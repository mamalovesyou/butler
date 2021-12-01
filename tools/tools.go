// +build tools

package tools

import (
	_ "github.com/cosmtrek/air"
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/joho/godotenv"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/matthieuberger/butler/cmd/victorinox"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
