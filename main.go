package main

import (
	config "api/config"
	models "api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Pong"})
}

func getDirtyJoke(c *gin.Context) {
	jokeURL, ok := os.LookupEnv("JOKE_API")
	if !ok {
		fmt.Printf("Error: No Joke API designated. Assuming local run, setting environment variable now \n")
		jokeUrl := "https://v2.jokeapi.dev/joke"
		os.Setenv("JOKE_API", jokeUrl)
	}
	jokeURL += "/Any"
	resp, err := http.Get(jokeURL)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var jokeResp models.JokeResponse
	err = json.Unmarshal(bodyBytes, &jokeResp)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, jokeResp)
}

func getMovieByTitle(c *gin.Context) {
	title := c.Param("title")

	omdbURL, ok := os.LookupEnv("MOVIE_API")
	if !ok {
		fmt.Printf("Error: No Movie API designated. Assuming local run, setting environment variable now \n")
		movieUrl := "http://www.omdbapi.com/?"
		os.Setenv("MOVIE_API", movieUrl)
	}

	config, err := config.InitializeViper("./config/")
	if err != nil {
		log.Fatal("Unable to load config:", err)
	}
	omdbURL += "apikey=" + config.MOVIE_API_KEY + "&t=" + title

	resp, err := http.Get(omdbURL)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var movieResp models.MovieByIDResponse
	err = json.Unmarshal(bodyBytes, &movieResp)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, movieResp)
}

func main() {
	port, ok := os.LookupEnv("HTTP_PORT")

	if !ok {
		fmt.Printf("Error: No HTTP Port designated. Defaulting to environment to gin default (:8080) \n")
		port = ":8080"
		os.Setenv("HTTP_PORT", port)
	}

	router := gin.Default()
	router.GET("/", ping)
	router.GET("/joke", getDirtyJoke)
	router.GET("/movie/:title", getMovieByTitle)
	router.Run(port)
}
