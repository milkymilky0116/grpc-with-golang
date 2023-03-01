package util

import (
	"context"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *BlogServer) DeletePost(ctx context.Context, in *pb.PostId) (*pb.Status, error) {
	oid, err := primitive.ObjectIDFromHex(in.GetId())

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": oid}

	_, err = Collection.DeleteOne(ctx, filter)

	if err != nil {
		return &pb.Status{
			Ok: false,
		}, nil
	}

	return &pb.Status{
		Ok: true,
	}, nil
}
