# syntax=docker/dockerfile:1

# ---- Build stage: compile a static Go binary ----
# Keep GO_VERSION in sync with the `go` line in go.mod.
ARG GO_VERSION=1.26.8
FROM golang:${GO_VERSION}-alpine AS build

WORKDIR /src

# Download dependencies first so this layer is cached until go.mod/go.sum change.
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

# CGO_ENABLED=0 gives a fully static binary that runs without libc.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/riven-api ./cmd/riven-api

# ---- Run stage: tiny image with just the binary ----
# distroless/static has CA certs and a non-root user, but no shell.
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/riven-api /riven-api

ENV PORT=8080
EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/riven-api"]
