package movies

import (
	"os"

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

type TMDBMovie struct {
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	PosterPath    string `json:"poster_path"`
}

type TMDBResults struct {
	Page    int         `json:"page"`
	Results []TMDBMovie `json:"results"`
}
type ActualMovie struct {
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	IDs      TraktIDs `json:"ids"`
	Tagline  string   `json:"tagline"`
	Overview string   `json:"overview"`
	Released string   `json:"released"`
	Runtime  int      `json:"runtime"`
	Trailer  string   `json:"trailer"`
	Status   string   `json:"status"`
	Rating   float64  `json:"rating"`
}
type Movie struct {
	Watchers int         `json:"watchers"`
	Movie    ActualMovie `json:"movie"`
}

type SingleImageData struct {
	PosterPath string `json:"poster_path"`
}

type ImageData struct {
	MovieResults []SingleImageData `json:"movie_results"`
}

var genres []Genre
var movies []Movie
var imageData ImageData

func addTraktHeadersToRequest(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", os.Getenv("TRAKT_API_KEY"))
}

func addMovieFetchParams(params *url.Values, genres, page string) {
	params.Add("limit", "30")
	params.Add("page", page)
	params.Add("genres", genres)
	params.Add("extended", "full")
}

func GetGenres(c *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.trakt.tv/genres/movies", nil)
	addTraktHeadersToRequest(req)

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

func GetTrendingMovies(c *gin.Context) {
	client := &http.Client{}

	var url string = fmt.Sprintf("https://api.themoviedb.org/3/trending/movie/day?api_key=%s", os.Getenv("TMDB_API_KEY"))
	req, err := http.NewRequest("GET", url, nil)

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

	var fetchResults TMDBResults
	json.Unmarshal(bodyBytes, &fetchResults)

	c.IndentedJSON(http.StatusOK, fetchResults.Results)
}

func GetTMDBMovies(c *gin.Context) {
	client := &http.Client{}
	paramsMap := c.Request.URL.Query()

	var url string = fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s&language=en-US&page=%s", paramsMap.Get("list"), os.Getenv("TMDB_API_KEY"), paramsMap.Get("page"))
	req, err := http.NewRequest("GET", url, nil)

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

	var fetchResults TMDBResults
	json.Unmarshal(bodyBytes, &fetchResults)

	c.IndentedJSON(http.StatusOK, fetchResults.Results)
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
	addTraktHeadersToRequest(req)

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

func GetImage(c *gin.Context) {
	client := &http.Client{}
	paramsMap := c.Request.URL.Query()

	var url string = fmt.Sprintf("https://api.themoviedb.org/3/find/%s?api_key=%s&language=en-US&external_source=imdb_id", paramsMap.Get("id"), os.Getenv("TMDB_API_KEY"))
	req, err := http.NewRequest("GET", url, nil)

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

	json.Unmarshal(bodyBytes, &imageData)

	c.IndentedJSON(http.StatusOK, imageData)
}
