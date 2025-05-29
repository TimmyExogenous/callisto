# Stage 1: Builder
FROM golang:1.21.12-bookworm AS builder

LABEL stage=builder

# Install necessary build dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends make git build-essential

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,id=callisto_gomod,target=/go/pkg/mod \
    go mod download -x
COPY . .
RUN --mount=type=cache,id=callisto_gobuild,target=/root/.cache/go-build \
    make build

# Stage 2: Final Runtime Image
FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN addgroup --system appgroup --gid 1001 && \
    adduser --system appuser --uid 1001 --ingroup appgroup

WORKDIR /home/appuser
COPY --from=builder /app/build/callisto /usr/local/bin/callisto
RUN chmod +x /usr/local/bin/callisto
USER appuser
ENTRYPOINT ["/usr/local/bin/callisto"]
CMD ["start", "--home", "/callisto/.callisto"]