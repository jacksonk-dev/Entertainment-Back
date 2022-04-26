package movies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Genre struct {
	Label string `json:"name"`
	Value string `json:"slug"`
}

type TraktIDs struct {
	Trakt int    `json:"trakt"`
	Slug  string `json:"slug"`
	IMDB  string `json:"imdb"`
	TMDB  int    `json:"tmdb"`
}

type ActualMovie struct {
	Title string   `json:"title"`
	Year  string   `json:"year"`
	IDs   TraktIDs `json:"ids"`
}
type Movie struct {
	Watchers int         `json:"watchers"`
	Movie    ActualMovie `json:"movie"`
}

var genres []Genre
var movies []Movie

func addGenericHeadersToRequest(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", "97d2684afdbb32dc5306041308ad7b334c1616a124c77afb58e73eec6dc02342")
}

func addMovieFetchParams(params *url.Values, genres, page string) {
	params.Add("limit", "30")
	params.Add("page", page)
	params.Add("genres", genres)
	params.Add("extended", "full")
}

func GetMovieGenres(c *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.trakt.tv/genres/movies", nil)
	addGenericHeadersToRequest(req)

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
	paramsMap := c.Request.URL.Query()

	base, err := url.Parse("https://api.trakt.tv/movies/" + paramsMap.Get("subLink"))
	if err != nil {
		return
	}
	params := url.Values{}
	addMovieFetchParams(&params, paramsMap.Get("genres"), paramsMap.Get("page"))

	base.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", base.String(), nil)
	addGenericHeadersToRequest(req)

	fmt.Println(base.String(), req)

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
	json.Unmarshal(bodyBytes, &movies)

	c.IndentedJSON(http.StatusOK, movies)
}
