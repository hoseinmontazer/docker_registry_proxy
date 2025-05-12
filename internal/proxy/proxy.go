package proxy

import (
	"io"
	"net/http"
	"registry_proxy/internal/auth"
	"registry_proxy/internal/logger"
)

const dockerHubBaseURL = "https://registry-1.docker.io"

func ProxyDockerRequest(w http.ResponseWriter, r *http.Request, dockerPath string) {
	fullURL := dockerHubBaseURL + dockerPath
	logger.Info("Proxying request to: %s", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		logger.Error("Failed to create initial request: %v", err)
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Accept", r.Header.Get("Accept"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("Upstream request failed: %v", err)
		http.Error(w, "Failed to contact upstream", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		logger.Info("Received 401 Unauthorized, attempting token flow...")
		authHeader := resp.Header.Get("WWW-Authenticate")
		resp.Body.Close()

		token, err := auth.GetAuthToken(authHeader)
		if err != nil {
			logger.Error("Failed to retrieve token: %v", err)
			http.Error(w, "Failed to get auth token", http.StatusInternalServerError)
			return
		}

		req, err = http.NewRequest("GET", fullURL, nil)
		if err != nil {
			logger.Error("Failed to create retry request: %v", err)
			http.Error(w, "Failed to create retry request", http.StatusInternalServerError)
			return
		}

		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Accept", r.Header.Get("Accept"))

		// logger.Info("Retrying with Authorization header: %s", req.Header.Get("Authorization"))
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			logger.Error("Upstream retry failed: %v", err)
			http.Error(w, "Upstream retry failed", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		// logger.Info("Retry successful, status code: %d", resp.StatusCode)

	}

	copyResponse(w, resp)
}

func copyResponse(w http.ResponseWriter, resp *http.Response) {
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
