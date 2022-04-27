package main

import (
	"H2EBack/packages/globals"
	"H2EBack/packages/movies"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config := globals.GetConfig()
	PORT := fmt.Sprintf(":%s", config.PORT)

	router := gin.Default()

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello"})
	})
	router.GET("/api/genres", movies.GetMovieGenres)
	router.GET("/api/movies", movies.GetMovies)
	router.GET("/api/movie-image", movies.GetMovieImage)

	router.Run(PORT)
}
