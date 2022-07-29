package test

import (
	"context"
	"testing"
	"time"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/database"
	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/service"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var TIME_WAIT = time.Second * 3

func TestMethodCreateByServiceService(t *testing.T) {
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
	c := pb.NewServiceServiceClient(conn)

	t.Run("When input is correct", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		input := &pb.NewService{
			Name: "Klever",
			Site: "https://klever.io/",
		}
		response, err := c.Create(ctx, input)
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Equal(t, response.Name, "Klever")
		assert.Equal(t, response.Site, "https://klever.io/")
		cancel()
	})

	t.Run("When name input is incorrect", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		input := &pb.NewService{
			Name: "",
			Site: "https://klever.io/",
		}
		_, err := c.Create(ctx, input)
		assert.Contains(t, err.Error(), "Name is required")
		if err == nil {
			t.Errorf("Except error")
		}
		cancel()
	})

	t.Run("When site input is incorrect", func(t *testing.T) {
		inputs := []pb.NewService{
			{
				Name: "klever",
				Site: "",
			},
			{
				Name: "klever",
				Site: "//klever.io/",
			},
		}
		outputs := []string{
			"Site is required",
			"Site is invalid",
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
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}

func TestMethodReadByServiceService(t *testing.T) {
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
	c := pb.NewServiceServiceClient(conn)

	// Add elements in db
	inputs := []pb.NewService{
		{
			Name: "Klever",
			Site: "https://klever.io/",
		},
		{
			Name: "Google",
			Site: "https://www.google.com.br/",
		},
		{
			Name: "Trybe",
			Site: "https://www.betrybe.com/",
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
		response, err := c.Read(ctx, &pb.FilterService{})
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		for index, user := range response.Services {
			assert.Equal(t, user.Name, inputs[index].Name)
			assert.Equal(t, user.Site, inputs[index].Site)
		}
		cancel()
	})

	t.Run("When has input as first name, filter elements", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Read(ctx, &pb.FilterService{Name: "Klever"})
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Equal(t, len(response.Services), 1)
		assert.Equal(t, response.Services[0].Name, "Klever")
		assert.Equal(t, response.Services[0].Site, "https://klever.io/")
		cancel()
	})

	t.Run("When has input as middle string, filter elements", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Read(ctx, &pb.FilterService{Name: "Le"})
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Equal(t, len(response.Services), 2)
		assert.Equal(t, response.Services[0].Name, "Klever")
		assert.Equal(t, response.Services[0].Site, "https://klever.io/")
		assert.Equal(t, response.Services[1].Name, "Google")
		assert.Equal(t, response.Services[1].Site, "https://www.google.com.br/")
		cancel()
	})

	// Clear database
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}

func TestMethodReadByIdByServiceService(t *testing.T) {
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
	c := pb.NewServiceServiceClient(conn)

	// Add elements in db
	inputs := []pb.NewService{
		{
			Name: "Klever",
			Site: "https://klever.io/",
		},
		{
			Name: "Google",
			Site: "https://www.google.com.br/",
		},
		{
			Name: "Trybe",
			Site: "https://www.betrybe.com/",
		},
	}
	idInputs := make([]string, 3)
	for index, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Create(ctx, &input)
		idInputs[index] = response.ServiceId
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When doesnt have input, return error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		_, err := c.ReadById(ctx, &pb.ServiceId{})
		if err == nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Contains(t, err.Error(), "the provided hex string is not a valid ObjectID")
		cancel()
	})

	t.Run("When have correct input, return element", func(t *testing.T) {
		for index := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			response, err := c.ReadById(ctx, &pb.ServiceId{ServiceId: idInputs[index]})
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.Name, inputs[index].Name)
			assert.Equal(t, response.Site, inputs[index].Site)
			cancel()
		}
	})

	// Clear database
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}

func TestUpdateByServiceService(t *testing.T) {
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
	c := pb.NewServiceServiceClient(conn)

	// Add elements in db
	inputs := []pb.NewService{
		{
			Name: "Klever",
			Site: "https://klever.io/",
		},
		{
			Name: "Google",
			Site: "https://www.google.com.br/",
		},
		{
			Name: "Trybe",
			Site: "https://www.betrybe.com/",
		},
	}
	idInputs := make([]string, 3)
	for index, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Create(ctx, &input)
		idInputs[index] = response.ServiceId
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When input is correct, return false", func(t *testing.T) {
		newInputs := []pb.Service{
			{
				ServiceId: idInputs[0],
				Name:      "Klever Site",
				Site:      "https://klever.io/",
			},
			{
				ServiceId: idInputs[1],
				Name:      "Google Site",
				Site:      "https://www.google.com.br/",
			},
			{
				ServiceId: idInputs[2],
				Name:      "Trybe Site",
				Site:      "https://www.betrybe.com/",
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
			response, err := c.ReadById(ctx, &pb.ServiceId{ServiceId: newInput.ServiceId})
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.Name, newInput.Name)
			assert.Equal(t, response.Site, newInput.Site)
			cancel()
		}
	})

	t.Run("When name is incorrect, return error", func(t *testing.T) {
		newInputs := []pb.Service{
			{
				ServiceId: idInputs[0],
				Name:      "",
				Site:      "https://klever.io/",
			},
		}
		outputs := []string{
			"Name is required",
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

	t.Run("When site is incorrect, return error", func(t *testing.T) {
		newInputs := []pb.Service{
			{
				ServiceId: idInputs[0],
				Name:      "klever",
				Site:      "",
			},
			{
				ServiceId: idInputs[0],
				Name:      "klever",
				Site:      "//klever.io/",
			},
		}
		outputs := []string{
			"Site is required",
			"Site is invalid",
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
		newInputs := []pb.Service{
			{
				ServiceId: "",
				Name:      "klever",
				Site:      "https://klever.io/",
			},
		}
		outputs := []string{
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
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}

func TestMethodDeleteByServiceService(t *testing.T) {
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
	c := pb.NewServiceServiceClient(conn)

	// Add elements in db
	inputs := []pb.NewService{
		{
			Name: "Klever",
			Site: "https://klever.io/",
		},
		{
			Name: "Google",
			Site: "https://www.google.com.br/",
		},
		{
			Name: "Trybe",
			Site: "https://www.betrybe.com/",
		},
	}
	idInputs := make([]string, 3)
	for index, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		response, err := c.Create(ctx, &input)
		idInputs[index] = response.ServiceId
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When doesnt have input, return error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
		_, err := c.Delete(ctx, &pb.ServiceId{})
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
			response, err := c.Delete(ctx, &pb.ServiceId{ServiceId: idInputs[index]})
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.Return, false)
			cancel()
		}
		// Check if elements were deleted
		for index := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
			_, err := c.ReadById(ctx, &pb.ServiceId{ServiceId: idInputs[index]})
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), "no documents in result")
			cancel()
		}
	})

	// Clear database
	ctx, cancel := context.WithTimeout(context.TODO(), TIME_WAIT)
	defer cancel()
	db.Drop(ctx)
}
