FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN go build -o registry_proxy ./cmd/proxy

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates libc6 && rm -rf /var/lib/apt/lists/*
WORKDIR /root/
COPY --from=builder /app/registry_proxy .
EXPOSE 5000
CMD ["./registry_proxy"]
