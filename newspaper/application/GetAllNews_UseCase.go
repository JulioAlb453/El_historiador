package application

import (
	"context"
	"main/newspaper/domain"
)

type GetAllNewsUseCase struct {
	repo domain.INews
}

func NewGetAllNewsUseCase(repo domain.INews) *GetAllNewsUseCase {
	return &GetAllNewsUseCase{repo: repo}
}

func (uc *GetAllNewsUseCase) Execute(ctx context.Context) ([]domain.News, error) {
	return uc.repo.GetAllNews(ctx)
}
