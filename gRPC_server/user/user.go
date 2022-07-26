package user_server

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/user"
	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) Create(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Recieved: %v", in.GetName())
	var user_id uint32 = uint32(rand.Intn(1000))
	return &pb.User{
		UserId:   user_id,
		Name:     in.GetName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}, nil
}

func UserService(s grpc.ServiceRegistrar, lis net.Listener) {
	pb.RegisterUserServiceServer(s, &UserServer{})
	log.Printf("server listening at %v", lis.Addr())
}
