package test

import (
	"context"
	"testing"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/database"
	pbService "github.com/aaguero96/Klever-Desafio-Tecnico/pb/service"
	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/upvote"
	pbUser "github.com/aaguero96/Klever-Desafio-Tecnico/pb/user"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestMethodCreateByUpvoteService(t *testing.T) {
	// Connect to database
	db, err := database.Connect()
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Connect to server gRPC
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Create new Upvote Client
	c := pb.NewUpvoteServiceClient(conn)

	// Add elements to user e server
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	cUser := pbUser.NewUserServiceClient(conn)
	user, err := cUser.Create(ctx, &pbUser.NewUser{
		Name:     "André Aguero",
		Email:    "andre.luiz_1996@hotmail.com",
		Password: "123@bC",
	})
	cService := pbService.NewServiceServiceClient(conn)
	service, err := cService.Create(ctx, &pbService.NewService{
		Name: "Klever",
		Site: "https://klever.io/",
	})
	idsToInput := map[string]string{
		"user":    user.UserId,
		"service": service.ServiceId,
	}
	cancel()

	t.Run("When input is correct", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		input := &pb.NewUpvote{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment",
		}
		response, err := c.Create(ctx, input)
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Equal(t, response.ServiceId, idsToInput["service"])
		assert.Equal(t, response.UserId, idsToInput["user"])
		assert.Equal(t, response.Vote, "up")
		assert.Equal(t, response.Comment, "My comment")
		cancel()
	})

	t.Run("When vote input is incorrect", func(t *testing.T) {
		inputs := []pb.NewUpvote{
			{
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["user"],
				Vote:      "",
				Comment:   "My comment",
			},
			{
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["user"],
				Vote:      "another",
				Comment:   "My comment",
			},
		}
		outputs := []string{
			"Vote is required",
			`Vote need to be "up" or "down"`,
		}
		for index, input := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			_, err := c.Create(ctx, &input)
			assert.Contains(t, err.Error(), outputs[index])
			if err == nil {
				t.Errorf("Except error")
			}
			cancel()
		}
	})

	t.Run("When serviceId input is incorrect", func(t *testing.T) {
		inputs := []pb.NewUpvote{
			{
				ServiceId: "",
				UserId:    idsToInput["user"],
				Vote:      "up",
				Comment:   "My comment",
			},
			{
				ServiceId: idsToInput["user"],
				UserId:    idsToInput["user"],
				Vote:      "up",
				Comment:   "My comment",
			},
			{
				ServiceId: "1",
				UserId:    idsToInput["user"],
				Vote:      "up",
				Comment:   "My comment",
			},
		}
		outputs := []string{
			"Id for collection services is required",
			"Id for collection services is invalid",
			"the provided hex string is not a valid ObjectID",
		}
		for index, input := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			_, err := c.Create(ctx, &input)
			assert.Contains(t, err.Error(), outputs[index])
			if err == nil {
				t.Errorf("Except error")
			}
			cancel()
		}
	})

	t.Run("When userId input is incorrect", func(t *testing.T) {
		inputs := []pb.NewUpvote{
			{
				ServiceId: idsToInput["service"],
				UserId:    "",
				Vote:      "up",
				Comment:   "My comment",
			},
			{
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["service"],
				Vote:      "up",
				Comment:   "My comment",
			},
			{
				ServiceId: idsToInput["service"],
				UserId:    "1",
				Vote:      "up",
				Comment:   "My comment",
			},
		}
		outputs := []string{
			"Id for collection users is required",
			"Id for collection users is invalid",
			"the provided hex string is not a valid ObjectID",
		}
		for index, input := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			_, err := c.Create(ctx, &input)
			assert.Contains(t, err.Error(), outputs[index])
			if err == nil {
				t.Errorf("Except error")
			}
			cancel()
		}
	})

	// Clear database
	ctx, cancel = context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}

func TestMethodReadByUpvoteService(t *testing.T) {
	// Connect to database
	db, err := database.Connect()
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Connect to server gRPC
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Create new User Client
	c := pb.NewUpvoteServiceClient(conn)

	// Add elements to user e server
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	cUser := pbUser.NewUserServiceClient(conn)
	user, err := cUser.Create(ctx, &pbUser.NewUser{
		Name:     "André Aguero",
		Email:    "andre.luiz_1996@hotmail.com",
		Password: "123@bC",
	})
	cService := pbService.NewServiceServiceClient(conn)
	service, err := cService.Create(ctx, &pbService.NewService{
		Name: "Klever",
		Site: "https://klever.io/",
	})
	idsToInput := map[string]string{
		"user":    user.UserId,
		"service": service.ServiceId,
	}
	cancel()

	// Add elements in db
	inputs := []pb.NewUpvote{
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 1",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 2",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 3",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "down",
			Comment:   "My comment 4",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "down",
			Comment:   "My comment 5",
		},
	}
	for _, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		_, err := c.Create(ctx, &input)
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When doesnt have input, return all services", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Read(ctx, &pb.FilterUpvote{})
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		for index, upvote := range response.Upvotes {
			assert.Equal(t, upvote.ServiceId, inputs[index].ServiceId)
			assert.Equal(t, upvote.UserId, inputs[index].UserId)
			assert.Equal(t, upvote.Vote, inputs[index].Vote)
			assert.Equal(t, upvote.Comment, inputs[index].Comment)
		}
		cancel()
	})

	t.Run("When has input type=up, filter elements", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Read(ctx, &pb.FilterUpvote{Type: "up"})
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Equal(t, len(response.Upvotes), 3)
		cancel()
	})

	t.Run("When has input type=down, filter elements", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Read(ctx, &pb.FilterUpvote{Type: "down"})
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Equal(t, len(response.Upvotes), 2)
		cancel()
	})

	// Clear database
	ctx, cancel = context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}

