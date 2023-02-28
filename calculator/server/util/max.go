package util

import (
	"io"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
)

// Bi-Directional Streaming : many to many
func (s *CalculatorServer) Max(stream pb.Calculator_MaxServer) error {
	currentMax := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receive Request : %v", err)
		}

		res := req.Input

		if res > uint32(currentMax) {
			currentMax = int(res)

			log.Printf("Current Max : %d", currentMax)
			err = stream.Send(&pb.MaxResult{
				Result: uint32(currentMax),
			})
		}

		if err != nil {
			log.Fatalf("Error while sending request : %v", err)
		}

	}
}
