package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostComment(token, repo string, issueNumber int, body string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d/comments", repo, issueNumber)

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
		return fmt.Errorf("GitHub API error: %s", resp.Status)
	}
	return nil
}
