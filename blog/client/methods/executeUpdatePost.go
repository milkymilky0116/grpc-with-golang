package methods

import (
	"context"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
)

func ExecuteUpdateBlog(c pb.BlogClient, id string) {
	updatePost := &pb.Post{
		Id:      id,
		UserId:  "Update Milky",
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	ok, err := c.UpdatePost(context.Background(), updatePost)

	if err != nil {
		log.Fatalf("Error occured While Update Post : %v", err)
	}

	if ok.Ok {
		log.Printf("%v Post has been updated Successfully.\n", id)
	}
}
