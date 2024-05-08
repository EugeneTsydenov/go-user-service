package main

import (
	"fmt"
	"github.com/EugeneTsydenov/go-user-service/cmd/user-service/db"
	"github.com/EugeneTsydenov/go-user-service/cmd/user-service/env"
	"log"
	"net"
)

func main() {
	env.InitEnv()
	db.InitDB()
	run()
}

func run() {
	servicePort := env.GetEnv()["SERVICE_PORT"]
	fmt.Println(servicePort)
	listener, err := net.Listen("tcp", servicePort)
	fmt.Println(err)
	if err != nil {
		log.Fatalf("Error listening on port %s", servicePort)
	}
	fmt.Println("Serving gRPC on port", servicePort)
	startGrpcServer(listener)
}
