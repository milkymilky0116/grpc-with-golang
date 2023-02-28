package methods

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
)

func ExecuteMax(c pb.CalculatorClient) {
	reqs := []*pb.MaxInput{
		{Input: 1},
		{Input: 5},
		{Input: 3},
		{Input: 6},
		{Input: 2},
		{Input: 20},
	}

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while Calling Max : %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending Request : %v", req.Input)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			_, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving request : %v", err)
			}

		}
		close(waitc)
	}()

	<-waitc
}
