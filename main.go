package main

import (
	"H2EBack/packages/movies"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	const appPort string = ":8082"

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello"})
	})
	router.GET("/api/genres", movies.GetMovieGenres)
	router.GET("/api/movies", movies.GetMovies)

	router.Run(appPort)
}
