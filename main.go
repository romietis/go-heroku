package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomepage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func main() {
	router := gin.Default()
	router.GET("/", GetHomepage)

	router.Run("0.0.0.0:8080")
	log.Println("Server started!")
}
