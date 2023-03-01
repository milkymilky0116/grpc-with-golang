package main

import (
	"log"

	"github.com/MangoSteen0903/go-grpc-tour/blog/client/methods"
	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:3000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewBlogClient(conn)

	methods.ExecuteCreateBlog(c)
	methods.ExecuteCreateBlog(c)
	methods.ExecuteCreateBlog(c)
	methods.ExecuteCreateBlog(c)

	methods.ExecuteListPost(c)
	if err != nil {
		log.Println("Error Occured")
	}

}
