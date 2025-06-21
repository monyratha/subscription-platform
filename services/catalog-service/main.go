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
	pb.UnimplementedCatalogServiceServer
}

func (s *server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return &pb.CreateProductResponse{ProductId: "prod-123"}, nil
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", ":50054")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterCatalogServiceServer(s, &server{})
		log.Println("CatalogService gRPC server running on :50054")
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
