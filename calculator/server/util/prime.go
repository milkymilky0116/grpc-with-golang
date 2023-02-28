package util

import (
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
)

// Server Streaming: 1 to many
func (s *CalculatorServer) Prime(in *pb.PrimeInput, stream pb.Calculator_PrimeServer) error {

	input := int(in.GetInput())
	temp := input
	log.Printf("Server : Receive %d\n", input)
	i := 2
	for {
		if temp%i == 0 {
			stream.Send(&pb.PrimeResult{
				Result: uint32(i),
			})
			temp = temp / i
		} else {
			i++
		}

		if temp <= 0 {
			break
		}

	}
	return nil
}
