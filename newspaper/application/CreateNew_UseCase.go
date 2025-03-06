package application

import (
	"context"
	"main/newspaper/domain"
)

type CreateNewUseCase struct {
	repo domain.INews
}

func NewCreateNewUseCase(repo domain.INews) *CreateNewUseCase {
	return &CreateNewUseCase{repo: repo}
}

func (uc *CreateNewUseCase) Execute(ctx context.Context, new domain.News) error {
	if err := uc.repo.Save(ctx, new); err != nil {
		return err
	}
	return nil
}
