package controllers

import (
	"context"
	"main/newspaper/application"
	"main/newspaper/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewSaveController struct {
	UseCase *application.CreateNewUseCase
}

func NewsSaveController(useCase *application.CreateNewUseCase) *NewSaveController {
	return &NewSaveController{UseCase: useCase}
}

func (nc *NewSaveController) NewCreateNewHandler(c *gin.Context) {
	var new domain.News

	if err := c.ShouldBindJSON(&new); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	validations := []struct {
		condition bool
		err       error
	}{
		{new.Title == "", domain.ErrMissingTitle},
		{new.Author == "", domain.ErrMissingAuthor},
		{new.Content == "", domain.ErrMissingContent},
		{new.Description == "", domain.ErrMissingDescription},
		{new.PublicationDate == "", domain.ErrMissingPublicationDate},
		{new.Topic == "", domain.ErrMissingTopic},
	}

	for _, v := range validations {
		if v.condition {
			c.JSON(400, gin.H{"error": v.err.Error()})
			return
		}
	}

	if err := nc.UseCase.Execute(context.Background(), new); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Noticia creada exitosamente"})

}