func TestMethodReadByIdByUpvoteService(t *testing.T) {
	// Connect to database
	db, err := database.Connect()
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Connect to server gRPC
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Create new User Client
	c := pb.NewUpvoteServiceClient(conn)

	// Add elements to user e server
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	cUser := pbUser.NewUserServiceClient(conn)
	user, err := cUser.Create(ctx, &pbUser.NewUser{
		Name:     "André Aguero",
		Email:    "andre.luiz_1996@hotmail.com",
		Password: "123@bC",
	})
	cService := pbService.NewServiceServiceClient(conn)
	service, err := cService.Create(ctx, &pbService.NewService{
		Name: "Klever",
		Site: "https://klever.io/",
	})
	idsToInput := map[string]string{
		"user":    user.UserId,
		"service": service.ServiceId,
	}
	cancel()

	// Add elements in db
	inputs := []pb.NewUpvote{
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 1",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 2",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 3",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "down",
			Comment:   "My comment 4",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "down",
			Comment:   "My comment 5",
		},
	}
	idInputs := make([]string, 5)
	for index, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Create(ctx, &input)
		idInputs[index] = response.UpvoteId
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When doesnt have input, return error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		_, err := c.ReadById(ctx, &pb.UpvoteId{})
		if err == nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Contains(t, err.Error(), "the provided hex string is not a valid ObjectID")
		cancel()
	})

	t.Run("When have correct input, return element", func(t *testing.T) {
		for index := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			response, err := c.ReadById(ctx, &pb.UpvoteId{UpvoteId: idInputs[index]})
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.ServiceId, inputs[index].ServiceId)
			assert.Equal(t, response.UserId, inputs[index].UserId)
			assert.Equal(t, response.Vote, inputs[index].Vote)
			assert.Equal(t, response.Comment, inputs[index].Comment)
			cancel()
		}
	})

	// Clear database
	ctx, cancel = context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}

