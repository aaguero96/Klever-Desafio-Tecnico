package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/database"
	user_server "github.com/aaguero96/Klever-Desafio-Tecnico/gRPC_server/user"
	pb "github.com/aaguero96/Klever-Desafio-Tecnico/pb/user"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestMethodCreateByUserService(t *testing.T) {
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
	c := pb.NewUserServiceClient(conn)

	t.Run("When input is correct", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		input := &pb.NewUser{
			Name:     "Andre Aguero",
			Email:    "andre.luiz_1996@hotmail.com",
			Password: "123@Bc",
		}
		response, err := c.Create(ctx, input)
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Equal(t, response.Name, "Andre Aguero")
		assert.Equal(t, response.Email, "andre.luiz_1996@hotmail.com")
		assert.Equal(t, user_server.CheckPasswordHash("123@Bc", response.Password), true)
		cancel()
	})

	t.Run("When name input is incorrect", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		input := &pb.NewUser{
			Name:     "",
			Email:    "andre.luiz_1996@hotmail.com",
			Password: "123@Bc",
		}
		_, err := c.Create(ctx, input)
		assert.Contains(t, err.Error(), "Name is required")
		if err == nil {
			t.Errorf("Except error")
		}
		cancel()
	})

	t.Run("When email input is incorrect", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		inputs := []pb.NewUser{
			{
				Name:     "Andre Aguero",
				Email:    "",
				Password: "123@Bc",
			},
			{
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996",
				Password: "123@Bc",
			},
			{
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@",
				Password: "123@Bc",
			},
			{
				Name:     "Andre Aguero",
				Email:    "@hotmail.com",
				Password: "123@Bc",
			},
		}
		outputs := []string{
			"Email is required",
			"missing '@' or angle-addr",
			"no angle-addr",
			"no angle-addr",
		}
		for index, input := range inputs {
			_, err := c.Create(ctx, &input)
			assert.Contains(t, err.Error(), outputs[index])
			if err == nil {
				t.Errorf("Except error")
			}
		}
		cancel()
	})

	t.Run("When password input is incorrect", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		inputs := []pb.NewUser{
			{
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "",
			},
			{
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "12345",
			},
			{
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "123456",
			},
			{
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "123456A",
			},
			{
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "123456Aa",
			},
			{
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "123456A@",
			},
			{
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "aaaaaa",
			},
		}
		text := fmt.Sprintf(`
			Password needs:
				minimun of %d chars;
				1 uppercase char;
				1 lowercase char;
				1 number char;
				1 special char.
			You forgot one or more.
		`, user_server.MIN_PASSWORD)
		outputs := []string{
			"Password is required",
			text,
			text,
			text,
			text,
			text,
			text,
		}
		for index, input := range inputs {
			_, err := c.Create(ctx, &input)
			assert.Contains(t, err.Error(), outputs[index])
			if err == nil {
				t.Errorf("Except error")
			}
		}
		cancel()
	})

	// Clear database
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	db.Drop(ctx)
}

func TestMethodReadByUserService(t *testing.T) {
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
	c := pb.NewUserServiceClient(conn)

	// Add elements in db
	inputs := []pb.NewUser{
		{
			Name:     "Andre Aguero",
			Email:    "andre.luiz_1996@hotmail.com",
			Password: "123@BcANDRE",
		},
		{
			Name:     "Bruna Aguero",
			Email:    "brunaa@hotmail.com",
			Password: "123@BcBRUNA",
		},
		{
			Name:     "Amanda Domingues",
			Email:    "amandaa@hotmail.com",
			Password: "123@BcAMANDA",
		},
	}
	for _, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		_, err := c.Create(ctx, &input)
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When doesnt have input, return all users", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		response, err := c.Read(ctx, &pb.Filter{})
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		for index, user := range response.Users {
			assert.Equal(t, user.Name, inputs[index].Name)
			assert.Equal(t, user.Email, inputs[index].Email)
			assert.Equal(t, user_server.CheckPasswordHash(inputs[index].Password, user.Password), true)
		}
		cancel()
	})

	t.Run("When has input as first name, filter elements", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		response, err := c.Read(ctx, &pb.Filter{Name: "Andre"})
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Equal(t, len(response.Users), 1)
		assert.Equal(t, response.Users[0].Name, "Andre Aguero")
		assert.Equal(t, response.Users[0].Email, "andre.luiz_1996@hotmail.com")
		assert.Equal(t, user_server.CheckPasswordHash("123@BcANDRE", response.Users[0].Password), true)
		cancel()
	})

	t.Run("When has input as middle string, filter elements", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		response, err := c.Read(ctx, &pb.Filter{Name: "agu"})
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Equal(t, len(response.Users), 2)
		assert.Equal(t, response.Users[0].Name, "Andre Aguero")
		assert.Equal(t, response.Users[0].Email, "andre.luiz_1996@hotmail.com")
		assert.Equal(t, user_server.CheckPasswordHash("123@BcANDRE", response.Users[0].Password), true)
		assert.Equal(t, response.Users[1].Name, "Bruna Aguero")
		assert.Equal(t, response.Users[1].Email, "brunaa@hotmail.com")
		assert.Equal(t, user_server.CheckPasswordHash("123@BcBRUNA", response.Users[1].Password), true)
		cancel()
	})

	// Clear database
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	db.Drop(ctx)
}

