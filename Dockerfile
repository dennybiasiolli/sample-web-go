# Build stage
FROM golang:1.22-bookworm AS builder
WORKDIR /usr/src/sample-web-go
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY main.go .
RUN go build -v -o /usr/local/bin/sample-web-go ./...


# Deploy stage
FROM debian:bookworm-slim
COPY --from=builder /usr/local/bin/sample-web-go /usr/local/bin/sample-web-go
ENV GIN_MODE=release PORT=8080
EXPOSE 8080
CMD ["sample-web-go"]
