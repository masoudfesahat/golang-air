package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World, I'm Masoud Fesahat :)")
	})

	router.GET("/env", func(c *gin.Context) {
		c.String(200, "Hello %s", os.Getenv("NAME"))
	})

	router.Run("0.0.0.0:" + port)
}
