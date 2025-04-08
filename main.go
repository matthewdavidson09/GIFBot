package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PullRequestEvent struct {
	Action      string `json:"action"`
	PullRequest struct {
		Number int `json:"number"`
	} `json:"pull_request"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	token := os.Getenv("GITHUB_TOKEN")

	data, err := os.ReadFile(eventPath)
	if err != nil {
		log.Fatalf("Error reading event: %v", err)
	}

	var event PullRequestEvent
	if err := json.Unmarshal(data, &event); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	gif := GetGif(event.Action)
	if gif == "" {
		log.Printf("No GIF configured for action: %s", event.Action)
		return
	}

	comment := fmt.Sprintf("![gif](%s)", gif)
	err = PostComment(token, event.Repository.FullName, event.PullRequest.Number, comment)
	if err != nil {
		log.Fatalf("Error posting comment: %v", err)
	}
}
