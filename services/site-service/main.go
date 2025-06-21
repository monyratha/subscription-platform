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
	pb.UnimplementedSiteServiceServer
}

func (s *server) CreateSite(ctx context.Context, req *pb.CreateSiteRequest) (*pb.CreateSiteResponse, error) {
	return &pb.CreateSiteResponse{SiteId: "site-123"}, nil
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterSiteServiceServer(s, &server{})
		log.Println("SiteService gRPC server running on :50052")
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
