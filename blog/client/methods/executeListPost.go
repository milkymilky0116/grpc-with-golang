package methods

import (
	"context"
	"io"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ExecuteListPost(c pb.BlogClient) {
	stream, err := c.ListPost(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while Calling ListPost : %v \n", err)
	}

	for {
		post, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream : %v \n", err)
		}
		log.Printf("Id: %v\n", post.Id)
		log.Printf("Title: %v\n", post.Title)
		log.Printf("Content: %v\n", post.Content)
		log.Printf("User: %v\n", post.UserId)

		log.Println("============================================")
	}
}
