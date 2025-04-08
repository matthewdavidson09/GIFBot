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
		Number int  `json:"number"`
		Merged bool `json:"merged"`
		Head   struct {
			Repo struct {
				FullName string `json:"full_name"`
			} `json:"repo"`
		} `json:"head"`
		Base struct {
			Repo struct {
				FullName string `json:"full_name"`
			} `json:"repo"`
		} `json:"base"`
	} `json:"pull_request"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	token := os.Getenv("INPUT_GITHUB_TOKEN") // ✅ Note: INPUT_ prefix

	data, err := os.ReadFile(eventPath)
	if err != nil {
		log.Fatalf("Error reading event file: %v", err)
	}

	var event PullRequestEvent
	if err := json.Unmarshal(data, &event); err != nil {
		log.Fatalf("Error parsing event JSON: %v", err)
	}

	// 🔁 Normalize action if needed
	action := event.Action
	if action == "synchronize" || action == "opened" {
		action = "opened_or_synchronize"
	} else if action == "closed" && event.PullRequest.Merged {
		action = "merged"
	}

	gif := GetGif(action)
	if gif == "" {
		log.Printf("No GIF configured for action: %s", action)
		return
	}

	comment := fmt.Sprintf("![gif](%s)", gif)
	if err := PostComment(token, event.Repository.FullName, event.PullRequest.Number, comment); err != nil {
		log.Fatalf("Error posting comment: %v", err)
	}

	if IsForkPR(event) {
		log.Println("Skipping comment: forked PR detected.")
		return
	}
}
