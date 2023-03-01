package methods

import (
	"context"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
	"google.golang.org/grpc/status"
)

func ExecuteSqrt(c pb.CalculatorClient, input int32) (float32, error) {
	sqrtResult, err := c.Sqrt(context.Background(), &pb.SqrtInput{
		Input: input,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error from Server : %v\n", e.Message())
			log.Printf("Error Code : %v\n", e.Code())
		} else {
			log.Fatalf("A none gRPC Error : %v", err)
		}
		return 0, err
	}

	return sqrtResult.Result, nil

}
