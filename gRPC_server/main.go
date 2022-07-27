package main

import (
	"log"
	"net"

	service_server "github.com/aaguero96/Klever-Desafio-Tecnico/gRPC_server/service"
	user_server "github.com/aaguero96/Klever-Desafio-Tecnico/gRPC_server/user"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	user_server.UserService(s, lis)
	service_server.ServiceService(s, lis)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
