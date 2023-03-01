package util

import (
	"context"
	"math"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CalculatorServer) Sqrt(ctx context.Context, in *pb.SqrtInput) (*pb.SqrtOutput, error) {
	input := float64(in.GetInput())

	if input < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Input should be positive number",
		)
	}

	return &pb.SqrtOutput{
		Result: float32(math.Sqrt(input)),
	}, nil
}
