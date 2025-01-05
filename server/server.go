package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net"
	api "simpleServer/api/grpc"
)

type Server struct {
	api.UnsafePersonServiceServer
	slog.Logger
	users map[int32]*api.Person
}

func (s *Server) GetPerson(_ context.Context, request *api.GetPersonRequest) (*api.PersonResponse, error) {
	if request.Id != nil {
		user, ok := s.users[*request.Id]
		if !ok {
			logger.Error("user not found")
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		logger.Info("user found")
		return &api.PersonResponse{Person: user}, nil
	}
	logger.Error("nil argument")
	return nil, status.Errorf(codes.InvalidArgument, "nil argument")
}
func (s *Server) AddPerson(_ context.Context, request *api.Person) (*api.PersonResponse, error) {
	if request.Id != nil && request.Name != "" {
		if _, ok := s.users[*request.Id]; ok {
			logger.Error("user already exists")
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}
		s.users[*request.Id] = request
		logger.Info("user added")
		return &api.PersonResponse{Person: request}, nil
	}
	logger.Error("nil argument")
	return nil, status.Errorf(codes.InvalidArgument, "nil argument")
}

func StartGrpcServer() error {
	lis, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", "localhost", 8080))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to listen: %v", err))
		return err
	}
	s := grpc.NewServer()
	api.RegisterPersonServiceServer(s, &Server{users: make(map[int32]*api.Person)})
	logger.Info(fmt.Sprintf("Starting server on port %d", 8080))
	if err = s.Serve(lis); err != nil {
		logger.Error(fmt.Sprintf("Failed to start server: %v", err))
		return err
	}
	return nil
}
