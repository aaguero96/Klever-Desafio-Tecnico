package upvote_server

import (
	"context"
	"errors"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/database"
	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/upvote"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpvoteServer struct {
	pb.UnimplementedUpvoteServiceServer
}

func (s UpvoteServer) Create(ctx context.Context, in *pb.NewUpvote) (*pb.Upvote, error) {
	db, err := database.Connect()
	if err != nil {
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
