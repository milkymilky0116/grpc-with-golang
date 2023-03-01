package util

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *BlogServer) ListPost(_ *emptypb.Empty, stream pb.Blog_ListPostServer) error {
	ctx := context.Background()
	cur, err := Collection.Find(ctx, primitive.D{})

	if err != nil {
		log.Fatalf("Error occured while Retrieving all Post.")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		data := &Post{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error occured while decoding data from mongo db : %v", err),
			)
		}
		stream.Send(DocumentToPost(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error : %v", err),
		)
	}
	return nil
}
