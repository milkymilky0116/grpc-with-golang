package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/MangoSteen0903/go-grpc-tour/blog/server/util"
	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const PORT = "3000"

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	util.Collection = client.Database("grpc_blog").Collection("blog")

	port := fmt.Sprintf(":%s", PORT)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBlogServer(s, &util.BlogServer{})

	log.Printf("Server Listening on http://localhost:%s", PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
