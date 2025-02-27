package main

import (
	"context"
	"log"
	"time"

	pb "github.com/D-arcjqw1/Auth/tree/main/grpc/api/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Создан пользователь: %s, %s", req.Name, req.Email)
	return &pb.CreateUserResponse{Id: time.Now().Unix()}, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Получение пользователя с ID: %d", req.Id)
	return &pb.GetUserResponse{
		Id:        req.Id,
		Name:      "Test User",
		Email:     "test@example.com",
		Role:      pb.Role_USER,
		CreatedAt: nil,
		UpdatedAt: nil,
	}, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Printf("Обновление пользователя ID: %d", req.Id)
	return &emptypb.Empty{}, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Printf("Удаление пользователя ID: %d", req.Id)
	return &emptypb.Empty{}, nil
}
