package controllers

import (
    "main/newspaper/application" 
    "main/newspaper/domain"
    "main/errores"    
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "net/http"
    "errors"
)

type NewsUpdateController struct {
    UseCase *application.UpdateNewUseCase
}

func NewNewsUpdateController(useCase *application.UpdateNewUseCase) *NewsUpdateController {
    return &NewsUpdateController{UseCase: useCase}
}

func (ac *NewsUpdateController) UpdateNewsHandler(c *gin.Context) {
    id := c.Param("id")
    var news domain.News

    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        errores.SendErrorRespose(c, http.StatusBadRequest, errors.New("ID inválido"))
        return
    }

    if err := c.ShouldBindJSON(&news); err != nil {
        errores.SendErrorRespose(c, http.StatusBadRequest, domain.ErrInvalidData)
        return
    }

    news.Id = objectId

    updatedNews, err := ac.UseCase.Execute(c.Request.Context(), news)
    if err != nil {
        errores.SendErrorRespose(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Noticia actualizada exitosamente",
        "news":    updatedNews,
    })
}
