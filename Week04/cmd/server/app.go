// solmyr

package main

import "net"

import "google.golang.org/grpc"

type App struct {
	listener 	net.Listener
	gsrv 		*grpc.Server
}

func (app App) Start() error {
	return app.gsrv.Serve(app.listener)
}

