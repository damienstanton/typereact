package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var port = os.Getenv("$PORT")

func main() {
	router := gin.Default()
	router.Static("/", "../")
	if port == "" {
		port = "8000"
	}
	log.Printf("Listening for *NEW* changes on port %v\n", port)
	router.Run(":" + port)
}
