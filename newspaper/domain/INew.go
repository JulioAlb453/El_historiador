package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type INews interface {
	Save(ctx context.Context, new News) error
	GetAllNews(ctx context.Context) ([]News, error)
	GetNewsById(ctx context.Context, id primitive.ObjectID) (News,error)
	Delete(ctx context.Context,id primitive.ObjectID ) (error)
	Update(ctx context.Context, new News) (News,error)
}