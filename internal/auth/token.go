package auth

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"registry_proxy/internal/logger"
)

func GetAuthToken(authHeader string) (string, error) {
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", nil
	}

	authParts := strings.Split(authHeader[len(bearerPrefix):], ",")
	params := make(map[string]string)
	for _, part := range authParts {
		kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(kv) == 2 {
			key := kv[0]
			val := strings.Trim(kv[1], `"`)
			params[key] = val
		}
	}

	tokenURL := params["realm"] + "?service=" + url.QueryEscape(params["service"]) + "&scope=" + url.QueryEscape(params["scope"])
	logger.Info("Fetching auth token from: %s", tokenURL)

	resp, err := http.Get(tokenURL)
	if err != nil {
		logger.Error("HTTP request for token failed: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		logger.Error("Failed to decode token response: %v", err)
		return "", err
	}

	logger.Info("Received token: [REDACTED]")
	return data.Token, nil
}
