package methods

import (
	"context"
	"log"
	"time"

	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
)

func ExecuteAvg(c pb.CalculatorClient) {
	reqs := []*pb.AvgInput{
		{Input: 1},
		{Input: 2},
		{Input: 3},
		{Input: 4},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average : %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req : %v", req.Input)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receving from Avg : %v\n", err)
	}

	log.Printf("Final Average : %v\n", res.Result)
}
