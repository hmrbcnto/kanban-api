package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"hmrbcnto.com/gin-api/config"
	db "hmrbcnto.com/gin-api/infastructure/db"
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
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run()
}
