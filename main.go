package main

import (
	"H2EBack/packages/globals"
	"H2EBack/packages/movies"
	"H2EBack/packages/shows"

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

	router.GET("/movie-genres", movies.GetGenres)
	router.GET("/movies", movies.GetMovies)
	router.GET("/movie-image", movies.GetImage)

	router.GET("/show-genres", shows.GetGenres)
	router.GET("/shows", shows.GetShows)
	router.GET("/show-image", shows.GetImage)

	router.Run(PORT)
}
