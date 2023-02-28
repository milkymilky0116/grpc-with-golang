package util

import (
	"io"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
)

func (s *CalculatorServer) Average(stream pb.Calculator_AverageServer) error {

	var total float32
	var length float32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			result := total / length
			return stream.SendAndClose(&pb.AvgResult{
				Result: float32(result),
			})
		} else {
			log.Printf("Server : Receive %d", req.Input)
		}
		if err != nil {
			log.Fatalf("Error occur while receive request : %v\n", err)

		}
		total += float32(req.GetInput())
		length++
	}
}
