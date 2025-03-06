package controllers

import (
	"context"
	"main/newspaper/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewGetAllController struct {
	UseCase *application.GetAllNewsUseCase
}

func GetAllNewsController(useCase *application.GetAllNewsUseCase) *NewGetAllController {
	return &NewGetAllController{UseCase: useCase}
}

func (nc *NewGetAllController) GetAllNewsHandler(c *gin.Context) {
	news, err := nc.UseCase.Execute(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, news)
}
