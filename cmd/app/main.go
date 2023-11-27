package main

import (
	"first_project/config"
	"first_project/internal/domain/author/usecase"
	"first_project/internal/infrastructure/controller/http"
	"first_project/internal/infrastructure/repository"
	"first_project/pkg/postgres"
	"first_project/pkg/server"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewConnectToPostgress(conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	handler := gin.Default()

	//Инициализация Repository
	authorRepo := repository.NewAuthorRepository(db)
	//Инициализация Usecase
	authorUC := usecase.NewAuthorUseCase(authorRepo)
	//Инициализация Controller
	http.NewAutorConthroller(handler, authorUC)
	//http.NewController(handler)

	if err = server.NewHttpServer(conf.Http, handler).Start(); err != nil {
		log.Fatal(err)
	}

}
