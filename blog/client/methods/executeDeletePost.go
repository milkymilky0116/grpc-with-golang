package methods

import (
	"context"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
)

func ExecuteDeleteBlog(c pb.BlogClient, id string) {
	result, err := c.DeletePost(context.Background(), &pb.PostId{Id: id})

	if err != nil {
		log.Fatalf("Error occured while deleting post : %v", err)
	}

	if result.Ok {
		log.Printf("%v was deleted successfully.", id)
	}
}
