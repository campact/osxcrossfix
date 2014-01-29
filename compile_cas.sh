#!/bin/bash

curl http://curl.haxx.se/ca/cacert.pem > ca.pem

cat > cas.go << EOF
package osxcrossfix

const (
	rootCAs = \`
EOF

cat ca.pem >> cas.go

cat >> cas.go << EOF
	\`)	
EOF

gofmt -w cas.go
