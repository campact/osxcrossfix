package osxcrossfix

import (
	"crypto/x509"
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

func BenchmarkCertificateParsing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RootCAPool = x509.NewCertPool()
		rootCAs, err := Asset("data/ca.pem")
		if err != nil {
			b.Fatalf("PEM loading failed")
		}
		if ok := RootCAPool.AppendCertsFromPEM([]byte(rootCAs)); !ok {
			b.Fatalf("Parsing failed")
		}
	}
}
