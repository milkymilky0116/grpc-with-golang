package util

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogServer) ReadPost(ctx context.Context, in *pb.PostId) (*pb.Post, error) {

	oid, err := primitive.ObjectIDFromHex(in.GetId())

	if err != nil {
		log.Fatalf("Cannot parse document id : %v", err)
	}

	data := &Post{}
	filter := bson.M{"_id": oid}

	post := Collection.FindOne(ctx, filter)

	if err := post.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find post with id %s", oid),
		)
	}

	return DocumentToPost(data), nil
}
