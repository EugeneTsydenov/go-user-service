package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func startGrpcServer(listener net.Listener) {
	grpcServer := grpc.NewServer()

	//services.AuthService(grpcServer)

	// Запускаем gRPC сервер в отдельной горутине
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
