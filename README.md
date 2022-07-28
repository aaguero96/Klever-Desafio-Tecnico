# Klever technical challange

# # Description

The Technical Challenge consists of creating an API with Golang using gRPC with stream pipes that exposes an upvote service endpoints.

# # Requirements

- Technical requirements:
  - Keep the code in Github.

- API:
  - The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string;
  - The structs used with your mongo model should support Marshal/Unmarshal with bson, json and struct;
  - The API should contain unit test of methods it uses.

- Extra:
  - Deliver the whole solution running in some free cloud service.

# # Deadline

- `📅 29 jul 2022 - Friday`

# # Run application (local)

1. clone git reposiory
- `git@github.com:aaguero96/Klever-Desafio-Tecnico.git` for SSH or `https://github.com/aaguero96/Klever-Desafio-Tecnico.git` for HTTPS.

2. Docker run mongoDB
- `docker run -d --name=mongo -p 27017:27017 bitnami/mongodb`, this step could be skipped if you have mongoDB runing on port 27017.

3. Install requirements
- `go mod tidy`.

4. Run gRPC service
- `go run gRPC_server/main.go`
- In this item you have permissions to do requests to localhost:50052 with gRPC API.

5. Run API service
- `go run api/main.go`
- In this item you have permissions to do requests to localhost:5000 with protocol http as client of gRPC server.
- This item is an extra for applicatio, because test requirements need only gRPC API.

# # Manual tests with Postman (local)

# # # For gRPC requests

1. Create API in postman
  - In field `APIs`, click on button `New` on top of page:
<img src="https://github.com/aaguero96/Klever-Desafio-Tecnico/blob/main/images/gPRC_1.png?raw=true"">
  - Click on `gRPC request`:
<img src="https://github.com/aaguero96/Klever-Desafio-Tecnico/blob/main/images/gPRC_2.png?raw=true"">
  - In `Select method` click on `Import a .proto file`:
<img src="https://github.com/aaguero96/Klever-Desafio-Tecnico/blob/main/images/gPRC_3.png?raw=true"">
  - Select an proto file, in this repository have three proto files at `proto/*.proto`, you have to do this, and next itens in item 1., procedures three times to test all requests.
  - Click on button `Next`:
<img src="https://github.com/aaguero96/Klever-Desafio-Tecnico/blob/main/images/gPRC_4.png?raw=true"">
  - Name API in fied `API name` and name Version in field `Version name`:
<img src="https://github.com/aaguero96/Klever-Desafio-Tecnico/blob/main/images/gPRC_5.png?raw=true"">
  - Click on `Import as API`.
  - In field `Enter server URL` insert `localhost:50052`.
  - Select field `message`:
<img src="https://github.com/aaguero96/Klever-Desafio-Tecnico/blob/main/images/gPRC_6.png?raw=true"">

2. Follow this methods:

  - For `UserService`:
    - `Create`
      - Input:
  ```javascript
  {
    "name": "André Aguero",
    "email": "andre@email.com",
    "password": "123456"
  }
      - Output:
  ```javascript
  {
    "userId": "62e26ffbb8caab0ad7c8db0a",
    "name": "André Aguero",
    "email": "andre@email.com",
    "password": "123456"
  }