package infraestructure

import (
	"log"
	"main/core"
	"main/newspaper/application"
	"main/newspaper/infraestructure/controllers"
	"main/newspaper/infraestructure/repository"
)

type Dependencies struct {
	NewsSaveController    *controllers.NewSaveController
	GetAllNewsController  *controllers.NewNewsGetAllController
	GetNewsByIdController *controllers.NewGetNewsByIdController
	DeleteNewsController  *controllers.NewsDeleteController
	UpdateNewsController *controllers.NewsUpdateController
}

func Init() *Dependencies {
	conn := core.Connect()

	if conn == nil {
		log.Fatal("Error al conectar la base de datos")
	}
	db := conn.Database("El_historiador")

	newsRepo := repository.NewMongoNewsRepository(db)

	createNewsUseCase := application.NewCreateNewsUseCase(newsRepo)
	getAllNewsUseCase := application.NewGetAllNewsUseCase(newsRepo)
	getNewsByIdUseCase := application.NewGetNewByIdUseCase(newsRepo)
	deleteNewsUseCase := application.NewDeleteNewUseCase(newsRepo)
	updateNewsUseCase := application.NewUpdateNewUseCase(newsRepo)

	NewsSaveController := controllers.NewsSaveController(createNewsUseCase)
	getAllNewsController := controllers.GetAllNewsController(getAllNewsUseCase)
	getNewsByIdController := controllers.NewNewsGetByIdController(getNewsByIdUseCase)
	deleteNewsController := controllers.NewNewsDeleteController(deleteNewsUseCase)
	updateNewsController := controllers.NewNewsUpdateController(updateNewsUseCase)


	return &Dependencies{
		NewsSaveController:    NewsSaveController,
        GetAllNewsController:  getAllNewsController,
        GetNewsByIdController: getNewsByIdController,
        DeleteNewsController:  deleteNewsController,
		UpdateNewsController : updateNewsController,
	}
}