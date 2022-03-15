package config

import (
	"os"
)

func ConfigEnv(cliConfig []string) {
	os.Setenv("JOKE_BASE_URL", "https://v2.jokeapi.dev/joke/")
	os.Setenv("HTTP_URL", "localhost:8080")
	os.Setenv("OMDB_BASE_URL", "http://www.omdbapi.com/?apikey=89174c9c")
}
