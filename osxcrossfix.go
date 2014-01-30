// +build darwin,!cgo

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
// It is called automatically when compiling for darwin with cgo disabled.
func InjectCertificates() {
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

func parseCertificates() {
	RootCAPool = x509.NewCertPool()
	rootCAs, err := Asset("data/ca.pem")
	if err != nil {
		log.Print("Certificate data could not be loaded")
		return
	}
	if ok := RootCAPool.AppendCertsFromPEM([]byte(rootCAs)); !ok {
		log.Printf("Certificate injection failed")
		return
	}
}

func init() {
	parseCertificates()
	InjectCertificates()
}
