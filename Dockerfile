# The first stage container, for building the application
FROM golang:1.17 as builder

# Add the keys
ARG GIT_TOKEN
ARG GIT_USER

COPY . /app
WORKDIR /app

RUN GIT_TERMINAL_PROMPT=1 \
    GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go build -v --installsuffix cgo --ldflags="-s" -o /contract-engine-edge


# for running the application
FROM alpine:3.15

WORKDIR /

COPY --from=builder /contract-engine-edge /contract-engine-edge

ENTRYPOINT ["/contract-engine-edge"]
