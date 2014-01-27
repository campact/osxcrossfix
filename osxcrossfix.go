package osxcrossfix

// Only include this file if we are compiling for osx
// with cgo disabled.
// +build darwin,!cgo

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
)

func init() {
	transport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		log.Printf("Unexpected underlying type of http.DefaultTransport, aborting certificate injection")
		return
	}

	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	}

	if transport.TLSClientConfig.RootCAs == nil {
		transport.TLSClientConfig.RootCAs = x509.NewCertPool()
	}
	if ok := transport.TLSClientConfig.RootCAs.AppendCertsFromPEM([]byte(rootCAs)); !ok {
		log.Printf("Certificates injection failed")
		return
	}
}
