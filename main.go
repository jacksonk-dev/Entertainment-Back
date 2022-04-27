package main

import (
	"H2EBack/packages/globals"
	"H2EBack/packages/movies"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var config globals.Config = globals.GetConfig()

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.CLIENT_ORIGIN)
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")

		c.Next()
	}
}

func main() {
	PORT := fmt.Sprintf(":%s", config.PORT)

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello"})
	})
	router.GET("/genres", movies.GetMovieGenres)
	router.GET("/movies", movies.GetMovies)
	router.GET("/movie-image", movies.GetMovieImage)

	router.Run(PORT)
}
