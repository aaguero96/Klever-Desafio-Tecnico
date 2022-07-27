package service_server

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/database"
	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type ServiceServer struct {
	pb.UnimplementedServiceServiceServer
}

func (s ServiceServer) Create(ctx context.Context, in *pb.NewService) (*pb.Service, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.Service{}, err
	}

	userCollection := db.Collection("services")

	newService := bson.M{
		"name": in.GetName(),
		"site": in.GetSite(),
	}

	result, err := userCollection.InsertOne(context.TODO(), newService)
	if err != nil {
		return &pb.Service{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return &pb.Service{}, errors.New("Error after transform id in string")
	}

	return &pb.Service{
		ServiceId: oid.Hex(),
		Name:      in.GetName(),
		Site:      in.GetSite(),
	}, nil
}

func ServiceService(s grpc.ServiceRegistrar, lis net.Listener) {
	pb.RegisterServiceServiceServer(s, &ServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
}
