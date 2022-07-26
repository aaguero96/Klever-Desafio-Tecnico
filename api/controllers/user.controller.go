package user_controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/models"
	"github.com/aaguero96/Klever-Desafio-Tecnico/api/responses"
	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/user"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.Create(ctx, &pb.NewUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	user.ID = response.UserId
	responses.JSON(w, http.StatusCreated, user)
}

func Read(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("userName"))

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.Read(ctx, &pb.Filter{
		Name: name,
	})

	responses.JSON(w, http.StatusCreated, response.Users)
}

func ReadById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.ReadById(ctx, &pb.UserId{
		UserId: userId,
	})

	responses.JSON(w, http.StatusCreated, response)
}
