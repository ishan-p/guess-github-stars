package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RepositoryInfo struct {
	Id              int    `json:"id"`
	Name            string `json:"full_name"`
	Description     string `json:"description"`
	StargazersCount int    `json:"stargazers_count"`
	Language        string `json:"language"`
	Forks           int    `json:"forks_count"`
}

type SearchResponse struct {
	Items      []RepositoryInfo `json:"items"`
	TotalCount int              `json:"total_count"`
}

func buildQueryParams(language string) string {
	var queryParams string
	if language != "" {
		queryParams = fmt.Sprintf("?q=language:%v&sort=stars&order=desc&per_page=50", language)
	} else {
		queryParams = fmt.Sprintf("?q=is:public&sort=stars&order=desc&per_page=50")
	}
	return queryParams
}

func SearchTrendingRepositories(language string) (SearchResponse, int) {
	url := "https://api.github.com/search/repositories"
	queryParams := buildQueryParams(language)
	request, _ := http.NewRequest(
		"GET",
		url+queryParams,
		nil,
	)
	request.Header.Add("accept", "application/vnd.github.v3+json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err, response.StatusCode)
	}
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	isValid := json.Valid(data)
	if !isValid {
		log.Fatal("Invalid JSON response")
	}

	var searchResult SearchResponse
	_ = json.Unmarshal(data, &searchResult)

	return searchResult, response.StatusCode
}
