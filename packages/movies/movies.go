package movies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Genre struct {
	Label string `json:"name"`
	Value string `json:"slug"`
}

type Movie struct {
}

var genres []Genre

func addHeadersToRequest(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", "97d2684afdbb32dc5306041308ad7b334c1616a124c77afb58e73eec6dc02342")
}

func GetMovieGenres(c *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.trakt.tv/genres/movies", nil)
	addHeadersToRequest(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	json.Unmarshal(bodyBytes, &genres)

	c.IndentedJSON(http.StatusOK, genres)
}

func GetMovies(c *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.trakt.tv/genres/movies", nil)
	addHeadersToRequest(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	json.Unmarshal(bodyBytes, &genres)

	c.IndentedJSON(http.StatusOK, genres)
}
