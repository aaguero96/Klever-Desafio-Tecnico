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

	serviceCollection := db.Collection("services")

	newService := bson.M{
		"name": in.GetName(),
		"site": in.GetSite(),
	}

	result, err := serviceCollection.InsertOne(context.TODO(), newService)
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

	serviceCollection := db.Collection("services")

	filter := bson.D{{}}
	filter = append(filter, bson.E{
		Key: "name",
		Value: bson.M{
			"$regex": primitive.Regex{Pattern: "^" + in.GetName() + ".*", Options: "i"},
		},
	})

	cur, err := serviceCollection.Find(context.TODO(), filter, options.Find())
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

func (s ServiceServer) ReadById(ctx context.Context, in *pb.ServiceId) (*pb.Service, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.Service{}, err
	}

	serviceCollection := db.Collection("services")

	userId, err := primitive.ObjectIDFromHex(in.GetServiceId())
	if err != nil {
		return &pb.Service{}, err
	}

	filter := bson.M{"_id": userId}

	var result *pb.Service

	err = serviceCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return &pb.Service{}, err
	}

	result.ServiceId = in.GetServiceId()
	return result, nil
}

func (s ServiceServer) Update(ctx context.Context, in *pb.Service) (*pb.EmptyService, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.EmptyService{}, err
	}

	serviceCollection := db.Collection("services")

	newUser := bson.M{
		"$set": bson.M{
			"name": in.GetName(),
			"site": in.GetSite(),
		},
	}

	userId, err := primitive.ObjectIDFromHex(in.GetServiceId())
	if err != nil {
		return &pb.EmptyService{}, err
	}

	filter := bson.M{"_id": userId}

	_, err = serviceCollection.UpdateOne(context.TODO(), filter, newUser)
	if err != nil {
		return &pb.EmptyService{}, err
	}

	return &pb.EmptyService{}, nil
}

func (s ServiceServer) Delete(ctx context.Context, in *pb.ServiceId) (*pb.EmptyService, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.EmptyService{}, err
	}

	serviceCollection := db.Collection("services")

	serviceId, err := primitive.ObjectIDFromHex(in.GetServiceId())
	if err != nil {
		return &pb.EmptyService{}, err
	}

	filter := bson.M{"_id": serviceId}

	_, err = serviceCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return &pb.EmptyService{}, nil
}

func ServiceService(s grpc.ServiceRegistrar, lis net.Listener) {
	pb.RegisterServiceServiceServer(s, &ServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
}
