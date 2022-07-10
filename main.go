package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var quotes = []string{
	"Too too cold!",
	"Beauty beautiful!",
	"I need to pee!",
	"I'm not a dumdum, I am smart smart!"}

func GetQuote() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprint(quotes[rand.Intn(len(quotes))])
}

func GetHomepage(c *gin.Context) {
	s := GetQuote()
	c.JSON(http.StatusOK, gin.H{"message": s})
}

func main() {
	router := gin.Default()
	router.GET("/", GetHomepage)

	router.Run()
}
