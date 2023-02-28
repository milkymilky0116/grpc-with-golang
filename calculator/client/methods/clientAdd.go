package methods

import (
	"context"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
)

func ExecuteAdd(c pb.CalculatorClient) uint32 {
	addResult, err := c.Add(context.Background(), &pb.AddInput{A: 3, B: 4})

	if err != nil {
		log.Fatalf("Could not Calculate: %v", err)
	}

	result := addResult.GetResult()

	return result
}