func TestMethodReadByIdByUserService(t *testing.T) {
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
	c := pb.NewUserServiceClient(conn)

	// Add elements in db
	inputs := []pb.NewUser{
		{
			Name:     "Andre Aguero",
			Email:    "andre.luiz_1996@hotmail.com",
			Password: "123@BcANDRE",
		},
		{
			Name:     "Bruna Aguero",
			Email:    "brunaa@hotmail.com",
			Password: "123@BcBRUNA",
		},
		{
			Name:     "Amanda Domingues",
			Email:    "amandaa@hotmail.com",
			Password: "123@BcAMANDA",
		},
	}
	idInputs := make([]string, 3)
	for index, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		response, err := c.Create(ctx, &input)
		idInputs[index] = response.UserId
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When doesnt have input, return error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		_, err := c.ReadById(ctx, &pb.UserId{})
		if err == nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Contains(t, err.Error(), "the provided hex string is not a valid ObjectID")
		cancel()
	})

	t.Run("When have correct input, return element", func(t *testing.T) {
		for index := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			response, err := c.ReadById(ctx, &pb.UserId{UserId: idInputs[index]})
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.Name, inputs[index].Name)
			assert.Equal(t, response.Email, inputs[index].Email)
			assert.Equal(t, user_server.CheckPasswordHash(inputs[index].Password, response.Password), true)
			cancel()
		}
	})

	// Clear database
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	db.Drop(ctx)
}

