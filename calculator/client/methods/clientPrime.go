package methods

import (
	"context"
	"io"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
)

func ExecutePrime(c pb.CalculatorClient) {
	req := &pb.PrimeInput{
		Input: 120,
	}

	stream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while Calling Prime : %v \n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream : %v \n", err)
		}

		log.Printf("%v\n", msg.Result)
	}
}
