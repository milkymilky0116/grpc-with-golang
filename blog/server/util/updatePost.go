package util

import (
	"context"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *BlogServer) UpdatePost(ctx context.Context, in *pb.Post) (*pb.Status, error) {
	oid, err := primitive.ObjectIDFromHex(in.GetId())

	if err != nil {
		log.Fatalf("Cannot parse document id : %v", err)
	}

	filter := bson.M{"_id": oid}
	data := &Post{
		UserID:  in.GetUserId(),
		Title:   in.GetTitle(),
		Content: in.GetContent(),
	}

	_, err = Collection.UpdateOne(ctx, filter, bson.M{"$set": data})

	if err != nil {
		return &pb.Status{
			Ok: false,
		}, err
	}

	return &pb.Status{
		Ok: true,
	}, nil
}
