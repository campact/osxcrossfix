#!/bin/bash

function crt2pem() {
	find . | grep -E '\.crt$' | while read line; do
		openssl x509 -inform DES -in $line -out ${line}.pem -outform PEM -text
	done
}

rm -rf cas &>/dev/null
mkdir cas; cd cas

# Google
(
	mkdir google; cd google
	curl -O http://pki.google.com/GIAG2.crt
	crt2pem
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

# DigiCert
(
	mkdir digicert; cd digicert
	curl -O https://www.digicert.com/CACerts/DigiCertHighAssuranceEVRootCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertAssuredIDRootCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertGlobalRootCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertAssuredIDRootG2.crt
	curl -O https://www.digicert.com/CACerts/DigiCertAssuredIDRootG3.crt
	curl -O https://www.digicert.com/CACerts/DigiCertGlobalRootG2.crt
	curl -O https://www.digicert.com/CACerts/DigiCertGlobalRootG3.crt
	curl -O https://www.digicert.com/CACerts/DigiCertTrustedRootG4.crt
	curl -O https://www.digicert.com/CACerts/DigiCertSecureServerCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertSHA2SecureServerCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertECCSecureServerCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertHighAssuranceEVCA-1.crt
	curl -O https://www.digicert.com/CACerts/DigiCertHighAssuranceCA-3.crt
	curl -O https://www.digicert.com/CACerts/DigiCertAssuredIDCA-1.crt
	curl -O https://www.digicert.com/CACerts/DigiCertGlobalCA-1.crt
	curl -O https://www.digicert.com/CACerts/DigiCertAssuredIDCodeSigningCA-1.crt
	curl -O https://www.digicert.com/CACerts/DigiCertHighAssuranceCodeSigningCA-1.crt
	curl -O https://www.digicert.com/CACerts/DigiCertEVCodeSigningCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertEVCodeSigningCA-SHA2.crt
	curl -O https://www.digicert.com/CACerts/DigiCertSHA2HighAssuranceServerCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertSHA2ExtendedValidationServerCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertSHA2AssuredIDCodeSigningCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertSHA2HighAssuranceCodeSigningCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertSHA2AssuredIDCA.crt
	curl -O https://www.digicert.com/CACerts/DigiCertDocumentSigningCA.crt
	crt2pem
)

# GlobalSign
(
	mkdir globalsign; cd globalsign
	curl -O http://secure.globalsign.net/cacert/primserver.crt
	curl -O http://secure.globalsign.net/cacert/ServerSign.crt
	crt2pem
)

# Build
cd ..
cat > cas.go << EOF
package osxcrossfix

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
