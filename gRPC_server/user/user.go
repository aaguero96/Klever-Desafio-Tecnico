package user_server

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/database"
	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s UserServer) Create(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.User{}, err
	}

	userCollection := db.Collection("users")

	newUser := bson.D{
		{"name", in.GetName()},
		{"email", in.GetEmail()},
		{"password", in.GetPassword()},
	}

	result, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return &pb.User{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return &pb.User{}, errors.New("Error after transform id in string")
	}

	return &pb.User{
		UserId:   oid.Hex(),
		Name:     in.GetName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}, nil
}

func UserService(s grpc.ServiceRegistrar, lis net.Listener) {
	pb.RegisterUserServiceServer(s, &UserServer{})
	log.Printf("server listening at %v", lis.Addr())
}
