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

		// Determine docker path
		trimmed := strings.TrimPrefix(path, "/v2/")
		segments := strings.Split(trimmed, "/")

		// Only insert "library/" if it's a single-part image (like "nginx")
		if len(segments) >= 2 && !strings.Contains(segments[0], "/") && !strings.Contains(segments[0], ".") {
			if len(segments) == 3 { // Example: /v2/nginx/manifests/latest
				segments[0] = "library/" + segments[0]
			}
		}

		// // If image is a top-level (official) image like "postgres", rewrite to "library/postgres"
		// if len(segments) >= 2 && !strings.Contains(segments[0], "/") && !strings.Contains(segments[0], ".") {
		// 	fmt.Println("seg --- >", segments, len(segments))
		// 	// segments[0] = "library/" + segments[0]
		// }

		dockerPath := "/v2/" + strings.Join(segments, "/")

		proxy.ProxyDockerRequest(w, r, dockerPath)
	}
}
