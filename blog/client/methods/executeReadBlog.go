package methods

import (
	"context"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
)

func ExecuteReadBlog(c pb.BlogClient, id string) *pb.Post {

	result, err := c.ReadPost(context.Background(), &pb.PostId{Id: id})

	if err != nil {
		log.Fatalf("Unexpected Error Occured : %v", err)
	}

	log.Printf("Blog Found : %s\n", result.Id)
	log.Printf("Blog Title : %s\n", result.Title)
	log.Printf("Blog User : %s\n", result.UserId)
	log.Printf("Blog Content : %s\n", result.Content)

	return result
}
