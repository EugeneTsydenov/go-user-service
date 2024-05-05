package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	initEnv()
	initDB()
	run()
}

func run() {
	servicePort := getEnv()["SERVICE_PORT"]
	fmt.Println(servicePort)
	listener, err := net.Listen("tcp", servicePort)
	if err != nil {
		log.Fatalf("Error listening on port %s", servicePort)
	}
	fmt.Println("Serving gRPC on port", servicePort)
	startGrpcServer(listener)
}
