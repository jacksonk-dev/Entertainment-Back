package main

import (
	"H2EBack/packages/movies"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	const appPort string = ":8082"

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello"})
	})
	router.GET("/movies", movies.GetMovies)

	router.Run(appPort)
}
