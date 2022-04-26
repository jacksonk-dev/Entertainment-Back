package movies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var movies = []Movie{}

func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}
