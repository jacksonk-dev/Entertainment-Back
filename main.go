package main

import (
	"EntertainmentBack/packages/globals"
	"EntertainmentBack/packages/movies"
	"EntertainmentBack/packages/shows"
	"os"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	var clientOrigin string = os.Getenv("CLIENT_ORIGIN")

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", clientOrigin)
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")

		c.Next()
	}
}

func main() {
	var port string = os.Getenv("PORT")
	if port == "" {
		globals.SetConfig()
		port = os.Getenv("PORT")
	}

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

	router.Run(fmt.Sprintf(":%s", port))
}
