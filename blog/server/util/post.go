package util

import (
	pb "github.com/MangoSteen0903/go-grpc-tour/blog/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  string             `bson:"user_id,omitempty"`
	Title   string             `bson:"title,omitempty"`
	Content string             `bson:"content,omitempty"`
}

func DocumentToPost(data *Post) *pb.Post {
	return &pb.Post{
		Id:      data.ID.Hex(),
		UserId:  data.UserID,
		Title:   data.Title,
		Content: data.Content,
	}
}
