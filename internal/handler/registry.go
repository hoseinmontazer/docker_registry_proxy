package handler

import (
	"log"
	"net/http"
	"registry_proxy/internal/proxy"
	"strings"
)

func RegistryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Println("Request:", path)

		if !strings.HasPrefix(path, "/v2/") {
			http.NotFound(w, r)
			return
		}

		// Handle /v2/ ping
		if path == "/v2/" {
			w.WriteHeader(http.StatusOK)
			return
		}

		dockerPath := path
		segments := strings.Split(strings.TrimPrefix(path, "/v2/"), "/")
		if len(segments) >= 2 && !strings.Contains(segments[0], "/") && !strings.Contains(segments[0], ".") {
			dockerPath = "/v2/library/" + strings.Join(segments, "/")
		}

		proxy.ProxyDockerRequest(w, r, dockerPath)
	}
}
