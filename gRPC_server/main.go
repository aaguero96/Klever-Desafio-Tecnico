package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aaguero96/Klever-Desafio-Tecnico/config"
	service_server "github.com/aaguero96/Klever-Desafio-Tecnico/gRPC_server/service"
	upvote_server "github.com/aaguero96/Klever-Desafio-Tecnico/gRPC_server/upvote"
	user_server "github.com/aaguero96/Klever-Desafio-Tecnico/gRPC_server/user"
	"google.golang.org/grpc"
)

func main() {
	config.LoadEnv()

	address := fmt.Sprintf(":%d", config.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	user_server.UserService(s, lis)
	service_server.ServiceService(s, lis)
	upvote_server.UpvoteService(s, lis)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
