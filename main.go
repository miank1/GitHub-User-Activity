package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// structure of GitHub events
type GitHubEvents struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Commits []struct {
			SHA string `json:"sha"`
		} `json:"commits"`
	} `json:"payload"`
}

func main() {
	var username string
	fmt.Println("github-activity")
	fmt.Scanln(&username) // input github username - eg - freecodecamp

	url := "https://api.github.com/users/" + username + "/events"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var events []GitHubEvents
	if err := json.Unmarshal(body, &events); err != nil { // json to struct
		log.Fatal(err)
	}

	fmt.Println("Output:")
	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			count := len(event.Payload.Commits)
			fmt.Printf("- Pushed %d commit(s) to %s \n", count, event.Repo.Name)
		case "IssueEvent":
			fmt.Printf("- Opened a new issue in %s\n", event.Repo.Name)
		case "WatchedEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		}
	}

}
