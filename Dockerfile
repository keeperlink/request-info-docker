# Build using a Go compiler
FROM golang:1.26-alpine AS builder
WORKDIR /src
COPY main.go .
ARG APP_VERSION=dev
# Compile a standalone Linux executable
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X main.version=${APP_VERSION}" -o /request-info main.go

# Minimal runtime image
FROM scratch
COPY --from=builder /request-info /request-info
# Run as non-root
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/request-info"]
