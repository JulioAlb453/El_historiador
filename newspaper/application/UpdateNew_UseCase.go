package application

import (
	"context"
	"errors"
	"main/newspaper/domain"
)

type UpdateNewUseCase struct {
	repo domain.INews
}

func NewUpdateNewUseCase(repo domain.INews) *UpdateNewUseCase {
	return &UpdateNewUseCase{repo: repo}
}

func (uc *UpdateNewUseCase) Execute(ctx context.Context, new domain.News) (domain.News, error) {
	if new.Title == "" {
		return domain.News{}, errors.New("El título es requerido")
	}
	if new.Author == "" {
		return domain.News{}, errors.New("El nombre del autor es requerido")
	}
	if new.Content == "" {
		return domain.News{}, errors.New("El contenido es requerido")
	}
	if new.Description == "" {
		return domain.News{}, errors.New("La descripción es requerida")
	}
	if new.PublicationDate == "" {
		return domain.News{}, errors.New("La fecha de publicación es requerida")
	}
	if new.Topic == "" {
		return domain.News{}, errors.New("El tema es requerido")
	}

	existingNew, err := uc.repo.GetNewsById(ctx, new.Id)

	if err != nil {
		return domain.News{}, err
	}

	existingNew.Title = new.Title
	existingNew.Author = new.Author
	existingNew.Content = new.Content
	existingNew.Description = new.Description
	existingNew.PublicationDate = new.PublicationDate
	existingNew.Topic = new.Topic

	updateNew, err := uc.repo.UpdateNews(ctx, existingNew)

	if err != nil {
		return domain.News{}, err
	}
	return updateNew, nil

}
