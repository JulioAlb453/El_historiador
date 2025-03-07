package controllers

import (
	"main/newspaper/application"
	"net/http"

	"main/errores"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewGetNewsByIdController struct {
	UseCase *application.GetNewByIdUseCase
}

func NewNewsGetByIdController(useCase *application.GetNewByIdUseCase) *NewGetNewsByIdController {
	return &NewGetNewsByIdController{UseCase: useCase}
}

func (nc *NewGetNewsByIdController) GetAlbumsByIdHanlder(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "ID Invalido por este motivo: " + err.Error()})
		return
	}

	new, err := nc.UseCase.Execute(c.Request.Context(), objectId)
	if err != nil {
		errores.SendErrorRespose(c, http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK, new)

}
