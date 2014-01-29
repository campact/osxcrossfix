package osxcrossfix

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
)

// RootCAPool contains the certificate pool of bundled certificates
// after ParseCertificates has been called.
var RootCAPool *x509.CertPool

// InjectCertificates injects RootCAPool into the http.DefaultTransport (used by
// http.DefaultClient), provided it hasn't been tampered with.
// Will call ParseCertificates if necessary.
// It is called automatically when compiling for darwin with cgo disabled.
func InjectCertificates() {
	if RootCAPool == nil {
		ParseCertificates()
	}
	transport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		log.Printf("Unexpected underlying type of http.DefaultTransport, aborting certificate injection")
		return
	}

	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	}

	if transport.TLSClientConfig.RootCAs != nil {
		log.Printf("RootCAs is not nil, cannot inject certificates")
		return
	}
	transport.TLSClientConfig.RootCAs = RootCAPool
}

// ParseCertificates will parse the bundled PEM file
// into the x509.CertPool structure.
func ParseCertificates() {
	RootCAPool = x509.NewCertPool()
	if ok := RootCAPool.AppendCertsFromPEM([]byte(rootCAs)); !ok {
		log.Printf("Certificates injection failed")
		return
	}
}
