package osxcrossfix

import (
	"testing"
)

func TestPopularPages() {
	websites := []string{"facebook.com", "google.com", "plus.google.com", "twitter.com"}
	for _, website := range websites {
		resp, err := http.Get("https://" + website)
		if err != nil {
			log.Fatalf("Could not contact facebook: %s", err)
		}
		log.Printf("Success: %d", resp.StatusCode)
	}
}
