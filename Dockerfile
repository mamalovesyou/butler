FROM golang:1.17-alpine as builder

# Install Alpine Dependencies
RUN apk add --update make

WORKDIR /butler
COPY . .

RUN make dependencies
RUN make services

# butler-users service
FROM scratch as service-users
COPY --from=builder /butler/bin/butler-users /butler-users
ENTRYPOINT ["/butler-users", "start"]

# butler-gateway service
FROM scratch as service-gateway
COPY --from=builder /butler/bin/butler-gateway /butler-gateway
ENTRYPOINT ["/butler-gateway", "start"]

