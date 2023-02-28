package util

import (
	"context"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
)

// Unary API: One to One Communication
func (s *CalculatorServer) Add(ctx context.Context, in *pb.AddInput) (*pb.Result, error) {
	output := in.GetA() + in.GetB()

	log.Printf("Server : Adding %d with %d", in.GetA(), in.GetB())
	return &pb.Result{
		Result: output,
	}, nil
}
