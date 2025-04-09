package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GetGifForEvent returns a random GIF URL for the given GitHub event action.
func GetGifForEvent(eventKey string, config map[string][]string) string {
	// Normalize eventKey (assuming keys are lower-case in our config)
	eventKey = strings.ToLower(eventKey)
	urls, ok := config[eventKey]
	if !ok || len(urls) == 0 {
		return ""
	}
	return urls[rnd.Intn(len(urls))]
}

func loadGifMapping() map[string][]string {
	data, err := os.ReadFile("gif.json")
	if err != nil {
		log.Printf("gif.json not found: %v", err)
		return map[string][]string{}
	}

	var mapping map[string][]string
	if err := json.Unmarshal(data, &mapping); err != nil {
		log.Fatalf("Error parsing gif.json: %v", err)
	}

	// Optionally normalize keys. If you want no transformation, remove this loop.
	normalized := make(map[string][]string)
	for key, urls := range mapping {
		normalized[strings.ToLower(key)] = urls
	}
	return normalized
}
