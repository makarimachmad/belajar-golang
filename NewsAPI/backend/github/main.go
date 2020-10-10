package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RepositoryOwner struct {
	Username        string          `json:"login"`
	Type            string          `json:"type"`
}

type Repository struct {
	ID              int             `json:"id"`
	Owner           RepositoryOwner `json:"owner"`
	Name            string          `json:"name"`
	ForksCount      int             `json:"forks_count"`
	StargazersCount int             `json:"stargazers_count"`
	WatcherCounts   int             `json:"watchers_count"`
}

func main() {
        // 1
	res, err := http.Get("https://api.github.com/users/99ridho/repos")
	if err != nil {
		panic(err.Error())
	}
        
        // 2
        body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	structJson := make([]Repository, 0)
	json.Unmarshal(body, &structJson)

	idx := 0
	for idx < len(structJson) {
		fmt.Println(structJson[idx])
		idx++
	}
}