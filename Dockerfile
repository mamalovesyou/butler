FROM golang:1.17-alpine as dependencies

# Install Alpine Dependencies
RUN apk add --update --no-cache make ca-certificates git

WORKDIR /butler

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

# butler-users service
FROM scratch as service-users
COPY --from=builder /butler/bin/butler-users /butler-users
ENTRYPOINT ["/butler-users", "start"]

# butler-users service
FROM scratch as service-octopus
COPY --from=builder /butler/bin/butler-octopus /butler-octopus
ENTRYPOINT ["/butler-octopus", "start"]

# butler-gateway service
FROM scratch as service-gateway
COPY --from=builder /butler/bin/butler-gateway /butler-gateway
ENTRYPOINT ["/butler-gateway", "start"]

