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
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s ServiceServer) Read(ctx context.Context, in *pb.FilterService) (*pb.Services, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.Services{}, err
	}

	userCollection := db.Collection("services")

	filter := bson.D{{}}
	filter = append(filter, bson.E{
		Key: "name",
		Value: bson.M{
			"$regex": primitive.Regex{Pattern: "^" + in.GetName() + ".*", Options: "i"},
		},
	})

	cur, err := userCollection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return &pb.Services{}, err
	}
	defer cur.Close(context.TODO())

	var result []*pb.Service

	for cur.Next(context.TODO()) {
		type DecodedUser struct {
			ObjectID primitive.ObjectID `bson:"_id"`
			Name     string
			Site     string
		}
		var decodedUser DecodedUser

		if err = cur.Decode(&decodedUser); err != nil {
			return &pb.Services{}, err
		}
		user := pb.Service{
			Name:      decodedUser.Name,
			Site:      decodedUser.Site,
			ServiceId: decodedUser.ObjectID.Hex(),
		}
		result = append(result, &user)
	}

	return &pb.Services{Services: result}, nil
}

func ServiceService(s grpc.ServiceRegistrar, lis net.Listener) {
	pb.RegisterServiceServiceServer(s, &ServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
}
