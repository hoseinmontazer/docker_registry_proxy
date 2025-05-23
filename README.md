# 🛡️ Docker Registry Proxy

A lightweight Go-based proxy for Docker Hub, designed to bypass access restrictions caused by censorship, sanctions, or network-level blocks — especially those that affect individuals who are **not involved in any political or unlawful activity** but are unfairly impacted.

---

## ✨ Overview

This proxy transparently forwards Docker Registry API requests while handling authentication automatically — allowing Docker clients and tools (e.g., `docker pull`, CI/CD, etc.) to work as expected, even in restricted regions.

It supports:

- ✔️ Transparent proxying of Docker Hub's `/v2/` registry API
- 🔒 Automatic token-based authentication handling
- 🌐 CORS and logging middleware
- ⚙️ Configurable via environment variables
- 📡 Designed for individual developers, students, and professionals facing access issues due to unjust blanket sanctions or censorship

---

## 🚀 Why This Exists

In some regions, developers are being blocked from using global services like Docker Hub due to government-imposed **censorship or sanctions** — even when those individuals have **no involvement in politics or wrongdoing**.

This project exists to help:

> 🧑‍💻 Developers, researchers, and citizens who simply want to **build, learn, and collaborate** — not be punished for things beyond their control.

I believe in **open access to technology and knowledge**, and this tool is a small step toward defending that principle.

---

## 📦 Usage

### Prerequisites

- Go 1.20+
- Docker (if deploying via container)

### Run Locally

```bash
go run main.go
```
---
## 🚀 Live Deployment

- 🌐 **URL**: [https://registry.shirpala.ir](https://registry.shirpala.ir)
- 📦 **API Proxying**: Supports Docker Registry HTTP API (`/v2/`)
- 🔐 **Auth Handling**: Auto-handles token-based `Bearer` authentication

---

## 🧪 How to Use

Pull images via the proxy by configuring your Docker daemon or client.

### Example: Manual Pull via Proxy

```bash
docker pull registry.shirpala.ir/nginx:latest
```

# Docker Configuration for Default Registry

This guide explains how to configure Docker to use `registry.shirpala.ir` as the default registry.

## Steps

1. **Locate or Create the `daemon.json` File**

   Docker's configuration is stored in a file called `daemon.json`. This file is usually located in one of the following paths:

   - **Linux**: `/etc/docker/daemon.json`
   - **Windows**: `C:\ProgramData\docker\config\daemon.json`
   - **Mac**: `/etc/docker/daemon.json`

2. **Edit the `daemon.json` File**

   Open the `daemon.json` file with your preferred text editor. If it does not exist, create a new file. Add the following JSON configuration to set the default registry:

   ```json
   {
     "registry-mirrors": ["https://registry.shirpala.ir"]
   }
   ```

## 📄 License

This project is licensed under the **MIT License**.

You are free to use, modify, and distribute this software with proper attribution — but please, **use it responsibly** and **never for malicious or unethical purposes**.

See the [LICENSE](./LICENSE) file for full details.
