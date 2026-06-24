package main

import (
	"log"
	"net/http"

	"hello-go-http/database"
	"hello-go-http/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using environment variables or defaults")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("database: %v", err)
	}

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello，golang")
	})

	orderHandler := handler.NewOrderHandler(db)
	orderHandler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server: %v", err)
	}
}