func TestUpdateByUpvoteService(t *testing.T) {
	// Connect to database
	db, err := database.Connect()
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Connect to server gRPC
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Create new User Client
	c := pb.NewUpvoteServiceClient(conn)

	// Add elements to user e server
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	cUser := pbUser.NewUserServiceClient(conn)
	user, err := cUser.Create(ctx, &pbUser.NewUser{
		Name:     "André Aguero",
		Email:    "andre.luiz_1996@hotmail.com",
		Password: "123@bC",
	})
	cService := pbService.NewServiceServiceClient(conn)
	service, err := cService.Create(ctx, &pbService.NewService{
		Name: "Klever",
		Site: "https://klever.io/",
	})
	idsToInput := map[string]string{
		"user":    user.UserId,
		"service": service.ServiceId,
	}
	cancel()

	// Add elements in db
	inputs := []pb.NewUpvote{
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 1",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 2",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 3",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "down",
			Comment:   "My comment 4",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "down",
			Comment:   "My comment 5",
		},
	}
	idInputs := make([]string, 5)
	for index, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Create(ctx, &input)
		idInputs[index] = response.UpvoteId
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When input is correct, return false", func(t *testing.T) {
		newInputs := []pb.Upvote{
			{
				UpvoteId:  idInputs[0],
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["user"],
				Vote:      "down",
				Comment:   "My comment 1.1",
			},
			{
				UpvoteId:  idInputs[1],
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["user"],
				Vote:      "down",
				Comment:   "My comment 2.1",
			},
			{
				UpvoteId:  idInputs[2],
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["user"],
				Vote:      "down",
				Comment:   "My comment 3.1",
			},
			{
				UpvoteId:  idInputs[3],
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["user"],
				Vote:      "up",
				Comment:   "My comment 4.1",
			},
			{
				UpvoteId:  idInputs[4],
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["user"],
				Vote:      "up",
				Comment:   "My comment 5.1",
			},
		}

		for _, newInput := range newInputs {
			// Update element
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			response, err := c.Update(ctx, &newInput)
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.Return, false)
			cancel()
		}
		for _, newInput := range newInputs {
			// Verify element
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			response, err := c.ReadById(ctx, &pb.UpvoteId{UpvoteId: newInput.UpvoteId})
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.ServiceId, newInput.ServiceId)
			assert.Equal(t, response.UserId, newInput.UserId)
			assert.Equal(t, response.Vote, newInput.Vote)
			assert.Equal(t, response.Comment, newInput.Comment)
			cancel()
		}
	})

	t.Run("When vote is incorrect, return error", func(t *testing.T) {
		newInputs := []pb.Upvote{
			{
				UpvoteId:  idInputs[0],
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["user"],
				Vote:      "",
				Comment:   "My comment 1.1",
			},
		}
		outputs := []string{
			"Vote is required",
		}

		for index, newInput := range newInputs {
			// Update element
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			_, err := c.Update(ctx, &newInput)
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), outputs[index])
			cancel()
		}
	})

	t.Run("When serviceId is incorrect, return error", func(t *testing.T) {
		newInputs := []pb.Upvote{
			{
				UpvoteId:  idInputs[0],
				ServiceId: "",
				UserId:    idsToInput["user"],
				Vote:      "up",
				Comment:   "My comment",
			},
			{
				UpvoteId:  idInputs[0],
				ServiceId: idsToInput["user"],
				UserId:    idsToInput["user"],
				Vote:      "up",
				Comment:   "My comment",
			},
			{
				UpvoteId:  idInputs[0],
				ServiceId: "1",
				UserId:    idsToInput["user"],
				Vote:      "up",
				Comment:   "My comment",
			},
		}
		outputs := []string{
			"Id for collection services is required",
			"Id for collection services is invalid",
			"the provided hex string is not a valid ObjectID",
		}

		for index, newInput := range newInputs {
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			_, err := c.Update(ctx, &newInput)
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), outputs[index])
			cancel()
		}
	})

	t.Run("When serviceId is incorrect, return error", func(t *testing.T) {
		newInputs := []pb.Upvote{
			{
				UpvoteId:  idInputs[0],
				ServiceId: idsToInput["service"],
				UserId:    "",
				Vote:      "up",
				Comment:   "My comment",
			},
			{
				UpvoteId:  idInputs[0],
				ServiceId: idsToInput["service"],
				UserId:    idsToInput["service"],
				Vote:      "up",
				Comment:   "My comment",
			},
			{
				UpvoteId:  idInputs[0],
				ServiceId: idsToInput["service"],
				UserId:    "1",
				Vote:      "up",
				Comment:   "My comment",
			},
		}
		outputs := []string{
			"Id for collection users is required",
			"Id for collection users is invalid",
			"the provided hex string is not a valid ObjectID",
		}

		for index, newInput := range newInputs {
			// Update element
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			_, err := c.Update(ctx, &newInput)
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), outputs[index])
			cancel()
		}
	})

	// Clear database
	ctx, cancel = context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}

func TestMethodDeleteByUpvoteService(t *testing.T) {
	// Connect to database
	db, err := database.Connect()
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Connect to server gRPC
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Errorf("Fail to connect with server")
	}

	// Create new User Client
	c := pb.NewUpvoteServiceClient(conn)

	// Add elements to user e server
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	cUser := pbUser.NewUserServiceClient(conn)
	user, err := cUser.Create(ctx, &pbUser.NewUser{
		Name:     "André Aguero",
		Email:    "andre.luiz_1996@hotmail.com",
		Password: "123@bC",
	})
	cService := pbService.NewServiceServiceClient(conn)
	service, err := cService.Create(ctx, &pbService.NewService{
		Name: "Klever",
		Site: "https://klever.io/",
	})
	idsToInput := map[string]string{
		"user":    user.UserId,
		"service": service.ServiceId,
	}
	cancel()

	// Add elements in db
	inputs := []pb.NewUpvote{
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 1",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 2",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "up",
			Comment:   "My comment 3",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "down",
			Comment:   "My comment 4",
		},
		{
			ServiceId: idsToInput["service"],
			UserId:    idsToInput["user"],
			Vote:      "down",
			Comment:   "My comment 5",
		},
	}
	idInputs := make([]string, 5)
	for index, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Create(ctx, &input)
		idInputs[index] = response.UpvoteId
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When doesnt have input, return error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		_, err := c.Delete(ctx, &pb.UpvoteId{})
		if err == nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Contains(t, err.Error(), "the provided hex string is not a valid ObjectID")
		cancel()
	})

	t.Run("When have correct input, return false", func(t *testing.T) {
		// Delete elements
		for index := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			response, err := c.Delete(ctx, &pb.UpvoteId{UpvoteId: idInputs[index]})
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.Return, false)
			cancel()
		}
		// Check if elements were deleted
		for index := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			_, err := c.ReadById(ctx, &pb.UpvoteId{UpvoteId: idInputs[index]})
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), "no documents in result")
			cancel()
		}
	})

	// Clear database
	ctx, cancel = context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}
