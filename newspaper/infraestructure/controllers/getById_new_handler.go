package controllers

import (
	"main/newspaper/application"
	"net/http"

	"main/errores"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetNewByIdController struct {
	UseCase *application.GetNewByIdUseCase
}

func NewGetByIdController(useCase *application.GetNewByIdUseCase) *GetNewByIdController {
	return &GetNewByIdController{UseCase: useCase}
}

func (nc *GetNewByIdController) GetAlbumsByIdHanlder(c *gin.Context) {
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
