FROM golang:1.17-alpine as builder

# Install Alpine Dependencies
RUN apk add --update make

WORKDIR go/src/github.com/butlerhq/butler
COPY . .

RUN make dependencies
RUN make services

# butler-users service
FROM scratch as service-users
COPY --from=builder go/src/github.com/butlerhq/butler/bin/butler-users /butler-users
ENTRYPOINT ["/butler-users", "start"]

# butler-gateway service
FROM scratch as service-gateway
COPY --from=builder go/src/github.com/butlerhq/butler/bin/butler-gateway /butler-gateway
ENTRYPOINT ["/butler-gateway", "start"]

