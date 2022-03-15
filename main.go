package main

import (
	config "api/config"
	models "api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getDirtyJoke(c *gin.Context) {
	jokeURL, ok := os.LookupEnv("JOKE_BASE_URL")
	if !ok {
		fmt.Printf("Error: %v", ok)
	}
	jokeURL += "Any"
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
	omdbURL, ok := os.LookupEnv("OMDB_BASE_URL")
	if !ok {
		fmt.Printf("Error: %v", ok)
	}

	omdbURL += "&t=" + title

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
	cliConfig := os.Args
	config.ConfigEnv(cliConfig)
	hostURL, ok := os.LookupEnv("HTTP_URL")
	if !ok {
		fmt.Printf("Error: %v", ok)
	}

	router := gin.Default()
	router.GET("/joke", getDirtyJoke)
	router.GET("/movie/:title", getMovieByTitle)
	router.Run(hostURL)
}
