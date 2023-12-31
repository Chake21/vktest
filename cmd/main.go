package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"vktest/server"
	"vktest/vktest/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("SERVICE_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("cannot listen port")
		return
	}
	s := grpc.NewServer()
	api.RegisterVKTestServer(s, &server.Server{})
	reflection.Register(s)
	fmt.Println("Service started at ", port, " port")
	if err = s.Serve(lis); err != nil {
		log.Fatal("Cant launch server: ", err)
	}
}
