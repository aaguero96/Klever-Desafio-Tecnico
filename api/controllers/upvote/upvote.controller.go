package user_controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/models"
	"github.com/aaguero96/Klever-Desafio-Tecnico/api/responses"
	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/upvote"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var upvote models.Upvote
	if err = json.Unmarshal(body, &upvote); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()

	c := pb.NewUpvoteServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.Create(ctx, &pb.NewUpvote{
		ServiceId: upvote.ServiceID,
		UserId:    upvote.UserID,
		Vote:      upvote.Vote,
		Comment:   upvote.Comment,
	})

	upvote.ID = response.UpvoteId
	responses.JSON(w, http.StatusCreated, upvote)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	upvoteId := params["upvoteId"]

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()

	c := pb.NewUpvoteServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.Delete(ctx, &pb.UpvoteId{
		UpvoteId: upvoteId,
	})

	responses.JSON(w, http.StatusNoContent, nil)
}
