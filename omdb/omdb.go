package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Search       []SearchResult `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}

type SearchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func Search(apiKey, title string) (Result, error){
	resp, err := http.Get("http://www.omdbapi.com/?apikey="+ apiKey +"&s=" + title)
	if err != nil {
		return Result{}, fmt.Errorf("failed to make request to omdb: %w", err)
	}
	
	defer resp.Body.Close()

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Result{}, fmt.Errorf("failed to decode response form omdb: %w", err)
	}

	return result, nil
}