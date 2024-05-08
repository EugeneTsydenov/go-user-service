package main

import (
	"github.com/EugeneTsydenov/go-user-service/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func startGrpcServer(listener net.Listener) {
	grpcServer := grpc.NewServer()

	service.UserService(grpcServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	grpcServer.GracefulStop()
}
