package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "subscription-platform/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// This is a stub implementation returning a dummy user id and token
	return &pb.RegisterResponse{UserId: "123", Token: "dummy-token"}, nil
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// gRPC server
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterAuthServiceServer(s, &server{})
		log.Println("AuthService gRPC server running on :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// HTTP server
	go func() {
		defer wg.Done()
		httpServer()
	}()

	wg.Wait()
}
