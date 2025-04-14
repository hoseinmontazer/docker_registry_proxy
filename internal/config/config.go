package config

import (
	"log"
	"os"
)

type Config struct {
	Addr         string
	DockerHubURL string
}

func Load() *Config {
	addr := os.Getenv("PROXY_ADDR")
	if addr == "" {
		addr = ":5000"
	}

	url := os.Getenv("DOCKER_HUB_URL")
	if url == "" {
		url = "https://registry-1.docker.io"
	}

	log.Println("Config loaded:", addr, url)
	return &Config{Addr: addr, DockerHubURL: url}
}
