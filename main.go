package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"hmrbcnto.com/gin-api/config"
	"hmrbcnto.com/gin-api/handler"
	db "hmrbcnto.com/gin-api/infastructure/db"
	"hmrbcnto.com/gin-api/repository"
	"hmrbcnto.com/gin-api/usecase"
)

func main() {
	// Loading config
	config, err := config.LoadConfig()

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// Creating db connection
	client, err := db.NewConnection(config.DbConfig.MongoURI)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	r := gin.Default()

	// Initializing routes

	// User
	userRepo := repository.NewUserRepo(client)
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler.NewUserHandler(r, userUsecase)

	// Auth
	authUsecase := usecase.NewAuthUsecase(userRepo)
	handler.NewAuthHandler(r, authUsecase)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run()
}
