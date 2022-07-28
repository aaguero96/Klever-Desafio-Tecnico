package upvote_server

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func validateId(db *mongo.Database, id string, collectionName string) error {
	if id == "" {
		text := fmt.Sprintf("Id for collection %s is required", collectionName)
		log.Println(text)
		return errors.New(text)
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	collection := db.Collection(collectionName)

	filter := bson.M{"_id": objectId}

	valid := collection.FindOne(context.TODO(), filter)
	if err = valid.Err(); err != nil {
		text := fmt.Sprintf("Id for collection %s is invalid", collectionName)
		log.Println(text)
		return errors.New(text)
	}

	return nil
}

func validateVote(vote string) error {
	if vote == "" {
		log.Println("Vote is required")
		return errors.New("Vote is required")
	}

	if vote != "up" && vote != "down" {
		log.Println(`Vote need to be "up" or "down"`)
		return errors.New(`Vote need to be "up" or "down"`)
	}

	return nil
}
