package infraestructure

import (
	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup, deps *Dependencies){
	group.GET("/", deps.GetAllNewsController.GetAllNewsHandler)
	group.GET("/:id", deps.GetNewsByIdController.GetAlbumsByIdHanlder)
	group.POST("/", deps.NewsSaveController.NewCreateNewHandler)
	group.PUT("/:id", deps.UpdateNewsController.UpdateNewsHandler)
	group.DELETE("/:id", deps.DeleteNewsController.DeleteNewsHandler)
	
}