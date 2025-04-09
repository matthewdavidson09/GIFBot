package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// PostComment posts a comment to a GitHub issue.
func PostComment(token, repo string, issueNumber int, body string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d/comments", repo, issueNumber)
	log.Printf("DEBUG: Repo=%s Issue#=%d Body=%s\n", repo, issueNumber, body)

	payload := map[string]string{"body": body}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GitHub API error: %s\nResponse Body: %s", resp.Status, string(bodyBytes))
	}
	return nil
}

// IsForkPR determines if the pull request is from a fork.
func IsForkPR(event PullRequestEvent) bool {
	return event.PullRequest.Head.Repo.FullName != event.PullRequest.Base.Repo.FullName
}

func getEventKey(event PullRequestEvent, config map[string][]string) string {
	// The event's action, assumed to be lowercase from GitHub, but normalized here.
	key := strings.ToLower(event.Action)
	if key == "closed" && event.PullRequest.Merged {
		// If the JSON config includes a "merged" key, use that.
		if _, ok := config["merged"]; ok {
			return "merged"
		}
	}
	return key
}
