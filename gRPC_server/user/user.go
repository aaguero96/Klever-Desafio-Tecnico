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
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s UserServer) Read(ctx context.Context, in *pb.Filter) (*pb.Users, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.Users{}, err
	}

	userCollection := db.Collection("users")

	filter := bson.D{{}}
	filter = append(filter, bson.E{
		Key: "name",
		Value: bson.D{
			{"$regex", primitive.Regex{Pattern: "^" + in.GetName() + ".*", Options: "i"}},
		},
	})

	cur, err := userCollection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return &pb.Users{}, err
	}
	defer cur.Close(context.TODO())

	var result []*pb.User

	for cur.Next(context.TODO()) {
		user := pb.User{}
		if err = cur.Decode(&user); err != nil {
			return &pb.Users{}, err
		}
		result = append(result, &user)
	}

	return &pb.Users{Users: result}, nil
}

func UserService(s grpc.ServiceRegistrar, lis net.Listener) {
	pb.RegisterUserServiceServer(s, &UserServer{})
	log.Printf("server listening at %v", lis.Addr())
}
