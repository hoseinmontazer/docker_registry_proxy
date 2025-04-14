package main

import (
	"net/http"

	"registry_proxy/internal/config"
	"registry_proxy/internal/handler"
	"registry_proxy/internal/logger"
	"registry_proxy/internal/middleware"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RegistryHandler())

	handlerWithMiddleware := middleware.Logging(middleware.CORS(mux))

	logger.Info("Docker registry proxy running on %s", cfg.Addr)
	err := http.ListenAndServe(cfg.Addr, handlerWithMiddleware)
	if err != nil {
		logger.Error("Failed to start server: %v", err)
	}
}
