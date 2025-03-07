package controllers

import (
	"main/errores"
	"main/newspaper/application"

	"net/http"
    "errors" 
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewsDeleteController struct {
	UseCase *application.DeleteNewUseCase
}

func NewNewsDeleteController(useCase *application.DeleteNewUseCase) *NewsDeleteController {
	return &NewsDeleteController{UseCase: useCase}
}

func (nc *NewsDeleteController) DeleteNewsHandler(c *gin.Context){
	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        errores.SendErrorRespose(c, http.StatusBadRequest, errors.New("ID inválido")) 
        return
    }

    _, err = nc.UseCase.Execute(c.Request.Context(), objectId)
    if err != nil {
        errores.SendErrorRespose(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Noticia eliminado exitosamente", 
    })
}
