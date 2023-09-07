package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Kazai0/grpc-client/cmd/api"
	"github.com/Kazai0/grpc-client/cmd/utils"
	"github.com/Kazai0/grpc-client/pb"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

// Create Server Client with fiber
func main() {
	app := fiber.New()

	app.Get("/user", func(c *fiber.Ctx) error {
		res := CallGrpc(c)
		return c.JSON(res)
	})

	app.Listen(":3000")
}

//call to server gRPC
func CallGrpc(c *fiber.Ctx) *pb.UserResponse {
	connection, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer connection.Close()

	body := &api.User{}

	c.BodyParser(body)

	// Gere os bytes aleatórios
	randomBytes, err := utils.GenerateRandomBytes()
	if err != nil {
		fmt.Printf("Erro ao gerar bytes aleatórios: %v\n", err)
	}

	client := pb.NewUserServiceClient(connection)

	req := &pb.UserRequest{User: &pb.User{
		Id:    randomBytes,
		Name:  body.Name,
		Email: body.Email,
	}}

	res, err := client.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro ao criar o usuário: %v", err)
	}

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	return res
}