func TestUpdateByUserService(t *testing.T) {
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
	c := pb.NewUserServiceClient(conn)

	// Add elements in db
	inputs := []pb.NewUser{
		{
			Name:     "Andre Aguero",
			Email:    "andre.luiz_1996@hotmail.com",
			Password: "123@BcANDRE",
		},
		{
			Name:     "Bruna Aguero",
			Email:    "brunaa@hotmail.com",
			Password: "123@BcBRUNA",
		},
		{
			Name:     "Amanda Domingues",
			Email:    "amandaa@hotmail.com",
			Password: "123@BcAMANDA",
		},
	}
	idInputs := make([]string, 3)
	for index, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		response, err := c.Create(ctx, &input)
		idInputs[index] = response.UserId
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When input is correct, return false", func(t *testing.T) {
		newInputs := []pb.User{
			{
				UserId:   idInputs[0],
				Name:     "Andre",
				Email:    "andre.luiz_1996@gmail.com",
				Password: "123@BcANDRE2",
			},
			{
				UserId:   idInputs[1],
				Name:     "Bruna",
				Email:    "brunaa@gmail.com",
				Password: "123@BcBRUNA2",
			},
			{
				UserId:   idInputs[2],
				Name:     "Amanda",
				Email:    "amandaa@gmail.com",
				Password: "123@BcAMANDA2",
			},
		}

		for _, newInput := range newInputs {
			// Update element
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			response, err := c.Update(ctx, &newInput)
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.Return, false)
			cancel()
		}
		for _, newInput := range newInputs {
			// Verify element
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			response, err := c.ReadById(ctx, &pb.UserId{UserId: newInput.UserId})
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.Name, newInput.Name)
			assert.Equal(t, response.Email, newInput.Email)
			assert.Equal(t, user_server.CheckPasswordHash(newInput.Password, response.Password), true)
			cancel()
		}
	})

	t.Run("When name is incorrect, return error", func(t *testing.T) {
		newInputs := []pb.User{
			{
				UserId:   idInputs[0],
				Name:     "",
				Email:    "andre.luiz_1996@gmail.com",
				Password: "123@BcANDRE2",
			},
		}
		outputs := []string{
			"Name is required",
		}

		for index, newInput := range newInputs {
			// Update element
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			_, err := c.Update(ctx, &newInput)
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), outputs[index])
			cancel()
		}
	})

	t.Run("When email is incorrect, return error", func(t *testing.T) {
		newInputs := []pb.User{
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "",
				Password: "123@Bc",
			},
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996",
				Password: "123@Bc",
			},
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@",
				Password: "123@Bc",
			},
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "@hotmail.com",
				Password: "123@Bc",
			},
		}
		outputs := []string{
			"Email is required",
			"missing '@' or angle-addr",
			"no angle-addr",
			"no angle-addr",
		}

		for index, newInput := range newInputs {
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			_, err := c.Update(ctx, &newInput)
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), outputs[index])
			cancel()
		}
	})

	t.Run("When password is incorrect, return error", func(t *testing.T) {
		newInputs := []pb.User{
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "",
			},
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "12345",
			},
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "123456",
			},
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "123456A",
			},
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "123456Aa",
			},
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "123456A@",
			},
			{
				UserId:   idInputs[0],
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@hotmail.com",
				Password: "aaaaaa",
			},
		}
		text := fmt.Sprintf(`
			Password needs:
				minimun of %d chars;
				1 uppercase char;
				1 lowercase char;
				1 number char;
				1 special char.
			You forgot one or more.
		`, user_server.MIN_PASSWORD)
		outputs := []string{
			"Password is required",
			text,
			text,
			text,
			text,
			text,
			text,
		}

		for index, newInput := range newInputs {
			// Update element
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			_, err := c.Update(ctx, &newInput)
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), outputs[index])
			cancel()
		}
	})

	t.Run("When userId is incorrect, return error", func(t *testing.T) {
		newInputs := []pb.User{
			{
				UserId:   "",
				Name:     "Andre Aguero",
				Email:    "andre.luiz_1996@gmail.com",
				Password: "123@BcANDRE2",
			},
		}
		outputs := []string{
			"the provided hex string is not a valid ObjectID",
		}

		for index, newInput := range newInputs {
			// Update element
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			_, err := c.Update(ctx, &newInput)
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), outputs[index])
			cancel()
		}
	})

	// Clear database
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	db.Drop(ctx)
}

func TestMethodDeleteByUserService(t *testing.T) {
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
	c := pb.NewUserServiceClient(conn)

	// Add elements in db
	inputs := []pb.NewUser{
		{
			Name:     "Andre Aguero",
			Email:    "andre.luiz_1996@hotmail.com",
			Password: "123@BcANDRE",
		},
		{
			Name:     "Bruna Aguero",
			Email:    "brunaa@hotmail.com",
			Password: "123@BcBRUNA",
		},
		{
			Name:     "Amanda Domingues",
			Email:    "amandaa@hotmail.com",
			Password: "123@BcAMANDA",
		},
	}
	idInputs := make([]string, 3)
	for index, input := range inputs {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		response, err := c.Create(ctx, &input)
		idInputs[index] = response.UserId
		if err != nil {
			t.Errorf("Internal error, problem with normal input")
		}
		cancel()
	}

	t.Run("When doesnt have input, return error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
		_, err := c.Delete(ctx, &pb.UserId{})
		if err == nil {
			t.Errorf("Internal error, problem with normal input")
		}
		assert.Contains(t, err.Error(), "the provided hex string is not a valid ObjectID")
		cancel()
	})

	t.Run("When have correct input, return false", func(t *testing.T) {
		// Delete elements
		for index := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			response, err := c.Delete(ctx, &pb.UserId{UserId: idInputs[index]})
			if err != nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Equal(t, response.Return, false)
			cancel()
		}
		// Check if elements were deleted
		for index := range inputs {
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			_, err := c.ReadById(ctx, &pb.UserId{UserId: idInputs[index]})
			if err == nil {
				t.Errorf("Internal error, problem with normal input")
			}
			assert.Contains(t, err.Error(), "no documents in result")
			cancel()
		}
	})

	// Clear database
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	db.Drop(ctx)
}
