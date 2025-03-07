package application

import (
	"context"
	"main/newspaper/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetNewByIdUseCase struct {
	repo domain.INews
}

func NewGetNewByIdUseCase(repo domain.INews) *GetNewByIdUseCase {
	return &GetNewByIdUseCase{repo: repo}
}

func (uc *GetNewByIdUseCase) Execute(ctx context.Context, id primitive.ObjectID) (domain.News, error) {
	new, err := uc.repo.GetNewsById(ctx, id)

	if err != nil {
		return domain.News{}, err
	}
	return new, nil
}
