package upvote_server

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/database"
	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/upvote"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type UpvoteServer struct {
	pb.UnimplementedUpvoteServiceServer
}

func (s UpvoteServer) Create(ctx context.Context, in *pb.NewUpvote) (*pb.Upvote, error) {
	if err := validateVote(in.GetVote()); err != nil {
		return &pb.Upvote{}, err
	}

	db, err := database.Connect()
	if err != nil {
		return &pb.Upvote{}, err
	}

	if err := validateId(db, in.GetServiceId(), "services"); err != nil {
		return &pb.Upvote{}, err
	}

	if err := validateId(db, in.GetUserId(), "users"); err != nil {
		return &pb.Upvote{}, err
	}

	upvoteCollection := db.Collection("upvotes")

	newUpvote := bson.M{
		"service_id": in.GetServiceId(),
		"user_id":    in.GetUserId(),
		"vote":       in.GetVote(),
		"comment":    in.GetComment(),
	}

	result, err := upvoteCollection.InsertOne(context.TODO(), newUpvote)
	if err != nil {
		return &pb.Upvote{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return &pb.Upvote{}, errors.New("Error after transform id in string")
	}

	return &pb.Upvote{
		UpvoteId:  oid.Hex(),
		ServiceId: in.GetServiceId(),
		UserId:    in.GetUserId(),
		Vote:      in.GetVote(),
		Comment:   in.GetComment(),
	}, nil
}

func (s UpvoteServer) Read(ctx context.Context, in *pb.FilterUpvote) (*pb.Upvotes, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.Upvotes{}, err
	}

	upvoteCollection := db.Collection("upvotes")

	filter := bson.D{{}}
	filter = append(filter, bson.E{
		Key: "vote",
		Value: bson.M{
			"$regex": primitive.Regex{Pattern: in.GetType(), Options: "i"},
		},
	})

	cur, err := upvoteCollection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return &pb.Upvotes{}, err
	}
	defer cur.Close(context.TODO())

	var result []*pb.Upvote

	for cur.Next(context.TODO()) {
		type DecodedUpvote struct {
			ObjectID  primitive.ObjectID `bson:"_id"`
			ServiceID string             `bson:"service_id"`
			UserID    string             `bson:"user_id"`
			Vote      string
			Comment   string
		}
		var decodedUpvote DecodedUpvote

		if err = cur.Decode(&decodedUpvote); err != nil {
			return &pb.Upvotes{}, err
		}
		upvote := pb.Upvote{
			UpvoteId:  decodedUpvote.ObjectID.Hex(),
			ServiceId: decodedUpvote.ServiceID,
			UserId:    decodedUpvote.ServiceID,
			Vote:      decodedUpvote.Vote,
			Comment:   decodedUpvote.Comment,
		}
		result = append(result, &upvote)
	}

	return &pb.Upvotes{Upvotes: result}, nil
}

func (s UpvoteServer) ReadById(ctx context.Context, in *pb.UpvoteId) (*pb.Upvote, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.Upvote{}, err
	}

	upvoteCollection := db.Collection("upvotes")

	upvoteId, err := primitive.ObjectIDFromHex(in.GetUpvoteId())
	if err != nil {
		return &pb.Upvote{}, err
	}

	filter := bson.M{"_id": upvoteId}

	type DecodedUpvote struct {
		ObjectID  primitive.ObjectID `bson:"_id"`
		ServiceID string             `bson:"service_id"`
		UserID    string             `bson:"user_id"`
		Vote      string
		Comment   string
	}
	var decodedUpvote DecodedUpvote

	err = upvoteCollection.FindOne(context.TODO(), filter).Decode(&decodedUpvote)
	if err != nil {
		return &pb.Upvote{}, err
	}

	return &pb.Upvote{
		UpvoteId:  in.GetUpvoteId(),
		ServiceId: decodedUpvote.ServiceID,
		UserId:    decodedUpvote.UserID,
		Vote:      decodedUpvote.Vote,
		Comment:   decodedUpvote.Comment,
	}, nil
}

func (s UpvoteServer) Update(ctx context.Context, in *pb.Upvote) (*pb.EmptyUpvote, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.EmptyUpvote{}, err
	}

	userCollection := db.Collection("users")

	newUpvote := bson.M{
		"$set": bson.M{
			"service_id": in.GetServiceId(),
			"user_id":    in.GetUserId(),
			"vote":       in.GetVote(),
			"comment":    in.GetComment(),
		},
	}

	userId, err := primitive.ObjectIDFromHex(in.GetUserId())
	if err != nil {
		return &pb.EmptyUpvote{}, err
	}

	filter := bson.M{"_id": userId}

	_, err = userCollection.UpdateOne(context.TODO(), filter, newUpvote)
	if err != nil {
		return &pb.EmptyUpvote{}, err
	}

	return &pb.EmptyUpvote{}, nil
}

func (s UpvoteServer) Delete(ctx context.Context, in *pb.UpvoteId) (*pb.EmptyUpvote, error) {
	db, err := database.Connect()
	if err != nil {
		return &pb.EmptyUpvote{}, err
	}

	upvoteCollection := db.Collection("upvotes")

	userId, err := primitive.ObjectIDFromHex(in.GetUpvoteId())
	if err != nil {
		return &pb.EmptyUpvote{}, err
	}

	filter := bson.M{"_id": userId}

	_, err = upvoteCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return &pb.EmptyUpvote{}, err
	}

	return &pb.EmptyUpvote{}, nil
}

func UpvoteService(s grpc.ServiceRegistrar, lis net.Listener) {
	pb.RegisterUpvoteServiceServer(s, &UpvoteServer{})
	log.Printf("server listening at %v", lis.Addr())
}
