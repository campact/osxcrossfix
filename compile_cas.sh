#!/bin/bash

rm -rf cas &>/dev/null
mkdir cas; cd cas

# Google
(
	mkdir google; cd google
	curl -O http://pki.google.com/GIAG2.crt
	openssl x509 -inform DES -in GIAG2.crt -out GIAG2.pem -outform PEM -text
)

# VeriSign
(
	mkdir verisign; cd verisign
	curl -O http://www.verisign.com/support/roots.zip
	unzip roots.zip
)

# CaCert
(
	mkdir cacert; cd cacert
	curl -O http://www.cacert.org/certs/root.crt
	curl -O http://www.cacert.org/certs/class3.crt
)

# Build
cd ..
cat > cas.go << EOF
package main

const (
	rootCAs = \`
EOF

find ./cas | grep -E '\.pem$' | while read line; do
	cat "${line}" >> cas.go
	echo "" >> cas.go
done

cat >> cas.go << EOF
	\`)	
EOF

gofmt -w cas.go
