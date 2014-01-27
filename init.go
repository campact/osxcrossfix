// Only include this file if we are compiling for osx
// with cgo disabled.
// +build darwin,!cgo

package osxcrossfix

import (
	"log"
)

func init() {
	log.Printf("Injecting certificates...")
	InjectCertificates()
}
