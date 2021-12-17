FROM golang:1.17-alpine as builder

# Install Alpine Dependencies
RUN apk add --update --no-cache make ca-certificates git

WORKDIR /butler
#COPY go.mod go.sum ./
#RUN go mod download

COPY . .
RUN make vendor
RUN make tools
RUN make services

# butler-victorinox
FROM scratch as victorinox
COPY --from=builder /butler/bin/butler-victorinox /bin/butler-victorinox

# butler-users service
FROM scratch as service-users
COPY --from=builder /butler/bin/butler-users /butler-users
ENTRYPOINT ["/butler-users", "start"]

# butler-gateway service
FROM scratch as service-gateway
COPY --from=builder /butler/bin/butler-gateway /butler-gateway
ENTRYPOINT ["/butler-gateway", "start"]

