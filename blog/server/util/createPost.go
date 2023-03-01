package util

import (
	"context"
	"fmt"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogServer) CreatePost(ctx context.Context, in *pb.Post) (*pb.PostId, error) {
	data := Post{
		UserID:  in.GetUserId(),
		Title:   in.GetTitle(),
		Content: in.GetContent(),
	}

	result, err := Collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.PostId{
		Id: oid.Hex(),
	}, nil
}
