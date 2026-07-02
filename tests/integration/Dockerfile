# Stage 1: build the Go binary.
# NOTE: golang:1.25 tag must match go.mod's `go` directive. Update if the tag
# is not yet on Docker Hub (e.g. use golang:1.24 during RC period).
FROM golang:1.25 AS builder

WORKDIR /app

# Download modules in a separate layer so source edits don't bust the cache.
COPY go.mod go.sum ./
RUN go mod download

# frontend/dist must exist (pre-built via `make build-ui`) — it is embedded
# into the binary by frontend/ui.go. The Go build fails without it.
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o /pragmata ./cmd/pragmata

# Stage 2: install sqlmesh into a venv.
# Uses python:3.12-slim (glibc/Debian) so duckdb and other native packages
# resolve to prebuilt manylinux aarch64 wheels instead of building from source.
FROM python:3.12-slim AS pybuilder

RUN python3 -m venv /opt/venv && \
    /opt/venv/bin/pip install --no-cache-dir 'sqlmesh[clickhouse]'

# Stage 3: minimal runtime image.
FROM python:3.12-slim

COPY --from=builder /pragmata /usr/local/bin/pragmata
COPY --from=pybuilder /opt/venv /opt/venv

ENV PATH="/opt/venv/bin:$PATH"

EXPOSE 7181

ENTRYPOINT ["pragmata"]
