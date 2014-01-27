package osxcrossfix

import (
	"net/http"
	"testing"
)

func TestPopularPages(t *testing.T) {
	InjectCertificates()
	websites := []string{"facebook.com", "google.com", "plus.google.com", "twitter.com", "linkedin.com"}
	for _, website := range websites {
		_, err := http.Get("https://" + website)
		if err != nil {
			t.Fatalf("Could not request %s: %s", website, err)
		}
	}
}
