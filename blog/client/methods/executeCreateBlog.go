package methods

import (
	"context"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
)

func ExecuteCreateBlog(c pb.BlogClient) string {

	data := &pb.Post{
		UserId:  "Milky",
		Title:   "This is test.",
		Content: "Wow!!",
	}
	result, err := c.CreatePost(context.Background(), data)

	if err != nil {
		log.Fatalf("Unexpected Error Occured : %v", err)
	}

	log.Printf("Blog has been created : %s", result.GetId())

	return result.Id
}
