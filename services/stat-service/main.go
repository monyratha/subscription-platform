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
	pb.UnimplementedStatServiceServer
}

func (s *server) RecordView(ctx context.Context, req *pb.RecordViewRequest) (*pb.RecordViewResponse, error) {
	return &pb.RecordViewResponse{Status: "ok"}, nil
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", ":50055")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterStatServiceServer(s, &server{})
		log.Println("StatService gRPC server running on :50055")
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
