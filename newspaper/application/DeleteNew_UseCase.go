package application

import (
	"context"
	"main/newspaper/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteNewUseCase struct {
	repo domain.INews
}

func NewDeleteNewUseCase(repo domain.INews) *DeleteNewUseCase {
	return &DeleteNewUseCase{repo: repo}
}

func (uc *DeleteNewUseCase) Execute(ctx context.Context, id primitive.ObjectID) (domain.News, error) {
	err := uc.repo.DeleteNews(ctx, id)

	if err != nil {
		return domain.News{}, err
	}
	return domain.News{}, nil
}
