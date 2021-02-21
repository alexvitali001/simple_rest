package controllers

import (
	"encoding/json"
	"go_simple_rest/utils"
	"net/http"
	"os"
	"strconv"
)

// Version :  Get the version of the API
var Version = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	type Result struct {
		Version string `json:"version"`
	}
	result := Result{
		Version: os.Getenv("VERSION"),
	}

	utils.Success(w, http.StatusOK, resp, result, "Successfully get the version of API")
}

// MovieList : Get list of movies from TMDB via discovery
var MovieList = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	// Get the movies from TMDB API
	url := os.Getenv("TMDB_API_URL") + "discover/movie?api_key=" + os.Getenv("TMDB_KEY")
	res, err := http.Get(url)
	if err != nil {
		utils.Fail(w, http.StatusInternalServerError, resp, err.Error())
		return
	}

	defer res.Body.Close()

	type Movie struct {
		ID           int    `json:"id"`
		Title        string `json:"title"`
		Description  string `json:"overview"`
		OriginalLink string `json:"poster_path"`
	}

	type Result struct {
		Page   int     `json:"page"`
		Movies []Movie `json:"results"`
	}
	result := Result{}
	json.NewDecoder(res.Body).Decode(&result)

	if len(result.Movies) == 0 {
		utils.Success(w, http.StatusOK, resp, result.Movies, "There is no movies to be discovered.")
		return
	}

	utils.Success(w, http.StatusOK, resp, result, strconv.Itoa(len(result.Movies))+" movie(s) have been discovered.")
}
