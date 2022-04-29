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

	// TMDB
	router.GET("/trending-movies", movies.GetTrendingMovies)
	router.GET("/trending-shows", shows.GetTrendingShows)
	router.GET("/get-movies", movies.GetTMDBMovies)
	router.GET("/get-shows", shows.GetTMDBShows)
	router.GET("/movie-image", movies.GetImage)
	router.GET("/show-image", shows.GetImage)

	// Trakt
	router.GET("/movie-genres", movies.GetGenres)
	router.GET("/movies", movies.GetMovies)
	router.GET("/show-genres", shows.GetGenres)
	router.GET("/shows", shows.GetShows)

	router.Run(PORT)
}
