package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "github.com/D-arcjqw1/Auth/grpc/pkg/user"
)

func main() {
	// Подключение к серверу
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Пример вызова метода Create
	req := &pb.CreateUserRequest{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	res, err := client.Create(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при создании пользователя: %v", err)
	}

	fmt.Printf("Создан пользователь с ID: %d\n", res.Id)
}
