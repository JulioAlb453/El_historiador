package domain

import "context"

type INews interface {
	Save(ctx context.Context, new News) error
	Get(ctx context.Context, id string) (*News, error)
	GetNewsById(ctx context.Context, id string) (*News,error)
	ListNews(ctx context.Context) ([]News, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, new News) error
}