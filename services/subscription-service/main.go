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
	pb.UnimplementedSubscriptionServiceServer
}

func (s *server) ActivatePlan(ctx context.Context, req *pb.ActivatePlanRequest) (*pb.ActivatePlanResponse, error) {
	return &pb.ActivatePlanResponse{SubscriptionId: "sub-123"}, nil
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", ":50053")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterSubscriptionServiceServer(s, &server{})
		log.Println("SubscriptionService gRPC server running on :50053")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		httpServer()
	}()

	wg.Wait()
}
