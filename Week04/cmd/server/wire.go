// solmyr
//+build wireinject

package main

import (
	"google.golang.org/grpc"
	"github.com/google/wire"
	"net"
)

func NewListener() (net.Listener, error) {
	return net.Listen("tcp4", "0.0.0.0:5000")
}

func NewGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func initApp() (*App, error) {
	wire.Build(
		NewListener,
		NewGRPCServer,
		wire.Struct(new(App), "*"))
	return &App{}, nil
}
