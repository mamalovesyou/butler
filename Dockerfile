FROM golang:1.17-alpine as dependencies

# Install Alpine Dependencies
RUN apk add --update --no-cache make ca-certificates git

WORKDIR /butler

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

COPY go.* ./
RUN go mod download

FROM dependencies as builder

COPY . .
RUN make vendor
RUN make tools
RUN make services

# butler-victorinox
FROM scratch as victorinox
COPY --from=builder /butler/bin/butler-victorinox /butler-victorinox


# butler base service
FROM alpine:3.14 as service-base
RUN apk add --no-cache ca-certificates openssl

# butler-users service
FROM service-base  as service-users
COPY --from=builder /butler/bin/butler-users /butler-users
COPY --from=builder /bin/grpc_health_probe /grpc_health_probe
ENTRYPOINT ["/butler-users", "start"]

# butler-users service
FROM service-base  as service-octopus
COPY --from=builder /butler/bin/butler-octopus /butler-octopus
COPY --from=builder /bin/grpc_health_probe /grpc_health_probe
ENTRYPOINT ["/butler-octopus", "start"]

# butler-gateway service
FROM scratch as service-gateway
COPY --from=builder /butler/bin/butler-gateway /butler-gateway
ENTRYPOINT ["/butler-gateway", "start"]

