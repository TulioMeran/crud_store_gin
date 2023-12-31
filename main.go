package main

import (
	"crud_store_gin/config"
	"crud_store_gin/controller"
	_ "crud_store_gin/docs"
	"crud_store_gin/models"
	"crud_store_gin/repository"
	"crud_store_gin/router"
	"crud_store_gin/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// @title 	Tag Service API
// @version 1.0
// @description	A Tag service API in Go using Gin framework
// @host localhost:8080
// @basePath /api
func main() {

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&models.Tags{})

	//Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	//Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	//Controller
	tagsController := controller.NewTagsController(tagsService)

	//Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Error launching the server: " + err.Error())
	}
}
