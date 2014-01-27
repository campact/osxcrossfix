package osxcrossfix

const (
	rootCAs = `
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 146025 (0x23a69)
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=GeoTrust Inc., CN=GeoTrust Global CA
        Validity
            Not Before: Apr  5 15:15:55 2013 GMT
            Not After : Apr  4 15:15:55 2015 GMT
        Subject: C=US, O=Google Inc, CN=Google Internet Authority G2
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:9c:2a:04:77:5c:d8:50:91:3a:06:a3:82:e0:d8:
                    50:48:bc:89:3f:f1:19:70:1a:88:46:7e:e0:8f:c5:
                    f1:89:ce:21:ee:5a:fe:61:0d:b7:32:44:89:a0:74:
                    0b:53:4f:55:a4:ce:82:62:95:ee:eb:59:5f:c6:e1:
                    05:80:12:c4:5e:94:3f:bc:5b:48:38:f4:53:f7:24:
                    e6:fb:91:e9:15:c4:cf:f4:53:0d:f4:4a:fc:9f:54:
                    de:7d:be:a0:6b:6f:87:c0:d0:50:1f:28:30:03:40:
                    da:08:73:51:6c:7f:ff:3a:3c:a7:37:06:8e:bd:4b:
                    11:04:eb:7d:24:de:e6:f9:fc:31:71:fb:94:d5:60:
                    f3:2e:4a:af:42:d2:cb:ea:c4:6a:1a:b2:cc:53:dd:
                    15:4b:8b:1f:c8:19:61:1f:cd:9d:a8:3e:63:2b:84:
                    35:69:65:84:c8:19:c5:46:22:f8:53:95:be:e3:80:
                    4a:10:c6:2a:ec:ba:97:20:11:c7:39:99:10:04:a0:
                    f0:61:7a:95:25:8c:4e:52:75:e2:b6:ed:08:ca:14:
                    fc:ce:22:6a:b3:4e:cf:46:03:97:97:03:7e:c0:b1:
                    de:7b:af:45:33:cf:ba:3e:71:b7:de:f4:25:25:c2:
                    0d:35:89:9d:9d:fb:0e:11:79:89:1e:37:c5:af:8e:
                    72:69
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Authority Key Identifier: 
                keyid:C0:7A:98:68:8D:89:FB:AB:05:64:0C:11:7D:AA:7D:65:B8:CA:CC:4E

            X509v3 Subject Key Identifier: 
                4A:DD:06:16:1B:BC:F6:68:B5:76:F5:81:B6:BB:62:1A:BA:5A:81:2F
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl.geotrust.com/crls/gtglobal.crl

            Authority Information Access: 
                OCSP - URI:http://gtglobal-ocsp.geotrust.com

            X509v3 Certificate Policies: 
                Policy: 1.3.6.1.4.1.11129.2.5.1

    Signature Algorithm: sha1WithRSAEncryption
         36:d7:06:80:11:27:ad:2a:14:9b:38:77:b3:23:a0:75:58:bb:
         b1:7e:83:42:ba:72:da:1e:d8:8e:36:06:97:e0:f0:95:3b:37:
         fd:1b:42:58:fe:22:c8:6b:bd:38:5e:d1:3b:25:6e:12:eb:5e:
         67:76:46:40:90:da:14:c8:78:0d:ed:95:66:da:8e:86:6f:80:
         a1:ba:56:32:95:86:dc:dc:6a:ca:04:8c:5b:7f:f6:bf:cc:6f:
         85:03:58:c3:68:51:13:cd:fd:c8:f7:79:3d:99:35:f0:56:a3:
         bd:e0:59:ed:4f:44:09:a3:9e:38:7a:f6:46:d1:1d:12:9d:4f:
         be:d0:40:fc:55:fe:06:5e:3c:da:1c:56:bd:96:51:7b:6f:57:
         2a:db:a2:aa:96:dc:8c:74:c2:95:be:f0:6e:95:13:ff:17:f0:
         3c:ac:b2:10:8d:cc:73:fb:e8:8f:02:c6:f0:fb:33:b3:95:3b:
         e3:c2:cb:68:58:73:db:a8:24:62:3b:06:35:9d:0d:a9:33:bd:
         78:03:90:2e:4c:78:5d:50:3a:81:d4:ee:a0:c8:70:38:dc:b2:
         f9:67:fa:87:40:5d:61:c0:51:8f:6b:83:6b:cd:05:3a:ca:e1:
         a7:05:78:fc:ca:da:94:d0:2c:08:3d:7e:16:79:c8:a0:50:20:
         24:54:33:71
-----BEGIN CERTIFICATE-----
MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
/iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIICPTCCAaYCEQDNun9W8N/kvFT+IqyzcqpVMA0GCSqGSIb3DQEBAgUAMF8xCzAJ
BgNVBAYTAlVTMRcwFQYDVQQKEw5WZXJpU2lnbiwgSW5jLjE3MDUGA1UECxMuQ2xh
c3MgMSBQdWJsaWMgUHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTAeFw05
NjAxMjkwMDAwMDBaFw0yODA4MDEyMzU5NTlaMF8xCzAJBgNVBAYTAlVTMRcwFQYD
VQQKEw5WZXJpU2lnbiwgSW5jLjE3MDUGA1UECxMuQ2xhc3MgMSBQdWJsaWMgUHJp
bWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTCBnzANBgkqhkiG9w0BAQEFAAOB
jQAwgYkCgYEA5Rm/baNWYS2ZSHH2Z965jeu3noaACpEO+jglr0aIguVzqKCbJF0N
H8xlbgyw0FaEGIeaBpsQoXPftFg5a27B9hXVqKg/qhIGjTGsf7A01480Z4gJzRQR
4k5FVmkfeAKA2txHkSm7NsljXMXg1y2He6G3MrB7MLoqLzGq7qNn2tsCAwEAATAN
BgkqhkiG9w0BAQIFAAOBgQBMP7iLxmjf7kMzDl3ppssHhE16M/+SG/Q2rdiVIjZo
EWx8QszznC7EBz8UsA9P/5CSdvnivErpj82ggAr3xSnxgiJduLHdgSOjeyUVRjB5
FvjqBUuUfx3CHMjjt/QQQDwTw18fU+hI5Ia0e6E1sHslurjTjqs/OJ0ANACY89Fx
lA==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIICPDCCAaUCEC0b/EoXjaOR6+f/9YtFvgswDQYJKoZIhvcNAQECBQAwXzELMAkG
A1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMTcwNQYDVQQLEy5DbGFz
cyAyIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MB4XDTk2
MDEyOTAwMDAwMFoXDTI4MDgwMTIzNTk1OVowXzELMAkGA1UEBhMCVVMxFzAVBgNV
BAoTDlZlcmlTaWduLCBJbmMuMTcwNQYDVQQLEy5DbGFzcyAyIFB1YmxpYyBQcmlt
YXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MIGfMA0GCSqGSIb3DQEBAQUAA4GN
ADCBiQKBgQC2WoujDWojg4BrzzmH9CETMwZMJaLtVRKXxaeAufqDwSCg+i8VDXyh
YGt+eSz6Bg86rvYbb7HS/y8oUl+DfUvEerf4Zh+AVPy3wo5ZShRXRtGak75BkQO7
FYCTXOvnzAhsPz6zSvz/S2wj1VCCJkQZjiPDceoZJEcEnnW/yKYAHwIDAQABMA0G
CSqGSIb3DQEBAgUAA4GBAIobK/o5wXTXXtgZZKJYSi034DNHD6zt96rbHuSLBlxg
J8pFUs4W7z8GZOeUaHxgMxURaa+dYo2jA1Rrpr7l7gUYYAS/QoD90KioHgE796Nc
r6Pc5iaAIzy4RHT3Cq5Ji2F4zCS/iIqnDupzGUH9TQPwiNHleI2lKk/2lw0Xd8rY
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIICPDCCAaUCEDyRMcsf9tAbDpq40ES/Er4wDQYJKoZIhvcNAQEFBQAwXzELMAkG
A1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMTcwNQYDVQQLEy5DbGFz
cyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MB4XDTk2
MDEyOTAwMDAwMFoXDTI4MDgwMjIzNTk1OVowXzELMAkGA1UEBhMCVVMxFzAVBgNV
BAoTDlZlcmlTaWduLCBJbmMuMTcwNQYDVQQLEy5DbGFzcyAzIFB1YmxpYyBQcmlt
YXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MIGfMA0GCSqGSIb3DQEBAQUAA4GN
ADCBiQKBgQDJXFme8huKARS0EN8EQNvjV69qRUCPhAwL0TPZ2RHP7gJYHyX3KqhE
BarsAx94f56TuZoAqiN91qyFomNFx3InzPRMxnVx0jnvT0Lwdd8KkMaOIG+YD/is
I19wKTakyYbnsZogy1Olhec9vn2a/iRFM9x2Fe0PonFkTGUugWhFpwIDAQABMA0G
CSqGSIb3DQEBBQUAA4GBABByUqkFFBkyCEHwxWsKzH4PIRnN5GfcX6kb5sroc50i
2JhucwNhkcV8sEVAbkSdjbCxlnRhLQ2pRdKkkirWmnWXbj9T/UWZYB2oK0z5XqcJ
2HUw19JlYD1n1khVdWk/kfVIC0dpImmClr7JyDiGSnoscxlIaU5rfGW/D/xwzoiQ
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDAjCCAmsCEEzH6qqYPnHTkxD4PTqJkZIwDQYJKoZIhvcNAQEFBQAwgcExCzAJ
BgNVBAYTAlVTMRcwFQYDVQQKEw5WZXJpU2lnbiwgSW5jLjE8MDoGA1UECxMzQ2xh
c3MgMSBQdWJsaWMgUHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eSAtIEcy
MTowOAYDVQQLEzEoYykgMTk5OCBWZXJpU2lnbiwgSW5jLiAtIEZvciBhdXRob3Jp
emVkIHVzZSBvbmx5MR8wHQYDVQQLExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMB4X
DTk4MDUxODAwMDAwMFoXDTI4MDgwMTIzNTk1OVowgcExCzAJBgNVBAYTAlVTMRcw
FQYDVQQKEw5WZXJpU2lnbiwgSW5jLjE8MDoGA1UECxMzQ2xhc3MgMSBQdWJsaWMg
UHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eSAtIEcyMTowOAYDVQQLEzEo
YykgMTk5OCBWZXJpU2lnbiwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5
MR8wHQYDVQQLExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMIGfMA0GCSqGSIb3DQEB
AQUAA4GNADCBiQKBgQCq0Lq+Fi24g9TK0g+8djHKlNgdk4xWArzZbxpvUjZudVYK
VdPfQ4chEWWKfo+9Id5rMj8bhDSVBZ1BNeuS65bdqlk/AVNtmU/t5eIqWpDBucSm
Fc/IReumXY6cPvBkJHalzasab7bYe1FhbqZ/h8jit+U03EGI6glAvnOSPWvndQID
AQABMA0GCSqGSIb3DQEBBQUAA4GBAKlPww3HZ74sy9mozS11534Vnjty637rXC0J
h9ZrbWB85a7FkCMMXErQr7Fd88e2CtvgFZMN3QO8x3aKtd1Pw5sTdbgBwObJW2ul
uIncrKTdcu1OofdPvAbT6shkdHvClUGcZXNY8ZCaPGqxmMnEh7zPRW1F4m4iP/68
DzFc6PLZ
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDAzCCAmwCEQC5L2DMiJ+hekYJuFtwbIqvMA0GCSqGSIb3DQEBBQUAMIHBMQsw
CQYDVQQGEwJVUzEXMBUGA1UEChMOVmVyaVNpZ24sIEluYy4xPDA6BgNVBAsTM0Ns
YXNzIDIgUHVibGljIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkgLSBH
MjE6MDgGA1UECxMxKGMpIDE5OTggVmVyaVNpZ24sIEluYy4gLSBGb3IgYXV0aG9y
aXplZCB1c2Ugb25seTEfMB0GA1UECxMWVmVyaVNpZ24gVHJ1c3QgTmV0d29yazAe
Fw05ODA1MTgwMDAwMDBaFw0yODA4MDEyMzU5NTlaMIHBMQswCQYDVQQGEwJVUzEX
MBUGA1UEChMOVmVyaVNpZ24sIEluYy4xPDA6BgNVBAsTM0NsYXNzIDIgUHVibGlj
IFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkgLSBHMjE6MDgGA1UECxMx
KGMpIDE5OTggVmVyaVNpZ24sIEluYy4gLSBGb3IgYXV0aG9yaXplZCB1c2Ugb25s
eTEfMB0GA1UECxMWVmVyaVNpZ24gVHJ1c3QgTmV0d29yazCBnzANBgkqhkiG9w0B
AQEFAAOBjQAwgYkCgYEAp4gBIXQs5xoD8JjhlzwPIQjxnNuX6Zr8wgQGE75fUsjM
HiwSViy4AWkszJkfrbCWrnkE8hM5wXuYuggs6MKEEyyqaekJ9MepAqRCwiNPStjw
DqL7MWzJ5m+ZJwf15vRMeJ5t60aG+rmGyVTyssSv1EYcWskVMP8NbPUtDm3Of3cC
AwEAATANBgkqhkiG9w0BAQUFAAOBgQByLvl/0fFx+8Se9sVeUYpAmLho+Jscg9ji
nb3/7aHmZuovCfTK1+qlK5X2JGCGTUQug6XELaDTrnhpb3LabK4I8GOSN+a7xDAX
rXfMSTWqz9iP0b63GJZHc2pUIjRkLbYWm1lbtFFZOrMLFPQS32eg9K0yZF6xRnIn
jBJ7xUS0rg==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDAjCCAmsCEH3Z/gfPqB63EHln+6eJNMYwDQYJKoZIhvcNAQEFBQAwgcExCzAJ
BgNVBAYTAlVTMRcwFQYDVQQKEw5WZXJpU2lnbiwgSW5jLjE8MDoGA1UECxMzQ2xh
c3MgMyBQdWJsaWMgUHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eSAtIEcy
MTowOAYDVQQLEzEoYykgMTk5OCBWZXJpU2lnbiwgSW5jLiAtIEZvciBhdXRob3Jp
emVkIHVzZSBvbmx5MR8wHQYDVQQLExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMB4X
DTk4MDUxODAwMDAwMFoXDTI4MDgwMTIzNTk1OVowgcExCzAJBgNVBAYTAlVTMRcw
FQYDVQQKEw5WZXJpU2lnbiwgSW5jLjE8MDoGA1UECxMzQ2xhc3MgMyBQdWJsaWMg
UHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eSAtIEcyMTowOAYDVQQLEzEo
YykgMTk5OCBWZXJpU2lnbiwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5
MR8wHQYDVQQLExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMIGfMA0GCSqGSIb3DQEB
AQUAA4GNADCBiQKBgQDMXtERXVxp0KvTuWpMmR9ZmDCOFoUgRm1HP9SFIIThbbP4
pO0M8RcPO/mn+SXXwc+EY/J8Y8+iR/LGWzOOZEAEaMGAuWQcRXfH2G71lSk8UOg0
13gfqLptQ5GVj0VXXn7F+8qkBOvqlzdUMG+7AUcyM83cV5tkaWH4mx0ciU9cZwID
AQABMA0GCSqGSIb3DQEBBQUAA4GBAFFNzb5cy5gZnBWyATl4Lk0PZ3BwmcYQWpSk
U01UbSuvDV1Ai2TT1+7eVmGSX6bEHRBhNtMsJzzoKQm5EWR0zLVznxxIqbxhAe7i
F6YM40AIOw7n60RzKprxaZLvcRTDOaxxp5EJb+RxBrO6WVcmeQD2+A2iMzAo1KpY
oJ2daZH9
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDAjCCAmsCEDKIjprS9esTR/h/xCA3JfgwDQYJKoZIhvcNAQEFBQAwgcExCzAJ
BgNVBAYTAlVTMRcwFQYDVQQKEw5WZXJpU2lnbiwgSW5jLjE8MDoGA1UECxMzQ2xh
c3MgNCBQdWJsaWMgUHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eSAtIEcy
MTowOAYDVQQLEzEoYykgMTk5OCBWZXJpU2lnbiwgSW5jLiAtIEZvciBhdXRob3Jp
emVkIHVzZSBvbmx5MR8wHQYDVQQLExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMB4X
DTk4MDUxODAwMDAwMFoXDTI4MDgwMTIzNTk1OVowgcExCzAJBgNVBAYTAlVTMRcw
FQYDVQQKEw5WZXJpU2lnbiwgSW5jLjE8MDoGA1UECxMzQ2xhc3MgNCBQdWJsaWMg
UHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eSAtIEcyMTowOAYDVQQLEzEo
YykgMTk5OCBWZXJpU2lnbiwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5
MR8wHQYDVQQLExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMIGfMA0GCSqGSIb3DQEB
AQUAA4GNADCBiQKBgQC68OTP+cSuhVS5B1f5j8V/aBH4xBewRNzjMHPVKmIquNDM
HO0oW369atyzkSTKQWI8/AIBvxwWMZQFl3Zuoq29YRdsTjCG8FE3KlDHqGKB3FtK
qsGgtG7rL+VXxbErQHDbWk2hjh+9Ax/YA9SPTJlxvOKCzFjomDqG04Y48wApHwID
AQABMA0GCSqGSIb3DQEBBQUAA4GBAIWMEsGnuVAVess+rLhDityq3RS6iYF+ATwj
cSGIL4LcY/oCRaxFWdcqWERbt5+BO5JoPeI3JPV7bI92NZYJqFmduc4jq3TWg/0y
cyfYaT5DdPauxYma51N86Xv2S/PBZYPejYqcPIiNOVn8qj8ijaHBZlCBckztImRP
T8qAkbYp
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEGjCCAwICEQCLW3VWhFSFCwDPrzhIzrGkMA0GCSqGSIb3DQEBBQUAMIHKMQsw
CQYDVQQGEwJVUzEXMBUGA1UEChMOVmVyaVNpZ24sIEluYy4xHzAdBgNVBAsTFlZl
cmlTaWduIFRydXN0IE5ldHdvcmsxOjA4BgNVBAsTMShjKSAxOTk5IFZlcmlTaWdu
LCBJbmMuIC0gRm9yIGF1dGhvcml6ZWQgdXNlIG9ubHkxRTBDBgNVBAMTPFZlcmlT
aWduIENsYXNzIDEgUHVibGljIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3Jp
dHkgLSBHMzAeFw05OTEwMDEwMDAwMDBaFw0zNjA3MTYyMzU5NTlaMIHKMQswCQYD
VQQGEwJVUzEXMBUGA1UEChMOVmVyaVNpZ24sIEluYy4xHzAdBgNVBAsTFlZlcmlT
aWduIFRydXN0IE5ldHdvcmsxOjA4BgNVBAsTMShjKSAxOTk5IFZlcmlTaWduLCBJ
bmMuIC0gRm9yIGF1dGhvcml6ZWQgdXNlIG9ubHkxRTBDBgNVBAMTPFZlcmlTaWdu
IENsYXNzIDEgUHVibGljIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkg
LSBHMzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAN2E1Lm0+afY8wR4
nN493GwTFtl63SRRZsDHJlkNrAYIwpTRMx/wgzUfbhvI3qpuFU5UJ+/EbRrsC+MO
8ESlV8dAWB6jRx9x7GD2bZTIGDnt/kIYVt/kTEkQeE4BdjVjEjbdZrwBBDajVWjV
ojYJrKshJlQGrT/KFOCsyq0GHZXi+J3x4GD/wn91K0zM2v6HmSHquv4+VNfSWXjb
PG7PoBMAGrgnoeS+Z5bKoMWznN3JdZ7rMJpfo83ZrngZPyPpXNspva1VyBtUjGP2
6KbqxzcSXKMpHgLZ2x87tNcPVkeBFQRKr4Mn0cVYiMHd9qqnoxjaaKptEVHhv2Vr
n5Z20T0CAwEAATANBgkqhkiG9w0BAQUFAAOCAQEAq2aN17O6x5q25lXQBfGfMY1a
qtmqRiYPce2lrVNWYgFHKkTp/j90CxObufRNG7LRX7K20ohcs5/Ny9Sn2WCVhDr4
wTcdYcrnsMXlkdpUpqwxga6X3s0IrLjAl4B/bnKk52kTlWUfxJM8/XmPBNQ+T+r3
ns7NZ3xPZQL/kYVUc8f/NveGLezQXk//EZ9yBta4GvFMDSZl4kSAHsef493oCtrs
pSCAaWihT37ha88HQfqDjrw43bAuEbFrskLMmrz5SCJ5ShkPshw+IHTZasO+8ih4
E1Z5T21Q6huwtVexN2ZYI/PcD98Kh8TvhgXVOBRgmaNL3gaWcSzy27YfpO8/7g==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEGTCCAwECEGFwy0mMX5hFKeewptlQW3owDQYJKoZIhvcNAQEFBQAwgcoxCzAJ
BgNVBAYTAlVTMRcwFQYDVQQKEw5WZXJpU2lnbiwgSW5jLjEfMB0GA1UECxMWVmVy
aVNpZ24gVHJ1c3QgTmV0d29yazE6MDgGA1UECxMxKGMpIDE5OTkgVmVyaVNpZ24s
IEluYy4gLSBGb3IgYXV0aG9yaXplZCB1c2Ugb25seTFFMEMGA1UEAxM8VmVyaVNp
Z24gQ2xhc3MgMiBQdWJsaWMgUHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0
eSAtIEczMB4XDTk5MTAwMTAwMDAwMFoXDTM2MDcxNjIzNTk1OVowgcoxCzAJBgNV
BAYTAlVTMRcwFQYDVQQKEw5WZXJpU2lnbiwgSW5jLjEfMB0GA1UECxMWVmVyaVNp
Z24gVHJ1c3QgTmV0d29yazE6MDgGA1UECxMxKGMpIDE5OTkgVmVyaVNpZ24sIElu
Yy4gLSBGb3IgYXV0aG9yaXplZCB1c2Ugb25seTFFMEMGA1UEAxM8VmVyaVNpZ24g
Q2xhc3MgMiBQdWJsaWMgUHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eSAt
IEczMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArwoNwtUs22e5LeWU
J92lvuCwTY+zYVY81nzD9M0+hsuiiOLh2KRpxbXiv8GmR1BeRjmL1Za6tW8UvxDO
JxOeBUebMXoT2B/Z0wI3i60sR/COgQanDTAM6/c8DyAd3HJG7qUCyFvDyVZpTMUY
wZF7C9UTAJu878NIPkZgIIUq1ZC2zYugzDLdt/1AVbJQHFauzI13TccgTacxdu9o
koqQHgiBVrKtaaNS0MscxCM9H5n+TOgWY47GCI72MfbS+uV23bUckqNJzc0BzWjN
qWm6o+sdDZykIKbBoMXRRkwXbdKsZj+WjOCE1Db/IlnF+RFgqF8EffIa9iVCYQ/E
Srg+iQIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQA0JhU8wI1NQ0kdvekhktdmnLfe
xbjQ5F1fdiLAJvmEOjr5jLX77GDx6M4EsMjdpwOPMPOY36TmpDHf0xwLRtxyID+u
7gU8pDM/CzmscHhzS5kr3zDCVLCoO1Wh/hYozUK9dG6A2ydEp85EXdQbkJgNHkKU
sQAsBNB0owIFImNjzYO1+8FtYmtpdf1dcEG59b98377BMnMiIYtYgXsVkXq642RI
sH/7NiXaldDxJBQX3RiAa0YjOVT1jmIJBB2UkKab5iXiQkWquJCtvgiPqQtCGJTP
cjnhsUPgKM+351psE2tJs//jGHyJizNdrDPXp/naOlXJWBD5qu9ats9LS98q
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEGjCCAwICEQCbfgZJoz5iudXukEhxKe9XMA0GCSqGSIb3DQEBBQUAMIHKMQsw
CQYDVQQGEwJVUzEXMBUGA1UEChMOVmVyaVNpZ24sIEluYy4xHzAdBgNVBAsTFlZl
cmlTaWduIFRydXN0IE5ldHdvcmsxOjA4BgNVBAsTMShjKSAxOTk5IFZlcmlTaWdu
LCBJbmMuIC0gRm9yIGF1dGhvcml6ZWQgdXNlIG9ubHkxRTBDBgNVBAMTPFZlcmlT
aWduIENsYXNzIDMgUHVibGljIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3Jp
dHkgLSBHMzAeFw05OTEwMDEwMDAwMDBaFw0zNjA3MTYyMzU5NTlaMIHKMQswCQYD
VQQGEwJVUzEXMBUGA1UEChMOVmVyaVNpZ24sIEluYy4xHzAdBgNVBAsTFlZlcmlT
aWduIFRydXN0IE5ldHdvcmsxOjA4BgNVBAsTMShjKSAxOTk5IFZlcmlTaWduLCBJ
bmMuIC0gRm9yIGF1dGhvcml6ZWQgdXNlIG9ubHkxRTBDBgNVBAMTPFZlcmlTaWdu
IENsYXNzIDMgUHVibGljIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkg
LSBHMzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMu6nFL8eB8aHm8b
N3O9+MlrlBIwT/A2R/XQkQr1F8ilYcEWQE37imGQ5XYgwREGfassbqb1EUGO+i2t
KmFZpGcmTNDovFJbcCAEWNF6yaRpvIMXZK0Fi7zQWM6NjPXr8EJJC52XJ2cybuGu
kxUccLwgTS8Y3pKI6GyFVxEa6X7jJhFUokWWVYPKMIno3Nij7SqAP395ZVc+FSBm
CC+Vk7+qRy+oRpfwEuL+wgorUeZ25rdGt+INpsyow0xZVYnm6FNcHOqd8GIWC6fJ
Xwzw3sJ2zq/3avL6QaaiMxTJ5Xpj055iN9WFZZ4O5lMkdBteHRJTW8cs54NJOxWu
imi5V5cCAwEAATANBgkqhkiG9w0BAQUFAAOCAQEAERSWwauSCPc/L8my/uRan2Te
2yFPhpk0djZX3dAVL8WtfxUfN2JzPtTnX84XA9s1+ivbrmAJXx5fj267Cz3qWhMe
DGBvtcC1IyIuBwvLqXTLR7sdwdela8wv0kL9Sd2nic9TutoAWii/gt/4uhMdUIaC
/Y4wjylGsB49Ndo4YhYYSq3mtlFs3q9i6wHQHiT+eo8SGhJouPtmmRQURVyu565p
F4ErWjfJXir0xuKhXFSbplQAz/DxwceYMBo7Nhbbo27q/a2ywtrvAkcTisDxszGt
TxzhT5yvDwyd93gN2PQ1VoDat20Xj50egWTh/sVFuq1ruQp6Tk9LhO5L8X3dEQ==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEGjCCAwICEQDsoKeLbnVqAc/EfMwvlF7XMA0GCSqGSIb3DQEBBQUAMIHKMQsw
CQYDVQQGEwJVUzEXMBUGA1UEChMOVmVyaVNpZ24sIEluYy4xHzAdBgNVBAsTFlZl
cmlTaWduIFRydXN0IE5ldHdvcmsxOjA4BgNVBAsTMShjKSAxOTk5IFZlcmlTaWdu
LCBJbmMuIC0gRm9yIGF1dGhvcml6ZWQgdXNlIG9ubHkxRTBDBgNVBAMTPFZlcmlT
aWduIENsYXNzIDQgUHVibGljIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3Jp
dHkgLSBHMzAeFw05OTEwMDEwMDAwMDBaFw0zNjA3MTYyMzU5NTlaMIHKMQswCQYD
VQQGEwJVUzEXMBUGA1UEChMOVmVyaVNpZ24sIEluYy4xHzAdBgNVBAsTFlZlcmlT
aWduIFRydXN0IE5ldHdvcmsxOjA4BgNVBAsTMShjKSAxOTk5IFZlcmlTaWduLCBJ
bmMuIC0gRm9yIGF1dGhvcml6ZWQgdXNlIG9ubHkxRTBDBgNVBAMTPFZlcmlTaWdu
IENsYXNzIDQgUHVibGljIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkg
LSBHMzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAK3LpRFpxlmr8Y+1
GQ9Wzsy1HyDkniYlS+BzZYlZ3tCD5PUPtbut8XzoIfzk6AzufEUiGXaStBO3IFsJ
+mGuqPKljYXCKtbeZjbSmwL0qJJgfJxptI8kHtCGUvYynEFYHiK9zUVilQhu0Gbd
U6LM8BDcVHOLBKFGMzNcF0C5nk3T875Vg+ixiY5afJqWIpA7iCXy0lOIAgwLePLm
NxdLMEYH5IBtptiWLugs+BGzOA1mppvqySNb247i8xOOGlktqgLw7KSHZtzBP/XY
ufTsgsbSPZUd5cBPhMnZo0QoBmrXRazwa2rvTl/4EYIeOGM0ZlDUPpNz+jDDZq3/
ky2X7wMCAwEAATANBgkqhkiG9w0BAQUFAAOCAQEAj/ola09b5KROJ1WrIhVZPMq1
CtRK26vdoV9TxaBXOcLORyu+OshWv8LZJxA6sQU8wHcxuzrTBXttmhwwjIDLk5Mq
g6sFUYICABFna/OIYUdfA5PVWw3g8dShMjWFsjrbsIKr0csKvE+MW8VLADsfKoKm
fjaF3H48ZwC15DtS4KjrXRX5xm3wrR0OhbepmnMUWluPQSjA1egtTaRezarZ7c7c
2NU8Qh0XwRJdRTjDOPP8hS6DRkiy1yBfkjaP53kPmF6Z6PDQpLv1U70qzlmwr25/
bLvSHgCwIe34QWKCudiyxLtGUPMxxY8BqHTr9Xgn2uf3ZkPznoM+IKrDNWCRzg==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDhDCCAwqgAwIBAgIQL4D+I4wOIg9IZxIokYesszAKBggqhkjOPQQDAzCByjEL
MAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQLExZW
ZXJpU2lnbiBUcnVzdCBOZXR3b3JrMTowOAYDVQQLEzEoYykgMjAwNyBWZXJpU2ln
biwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5MUUwQwYDVQQDEzxWZXJp
U2lnbiBDbGFzcyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9y
aXR5IC0gRzQwHhcNMDcxMTA1MDAwMDAwWhcNMzgwMTE4MjM1OTU5WjCByjELMAkG
A1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQLExZWZXJp
U2lnbiBUcnVzdCBOZXR3b3JrMTowOAYDVQQLEzEoYykgMjAwNyBWZXJpU2lnbiwg
SW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5MUUwQwYDVQQDEzxWZXJpU2ln
biBDbGFzcyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5
IC0gRzQwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAASnVnp8Utpkmw4tXNherJI9/gHm
GUo9FANL+mAnINmDiWn6VMaaGF5VKmTeBvaNSjutEDxlPZCIBIngMGGzrl0Bp3ve
fLK+ymVhAIau2o970ImtTR1ZmkGxvEeA3J5iw/mjgbIwga8wDwYDVR0TAQH/BAUw
AwEB/zAOBgNVHQ8BAf8EBAMCAQYwbQYIKwYBBQUHAQwEYTBfoV2gWzBZMFcwVRYJ
aW1hZ2UvZ2lmMCEwHzAHBgUrDgMCGgQUj+XTGoasjY5rw8+AatRIGCx7GS4wJRYj
aHR0cDovL2xvZ28udmVyaXNpZ24uY29tL3ZzbG9nby5naWYwHQYDVR0OBBYEFLMW
kf3upm7ktS5Jj4d4gYDs5bG1MAoGCCqGSM49BAMDA2gAMGUCMGYhDBgmYFo4e1ZC
4Kf8NoRRkSAsdk1DPcQdhCPQrNZ8NQbOzWm9kA3bbEhCHQ6qQgIxAJw9SDkjOVga
FRJZap7v1VmyHVIsmXHNxynfGyphe3HR3vPA5Q06Sqotp9iGKt0uEA==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIE0zCCA7ugAwIBAgIQGNrRniZ96LtKIVjNzGs7SjANBgkqhkiG9w0BAQUFADCB
yjELMAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQL
ExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMTowOAYDVQQLEzEoYykgMjAwNiBWZXJp
U2lnbiwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5MUUwQwYDVQQDEzxW
ZXJpU2lnbiBDbGFzcyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0
aG9yaXR5IC0gRzUwHhcNMDYxMTA4MDAwMDAwWhcNMzYwNzE2MjM1OTU5WjCByjEL
MAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQLExZW
ZXJpU2lnbiBUcnVzdCBOZXR3b3JrMTowOAYDVQQLEzEoYykgMjAwNiBWZXJpU2ln
biwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5MUUwQwYDVQQDEzxWZXJp
U2lnbiBDbGFzcyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9y
aXR5IC0gRzUwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCvJAgIKXo1
nmAMqudLO07cfLw8RRy7K+D+KQL5VwijZIUVJ/XxrcgxiV0i6CqqpkKzj/i5Vbex
t0uz/o9+B1fs70PbZmIVYc9gDaTY3vjgw2IIPVQT60nKWVSFJuUrjxuf6/WhkcIz
SdhDY2pSS9KP6HBRTdGJaXvHcPaz3BJ023tdS1bTlr8Vd6Gw9KIl8q8ckmcY5fQG
BO+QueQA5N06tRn/Arr0PO7gi+s3i+z016zy9vA9r911kTMZHRxAy3QkGSGT2RT+
rCpSx4/VBEnkjWNHiDxpg8v+R70rfk/Fla4OndTRQ8Bnc+MUCH7lP59zuDMKz10/
NIeWiu5T6CUVAgMBAAGjgbIwga8wDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8E
BAMCAQYwbQYIKwYBBQUHAQwEYTBfoV2gWzBZMFcwVRYJaW1hZ2UvZ2lmMCEwHzAH
BgUrDgMCGgQUj+XTGoasjY5rw8+AatRIGCx7GS4wJRYjaHR0cDovL2xvZ28udmVy
aXNpZ24uY29tL3ZzbG9nby5naWYwHQYDVR0OBBYEFH/TZafC3ey78DAJ80M5+gKv
MzEzMA0GCSqGSIb3DQEBBQUAA4IBAQCTJEowX2LP2BqYLz3q3JktvXf2pXkiOOzE
p6B4Eq1iDkVwZMXnl2YtmAl+X6/WzChl8gGqCBpH3vn5fJJaCGkgDdk+bW48DW7Y
5gaRQBi5+MHt39tBquCWIMnNZBU4gcmU7qKEKQsTb47bDN0lAtukixlE0kF6BWlK
WE9gyn6CagsCqiUXObXbf+eEZSqVir2G3l6BFoMtEMze/aiCKm0oHw0LxOXnGiYZ
4fQRbxC1lfznQgUy286dUV4otp6F01vvpX1FQHKOtw5rDgb7MzVIcbidJ4vEZV8N
hnacRHr2lVz2XTIIM6RUthg/aFzyQkqFOFSDX9HoLPKsEdao7WNq
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEuTCCA6GgAwIBAgIQQBrEZCGzEyEDDrvkEhrFHTANBgkqhkiG9w0BAQsFADCB
vTELMAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQL
ExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMTowOAYDVQQLEzEoYykgMjAwOCBWZXJp
U2lnbiwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5MTgwNgYDVQQDEy9W
ZXJpU2lnbiBVbml2ZXJzYWwgUm9vdCBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTAe
Fw0wODA0MDIwMDAwMDBaFw0zNzEyMDEyMzU5NTlaMIG9MQswCQYDVQQGEwJVUzEX
MBUGA1UEChMOVmVyaVNpZ24sIEluYy4xHzAdBgNVBAsTFlZlcmlTaWduIFRydXN0
IE5ldHdvcmsxOjA4BgNVBAsTMShjKSAyMDA4IFZlcmlTaWduLCBJbmMuIC0gRm9y
IGF1dGhvcml6ZWQgdXNlIG9ubHkxODA2BgNVBAMTL1ZlcmlTaWduIFVuaXZlcnNh
bCBSb290IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAx2E3XrEBNNti1xWb/1hajCMj1mCOkdeQmIN65lgZOIzF
9uVkhbSicfvtvbnazU0AtMgtc6XHaXGVHzk8skQHnOgO+k1KxCHfKWGPMiJhgsWH
H26MfF8WIFFE0XBPV+rjHOPMee5Y2A7Cs0WTwCznmhcrewA3ekEzeOEz4vMQGn+H
LL729fdC4uW/h2KJXwBL38Xd5HVEMkE6HnFuacsLdUYI0crSK5XQz/u5QGtkjFdN
/BMReYTtXlT2NJ8IAfMQJQYXStrxHXpma5hgZqTZ79IugvHw7wnqRMkVauIDbjPT
rJ9VAMf2CGqUuV/c4DPxhGD5WycRtPwW8rtWaoAljQIDAQABo4GyMIGvMA8GA1Ud
EwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgEGMG0GCCsGAQUFBwEMBGEwX6FdoFsw
WTBXMFUWCWltYWdlL2dpZjAhMB8wBwYFKw4DAhoEFI/l0xqGrI2Oa8PPgGrUSBgs
exkuMCUWI2h0dHA6Ly9sb2dvLnZlcmlzaWduLmNvbS92c2xvZ28uZ2lmMB0GA1Ud
DgQWBBS2d/ppSEefUxLVwuoHMnYH0ZcHGTANBgkqhkiG9w0BAQsFAAOCAQEASvj4
sAPmLGd75JR3Y8xuTPl9Dg3cyLk1uXBPY/ok+myDjEedO2Pzmvl2MpWRsXe8rJq+
seQxIcaBlVZaDrHC1LGmWazxY8u4TB1ZkErvkBYoH1quEPuBUDgMbMzxPcP1Y+Oz
4yHJJDnp/RVmRvQbEdBNc6N9Rvk97ahfYtTxP/jgdFcrGJ2BtMQo2pSXpXDrrB2+
BxHw1dvd5Yzw1TKwg+ZX4o+/vqGqvz0dtdQ46tewXDpPaj+PwGZsY6rp2aQW9IHR
lRQOfc2VNNnSj3BzgXucfr2YYdhFh5iQxeuGMMY1v/D/w1WIg0vvBZIGcfK4mJO3
7M2CYfE45k+XmCpajQ==
-----END CERTIFICATE-----
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            02:ac:5c:26:6a:0b:40:9b:8f:0b:79:f2:ae:46:25:77
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Validity
            Not Before: Nov 10 00:00:00 2006 GMT
            Not After : Nov 10 00:00:00 2031 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:c6:cc:e5:73:e6:fb:d4:bb:e5:2d:2d:32:a6:df:
                    e5:81:3f:c9:cd:25:49:b6:71:2a:c3:d5:94:34:67:
                    a2:0a:1c:b0:5f:69:a6:40:b1:c4:b7:b2:8f:d0:98:
                    a4:a9:41:59:3a:d3:dc:94:d6:3c:db:74:38:a4:4a:
                    cc:4d:25:82:f7:4a:a5:53:12:38:ee:f3:49:6d:71:
                    91:7e:63:b6:ab:a6:5f:c3:a4:84:f8:4f:62:51:be:
                    f8:c5:ec:db:38:92:e3:06:e5:08:91:0c:c4:28:41:
                    55:fb:cb:5a:89:15:7e:71:e8:35:bf:4d:72:09:3d:
                    be:3a:38:50:5b:77:31:1b:8d:b3:c7:24:45:9a:a7:
                    ac:6d:00:14:5a:04:b7:ba:13:eb:51:0a:98:41:41:
                    22:4e:65:61:87:81:41:50:a6:79:5c:89:de:19:4a:
                    57:d5:2e:e6:5d:1c:53:2c:7e:98:cd:1a:06:16:a4:
                    68:73:d0:34:04:13:5c:a1:71:d3:5a:7c:55:db:5e:
                    64:e1:37:87:30:56:04:e5:11:b4:29:80:12:f1:79:
                    39:88:a2:02:11:7c:27:66:b7:88:b7:78:f2:ca:0a:
                    a8:38:ab:0a:64:c2:bf:66:5d:95:84:c1:a1:25:1e:
                    87:5d:1a:50:0b:20:12:cc:41:bb:6e:0b:51:38:b8:
                    4b:cb
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier: 
                B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3
            X509v3 Authority Key Identifier: 
                keyid:B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3

    Signature Algorithm: sha1WithRSAEncryption
         1c:1a:06:97:dc:d7:9c:9f:3c:88:66:06:08:57:21:db:21:47:
         f8:2a:67:aa:bf:18:32:76:40:10:57:c1:8a:f3:7a:d9:11:65:
         8e:35:fa:9e:fc:45:b5:9e:d9:4c:31:4b:b8:91:e8:43:2c:8e:
         b3:78:ce:db:e3:53:79:71:d6:e5:21:94:01:da:55:87:9a:24:
         64:f6:8a:66:cc:de:9c:37:cd:a8:34:b1:69:9b:23:c8:9e:78:
         22:2b:70:43:e3:55:47:31:61:19:ef:58:c5:85:2f:4e:30:f6:
         a0:31:16:23:c8:e7:e2:65:16:33:cb:bf:1a:1b:a0:3d:f8:ca:
         5e:8b:31:8b:60:08:89:2d:0c:06:5c:52:b7:c4:f9:0a:98:d1:
         15:5f:9f:12:be:7c:36:63:38:bd:44:a4:7f:e4:26:2b:0a:c4:
         97:69:0d:e9:8c:e2:c0:10:57:b8:c8:76:12:91:55:f2:48:69:
         d8:bc:2a:02:5b:0f:44:d4:20:31:db:f4:ba:70:26:5d:90:60:
         9e:bc:4b:17:09:2f:b4:cb:1e:43:68:c9:07:27:c1:d2:5c:f7:
         ea:21:b9:68:12:9c:3c:9c:bf:9e:fc:80:5c:9b:63:cd:ec:47:
         aa:25:27:67:a0:37:f3:00:82:7d:54:d7:a9:f8:e9:2e:13:a3:
         77:e8:1f:4a
-----BEGIN CERTIFICATE-----
MIIDxTCCAq2gAwIBAgIQAqxcJmoLQJuPC3nyrkYldzANBgkqhkiG9w0BAQUFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTA2MTExMDAwMDAwMFoXDTMxMTExMDAwMDAwMFowbDEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTErMCkGA1UEAxMiRGlnaUNlcnQgSGlnaCBBc3N1cmFuY2Ug
RVYgUm9vdCBDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMbM5XPm
+9S75S0tMqbf5YE/yc0lSbZxKsPVlDRnogocsF9ppkCxxLeyj9CYpKlBWTrT3JTW
PNt0OKRKzE0lgvdKpVMSOO7zSW1xkX5jtqumX8OkhPhPYlG++MXs2ziS4wblCJEM
xChBVfvLWokVfnHoNb9Ncgk9vjo4UFt3MRuNs8ckRZqnrG0AFFoEt7oT61EKmEFB
Ik5lYYeBQVCmeVyJ3hlKV9Uu5l0cUyx+mM0aBhakaHPQNAQTXKFx01p8VdteZOE3
hzBWBOURtCmAEvF5OYiiAhF8J2a3iLd48soKqDirCmTCv2ZdlYTBoSUeh10aUAsg
EsxBu24LUTi4S8sCAwEAAaNjMGEwDgYDVR0PAQH/BAQDAgGGMA8GA1UdEwEB/wQF
MAMBAf8wHQYDVR0OBBYEFLE+w2kD+L9HAdSYJhoIAu9jZCvDMB8GA1UdIwQYMBaA
FLE+w2kD+L9HAdSYJhoIAu9jZCvDMA0GCSqGSIb3DQEBBQUAA4IBAQAcGgaX3Nec
nzyIZgYIVyHbIUf4KmeqvxgydkAQV8GK83rZEWWONfqe/EW1ntlMMUu4kehDLI6z
eM7b41N5cdblIZQB2lWHmiRk9opmzN6cN82oNLFpmyPInngiK3BD41VHMWEZ71jF
hS9OMPagMRYjyOfiZRYzy78aG6A9+MpeizGLYAiJLQwGXFK3xPkKmNEVX58Svnw2
Yzi9RKR/5CYrCsSXaQ3pjOLAEFe4yHYSkVXySGnYvCoCWw9E1CAx2/S6cCZdkGCe
vEsXCS+0yx5DaMkHJ8HSXPfqIbloEpw8nL+e/IBcm2PN7EeqJSdnoDfzAIJ9VNep
+OkuE6N36B9K
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0c:e7:e0:e5:17:d8:46:fe:8f:e5:60:fc:1b:f0:30:39
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root CA
        Validity
            Not Before: Nov 10 00:00:00 2006 GMT
            Not After : Nov 10 00:00:00 2031 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:ad:0e:15:ce:e4:43:80:5c:b1:87:f3:b7:60:f9:
                    71:12:a5:ae:dc:26:94:88:aa:f4:ce:f5:20:39:28:
                    58:60:0c:f8:80:da:a9:15:95:32:61:3c:b5:b1:28:
                    84:8a:8a:dc:9f:0a:0c:83:17:7a:8f:90:ac:8a:e7:
                    79:53:5c:31:84:2a:f6:0f:98:32:36:76:cc:de:dd:
                    3c:a8:a2:ef:6a:fb:21:f2:52:61:df:9f:20:d7:1f:
                    e2:b1:d9:fe:18:64:d2:12:5b:5f:f9:58:18:35:bc:
                    47:cd:a1:36:f9:6b:7f:d4:b0:38:3e:c1:1b:c3:8c:
                    33:d9:d8:2f:18:fe:28:0f:b3:a7:83:d6:c3:6e:44:
                    c0:61:35:96:16:fe:59:9c:8b:76:6d:d7:f1:a2:4b:
                    0d:2b:ff:0b:72:da:9e:60:d0:8e:90:35:c6:78:55:
                    87:20:a1:cf:e5:6d:0a:c8:49:7c:31:98:33:6c:22:
                    e9:87:d0:32:5a:a2:ba:13:82:11:ed:39:17:9d:99:
                    3a:72:a1:e6:fa:a4:d9:d5:17:31:75:ae:85:7d:22:
                    ae:3f:01:46:86:f6:28:79:c8:b1:da:e4:57:17:c4:
                    7e:1c:0e:b0:b4:92:a6:56:b3:bd:b2:97:ed:aa:a7:
                    f0:b7:c5:a8:3f:95:16:d0:ff:a1:96:eb:08:5f:18:
                    77:4f
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier: 
                45:EB:A2:AF:F4:92:CB:82:31:2D:51:8B:A7:A7:21:9D:F3:6D:C8:0F
            X509v3 Authority Key Identifier: 
                keyid:45:EB:A2:AF:F4:92:CB:82:31:2D:51:8B:A7:A7:21:9D:F3:6D:C8:0F

    Signature Algorithm: sha1WithRSAEncryption
         a2:0e:bc:df:e2:ed:f0:e3:72:73:7a:64:94:bf:f7:72:66:d8:
         32:e4:42:75:62:ae:87:eb:f2:d5:d9:de:56:b3:9f:cc:ce:14:
         28:b9:0d:97:60:5c:12:4c:58:e4:d3:3d:83:49:45:58:97:35:
         69:1a:a8:47:ea:56:c6:79:ab:12:d8:67:81:84:df:7f:09:3c:
         94:e6:b8:26:2c:20:bd:3d:b3:28:89:f7:5f:ff:22:e2:97:84:
         1f:e9:65:ef:87:e0:df:c1:67:49:b3:5d:eb:b2:09:2a:eb:26:
         ed:78:be:7d:3f:2b:f3:b7:26:35:6d:5f:89:01:b6:49:5b:9f:
         01:05:9b:ab:3d:25:c1:cc:b6:7f:c2:f1:6f:86:c6:fa:64:68:
         eb:81:2d:94:eb:42:b7:fa:8c:1e:dd:62:f1:be:50:67:b7:6c:
         bd:f3:f1:1f:6b:0c:36:07:16:7f:37:7c:a9:5b:6d:7a:f1:12:
         46:60:83:d7:27:04:be:4b:ce:97:be:c3:67:2a:68:11:df:80:
         e7:0c:33:66:bf:13:0d:14:6e:f3:7f:1f:63:10:1e:fa:8d:1b:
         25:6d:6c:8f:a5:b7:61:01:b1:d2:a3:26:a1:10:71:9d:ad:e2:
         c3:f9:c3:99:51:b7:2b:07:08:ce:2e:e6:50:b2:a7:fa:0a:45:
         2f:a2:f0:f2
-----BEGIN CERTIFICATE-----
MIIDtzCCAp+gAwIBAgIQDOfg5RfYRv6P5WD8G/AwOTANBgkqhkiG9w0BAQUFADBl
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJv
b3QgQ0EwHhcNMDYxMTEwMDAwMDAwWhcNMzExMTEwMDAwMDAwWjBlMQswCQYDVQQG
EwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNl
cnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJvb3QgQ0EwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCtDhXO5EOAXLGH87dg+XESpa7c
JpSIqvTO9SA5KFhgDPiA2qkVlTJhPLWxKISKityfCgyDF3qPkKyK53lTXDGEKvYP
mDI2dsze3Tyoou9q+yHyUmHfnyDXH+Kx2f4YZNISW1/5WBg1vEfNoTb5a3/UsDg+
wRvDjDPZ2C8Y/igPs6eD1sNuRMBhNZYW/lmci3Zt1/GiSw0r/wty2p5g0I6QNcZ4
VYcgoc/lbQrISXwxmDNsIumH0DJaoroTghHtORedmTpyoeb6pNnVFzF1roV9Iq4/
AUaG9ih5yLHa5FcXxH4cDrC0kqZWs72yl+2qp/C3xag/lRbQ/6GW6whfGHdPAgMB
AAGjYzBhMA4GA1UdDwEB/wQEAwIBhjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQW
BBRF66Kv9JLLgjEtUYunpyGd823IDzAfBgNVHSMEGDAWgBRF66Kv9JLLgjEtUYun
pyGd823IDzANBgkqhkiG9w0BAQUFAAOCAQEAog683+Lt8ONyc3pklL/3cmbYMuRC
dWKuh+vy1dneVrOfzM4UKLkNl2BcEkxY5NM9g0lFWJc1aRqoR+pWxnmrEthngYTf
fwk8lOa4JiwgvT2zKIn3X/8i4peEH+ll74fg38FnSbNd67IJKusm7Xi+fT8r87cm
NW1fiQG2SVufAQWbqz0lwcy2f8Lxb4bG+mRo64EtlOtCt/qMHt1i8b5QZ7dsvfPx
H2sMNgcWfzd8qVttevESRmCD1ycEvkvOl77DZypoEd+A5wwzZr8TDRRu838fYxAe
+o0bJW1sj6W3YQGx0qMmoRBxna3iw/nDmVG3KwcIzi7mULKn+gpFL6Lw8g==
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            08:3b:e0:56:90:42:46:b1:a1:75:6a:c9:59:91:c7:4a
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root CA
        Validity
            Not Before: Nov 10 00:00:00 2006 GMT
            Not After : Nov 10 00:00:00 2031 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:e2:3b:e1:11:72:de:a8:a4:d3:a3:57:aa:50:a2:
                    8f:0b:77:90:c9:a2:a5:ee:12:ce:96:5b:01:09:20:
                    cc:01:93:a7:4e:30:b7:53:f7:43:c4:69:00:57:9d:
                    e2:8d:22:dd:87:06:40:00:81:09:ce:ce:1b:83:bf:
                    df:cd:3b:71:46:e2:d6:66:c7:05:b3:76:27:16:8f:
                    7b:9e:1e:95:7d:ee:b7:48:a3:08:da:d6:af:7a:0c:
                    39:06:65:7f:4a:5d:1f:bc:17:f8:ab:be:ee:28:d7:
                    74:7f:7a:78:99:59:85:68:6e:5c:23:32:4b:bf:4e:
                    c0:e8:5a:6d:e3:70:bf:77:10:bf:fc:01:f6:85:d9:
                    a8:44:10:58:32:a9:75:18:d5:d1:a2:be:47:e2:27:
                    6a:f4:9a:33:f8:49:08:60:8b:d4:5f:b4:3a:84:bf:
                    a1:aa:4a:4c:7d:3e:cf:4f:5f:6c:76:5e:a0:4b:37:
                    91:9e:dc:22:e6:6d:ce:14:1a:8e:6a:cb:fe:cd:b3:
                    14:64:17:c7:5b:29:9e:32:bf:f2:ee:fa:d3:0b:42:
                    d4:ab:b7:41:32:da:0c:d4:ef:f8:81:d5:bb:8d:58:
                    3f:b5:1b:e8:49:28:a2:70:da:31:04:dd:f7:b2:16:
                    f2:4c:0a:4e:07:a8:ed:4a:3d:5e:b5:7f:a3:90:c3:
                    af:27
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier: 
                03:DE:50:35:56:D1:4C:BB:66:F0:A3:E2:1B:1B:C3:97:B2:3D:D1:55
            X509v3 Authority Key Identifier: 
                keyid:03:DE:50:35:56:D1:4C:BB:66:F0:A3:E2:1B:1B:C3:97:B2:3D:D1:55

    Signature Algorithm: sha1WithRSAEncryption
         cb:9c:37:aa:48:13:12:0a:fa:dd:44:9c:4f:52:b0:f4:df:ae:
         04:f5:79:79:08:a3:24:18:fc:4b:2b:84:c0:2d:b9:d5:c7:fe:
         f4:c1:1f:58:cb:b8:6d:9c:7a:74:e7:98:29:ab:11:b5:e3:70:
         a0:a1:cd:4c:88:99:93:8c:91:70:e2:ab:0f:1c:be:93:a9:ff:
         63:d5:e4:07:60:d3:a3:bf:9d:5b:09:f1:d5:8e:e3:53:f4:8e:
         63:fa:3f:a7:db:b4:66:df:62:66:d6:d1:6e:41:8d:f2:2d:b5:
         ea:77:4a:9f:9d:58:e2:2b:59:c0:40:23:ed:2d:28:82:45:3e:
         79:54:92:26:98:e0:80:48:a8:37:ef:f0:d6:79:60:16:de:ac:
         e8:0e:cd:6e:ac:44:17:38:2f:49:da:e1:45:3e:2a:b9:36:53:
         cf:3a:50:06:f7:2e:e8:c4:57:49:6c:61:21:18:d5:04:ad:78:
         3c:2c:3a:80:6b:a7:eb:af:15:14:e9:d8:89:c1:b9:38:6c:e2:
         91:6c:8a:ff:64:b9:77:25:57:30:c0:1b:24:a3:e1:dc:e9:df:
         47:7c:b5:b4:24:08:05:30:ec:2d:bd:0b:bf:45:bf:50:b9:a9:
         f3:eb:98:01:12:ad:c8:88:c6:98:34:5f:8d:0a:3c:c6:e9:d5:
         95:95:6d:de
-----BEGIN CERTIFICATE-----
MIIDrzCCApegAwIBAgIQCDvgVpBCRrGhdWrJWZHHSjANBgkqhkiG9w0BAQUFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0wNjExMTAwMDAwMDBaFw0zMTExMTAwMDAwMDBaMGExCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j
b20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IENBMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4jvhEXLeqKTTo1eqUKKPC3eQyaKl7hLOllsB
CSDMAZOnTjC3U/dDxGkAV53ijSLdhwZAAIEJzs4bg7/fzTtxRuLWZscFs3YnFo97
nh6Vfe63SKMI2tavegw5BmV/Sl0fvBf4q77uKNd0f3p4mVmFaG5cIzJLv07A6Fpt
43C/dxC//AH2hdmoRBBYMql1GNXRor5H4idq9Joz+EkIYIvUX7Q6hL+hqkpMfT7P
T19sdl6gSzeRntwi5m3OFBqOasv+zbMUZBfHWymeMr/y7vrTC0LUq7dBMtoM1O/4
gdW7jVg/tRvoSSiicNoxBN33shbyTApOB6jtSj1etX+jkMOvJwIDAQABo2MwYTAO
BgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUA95QNVbR
TLtm8KPiGxvDl7I90VUwHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUw
DQYJKoZIhvcNAQEFBQADggEBAMucN6pIExIK+t1EnE9SsPTfrgT1eXkIoyQY/Esr
hMAtudXH/vTBH1jLuG2cenTnmCmrEbXjcKChzUyImZOMkXDiqw8cvpOp/2PV5Adg
06O/nVsJ8dWO41P0jmP6P6fbtGbfYmbW0W5BjfIttep3Sp+dWOIrWcBAI+0tKIJF
PnlUkiaY4IBIqDfv8NZ5YBberOgOzW6sRBc4L0na4UU+Krk2U886UAb3LujEV0ls
YSEY1QSteDwsOoBrp+uvFRTp2InBuThs4pFsiv9kuXclVzDAGySj4dzp30d8tbQk
CAUw7C29C79Fv1C5qfPrmAESrciIxpg0X40KPMbp1ZWVbd4=
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0b:93:1c:3a:d6:39:67:ea:67:23:bf:c3:af:9a:f4:4b
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root G2
        Validity
            Not Before: Aug  1 12:00:00 2013 GMT
            Not After : Jan 15 12:00:00 2038 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root G2
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:d9:e7:28:2f:52:3f:36:72:49:88:93:34:f3:f8:
                    6a:1e:31:54:80:9f:ad:54:41:b5:47:df:96:a8:d4:
                    af:80:2d:b9:0a:cf:75:fd:89:a5:7d:24:fa:e3:22:
                    0c:2b:bc:95:17:0b:33:bf:19:4d:41:06:90:00:bd:
                    0c:4d:10:fe:07:b5:e7:1c:6e:22:55:31:65:97:bd:
                    d3:17:d2:1e:62:f3:db:ea:6c:50:8c:3f:84:0c:96:
                    cf:b7:cb:03:e0:ca:6d:a1:14:4c:1b:89:dd:ed:00:
                    b0:52:7c:af:91:6c:b1:38:13:d1:e9:12:08:c0:00:
                    b0:1c:2b:11:da:77:70:36:9b:ae:ce:79:87:dc:82:
                    70:e6:09:74:70:55:69:af:a3:68:9f:bf:dd:b6:79:
                    b3:f2:9d:70:29:55:f4:ab:ff:95:61:f3:c9:40:6f:
                    1d:d1:be:93:bb:d3:88:2a:bb:9d:bf:72:5a:56:71:
                    3b:3f:d4:f3:d1:0a:fe:28:ef:a3:ee:d9:99:af:03:
                    d3:8f:60:b7:f2:92:a1:b1:bd:89:89:1f:30:cd:c3:
                    a6:2e:62:33:ae:16:02:77:44:5a:e7:81:0a:3c:a7:
                    44:2e:79:b8:3f:04:bc:5c:a0:87:e1:1b:af:51:8e:
                    cd:ec:2c:fa:f8:fe:6d:f0:3a:7c:aa:8b:e4:67:95:
                    31:8d
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Subject Key Identifier: 
                CE:C3:4A:B9:99:55:F2:B8:DB:60:BF:A9:7E:BD:56:B5:97:36:A7:D6
    Signature Algorithm: sha256WithRSAEncryption
         ca:a5:55:8c:e3:c8:41:6e:69:27:a7:75:11:ef:3c:86:36:6f:
         d2:9d:c6:78:38:1d:69:96:a2:92:69:2e:38:6c:9b:7d:04:d4:
         89:a5:b1:31:37:8a:c9:21:cc:ab:6c:cd:8b:1c:9a:d6:bf:48:
         d2:32:66:c1:8a:c0:f3:2f:3a:ef:c0:e3:d4:91:86:d1:50:e3:
         03:db:73:77:6f:4a:39:53:ed:de:26:c7:b5:7d:af:2b:42:d1:
         75:62:e3:4a:2b:02:c7:50:4b:e0:69:e2:96:6c:0e:44:66:10:
         44:8f:ad:05:eb:f8:79:ac:a6:1b:e8:37:34:9d:53:c9:61:aa:
         a2:52:af:4a:70:16:86:c2:3a:c8:b1:13:70:36:d8:cf:ee:f4:
         0a:34:d5:5b:4c:fd:07:9c:a2:ba:d9:01:72:5c:f3:4d:c1:dd:
         0e:b1:1c:0d:c4:63:be:ad:f4:14:fb:89:ec:a2:41:0e:4c:cc:
         c8:57:40:d0:6e:03:aa:cd:0c:8e:89:99:99:6c:f0:3c:30:af:
         38:df:6f:bc:a3:be:29:20:27:ab:74:ff:13:22:78:de:97:52:
         55:1e:83:b5:54:20:03:ee:ae:c0:4f:56:de:37:cc:c3:7f:aa:
         04:27:bb:d3:77:b8:62:db:17:7c:9c:28:22:13:73:6c:cf:26:
         f5:8a:29:e7
-----BEGIN CERTIFICATE-----
MIIDljCCAn6gAwIBAgIQC5McOtY5Z+pnI7/Dr5r0SzANBgkqhkiG9w0BAQsFADBl
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJv
b3QgRzIwHhcNMTMwODAxMTIwMDAwWhcNMzgwMTE1MTIwMDAwWjBlMQswCQYDVQQG
EwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNl
cnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJvb3QgRzIwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDZ5ygvUj82ckmIkzTz+GoeMVSA
n61UQbVH35ao1K+ALbkKz3X9iaV9JPrjIgwrvJUXCzO/GU1BBpAAvQxNEP4Htecc
biJVMWWXvdMX0h5i89vqbFCMP4QMls+3ywPgym2hFEwbid3tALBSfK+RbLE4E9Hp
EgjAALAcKxHad3A2m67OeYfcgnDmCXRwVWmvo2ifv922ebPynXApVfSr/5Vh88lA
bx3RvpO704gqu52/clpWcTs/1PPRCv4o76Pu2ZmvA9OPYLfykqGxvYmJHzDNw6Yu
YjOuFgJ3RFrngQo8p0Quebg/BLxcoIfhG69Rjs3sLPr4/m3wOnyqi+RnlTGNAgMB
AAGjQjBAMA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgGGMB0GA1UdDgQW
BBTOw0q5mVXyuNtgv6l+vVa1lzan1jANBgkqhkiG9w0BAQsFAAOCAQEAyqVVjOPI
QW5pJ6d1Ee88hjZv0p3GeDgdaZaikmkuOGybfQTUiaWxMTeKySHMq2zNixya1r9I
0jJmwYrA8y8678Dj1JGG0VDjA9tzd29KOVPt3ibHtX2vK0LRdWLjSisCx1BL4Gni
lmwORGYQRI+tBev4eaymG+g3NJ1TyWGqolKvSnAWhsI6yLETcDbYz+70CjTVW0z9
B5yiutkBclzzTcHdDrEcDcRjvq30FPuJ7KJBDkzMyFdA0G4Dqs0MjomZmWzwPDCv
ON9vvKO+KSAnq3T/EyJ43pdSVR6DtVQgA+6uwE9W3jfMw3+qBCe703e4YtsXfJwo
IhNzbM8m9Yop5w==
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0b:a1:5a:fa:1d:df:a0:b5:49:44:af:cd:24:a0:6c:ec
    Signature Algorithm: ecdsa-with-SHA384
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root G3
        Validity
            Not Before: Aug  1 12:00:00 2013 GMT
            Not After : Jan 15 12:00:00 2038 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root G3
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (384 bit)
                pub: 
                    04:19:e7:bc:ac:44:65:ed:cd:b8:3f:58:fb:8d:b1:
                    57:a9:44:2d:05:15:f2:ef:0b:ff:10:74:9f:b5:62:
                    52:5f:66:7e:1f:e5:dc:1b:45:79:0b:cc:c6:53:0a:
                    9d:8d:5d:02:d9:a9:59:de:02:5a:f6:95:2a:0e:8d:
                    38:4a:8a:49:c6:bc:c6:03:38:07:5f:55:da:7e:09:
                    6e:e2:7f:5e:d0:45:20:0f:59:76:10:d6:a0:24:f0:
                    2d:de:36:f2:6c:29:39
                ASN1 OID: secp384r1
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Subject Key Identifier: 
                CB:D0:BD:A9:E1:98:05:51:A1:4D:37:A2:83:79:CE:8D:1D:2A:E4:84
    Signature Algorithm: ecdsa-with-SHA384
         30:64:02:30:25:a4:81:45:02:6b:12:4b:75:74:4f:c8:23:e3:
         70:f2:75:72:de:7c:89:f0:cf:91:72:61:9e:5e:10:92:59:56:
         b9:83:c7:10:e7:38:e9:58:26:36:7d:d5:e4:34:86:39:02:30:
         7c:36:53:f0:30:e5:62:63:3a:99:e2:b6:a3:3b:9b:34:fa:1e:
         da:10:92:71:5e:91:13:a7:dd:a4:6e:92:cc:32:d6:f5:21:66:
         c7:2f:ea:96:63:6a:65:45:92:95:01:b4
-----BEGIN CERTIFICATE-----
MIICRjCCAc2gAwIBAgIQC6Fa+h3foLVJRK/NJKBs7DAKBggqhkjOPQQDAzBlMQsw
CQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cu
ZGlnaWNlcnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJvb3Qg
RzMwHhcNMTMwODAxMTIwMDAwWhcNMzgwMTE1MTIwMDAwWjBlMQswCQYDVQQGEwJV
UzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNlcnQu
Y29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJvb3QgRzMwdjAQBgcq
hkjOPQIBBgUrgQQAIgNiAAQZ57ysRGXtzbg/WPuNsVepRC0FFfLvC/8QdJ+1YlJf
Zn4f5dwbRXkLzMZTCp2NXQLZqVneAlr2lSoOjThKiknGvMYDOAdfVdp+CW7if17Q
RSAPWXYQ1qAk8C3eNvJsKTmjQjBAMA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/
BAQDAgGGMB0GA1UdDgQWBBTL0L2p4ZgFUaFNN6KDec6NHSrkhDAKBggqhkjOPQQD
AwNnADBkAjAlpIFFAmsSS3V0T8gj43DydXLefInwz5FyYZ5eEJJZVrmDxxDnOOlY
JjZ91eQ0hjkCMHw2U/Aw5WJjOpnitqM7mzT6HtoQknFekROn3aRukswy1vUhZscv
6pZjamVFkpUBtA==
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            03:3a:f1:e6:a7:11:a9:a0:bb:28:64:b1:1d:09:fa:e5
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root G2
        Validity
            Not Before: Aug  1 12:00:00 2013 GMT
            Not After : Jan 15 12:00:00 2038 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root G2
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:bb:37:cd:34:dc:7b:6b:c9:b2:68:90:ad:4a:75:
                    ff:46:ba:21:0a:08:8d:f5:19:54:c9:fb:88:db:f3:
                    ae:f2:3a:89:91:3c:7a:e6:ab:06:1a:6b:cf:ac:2d:
                    e8:5e:09:24:44:ba:62:9a:7e:d6:a3:a8:7e:e0:54:
                    75:20:05:ac:50:b7:9c:63:1a:6c:30:dc:da:1f:19:
                    b1:d7:1e:de:fd:d7:e0:cb:94:83:37:ae:ec:1f:43:
                    4e:dd:7b:2c:d2:bd:2e:a5:2f:e4:a9:b8:ad:3a:d4:
                    99:a4:b6:25:e9:9b:6b:00:60:92:60:ff:4f:21:49:
                    18:f7:67:90:ab:61:06:9c:8f:f2:ba:e9:b4:e9:92:
                    32:6b:b5:f3:57:e8:5d:1b:cd:8c:1d:ab:95:04:95:
                    49:f3:35:2d:96:e3:49:6d:dd:77:e3:fb:49:4b:b4:
                    ac:55:07:a9:8f:95:b3:b4:23:bb:4c:6d:45:f0:f6:
                    a9:b2:95:30:b4:fd:4c:55:8c:27:4a:57:14:7c:82:
                    9d:cd:73:92:d3:16:4a:06:0c:8c:50:d1:8f:1e:09:
                    be:17:a1:e6:21:ca:fd:83:e5:10:bc:83:a5:0a:c4:
                    67:28:f6:73:14:14:3d:46:76:c3:87:14:89:21:34:
                    4d:af:0f:45:0c:a6:49:a1:ba:bb:9c:c5:b1:33:83:
                    29:85
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Subject Key Identifier: 
                4E:22:54:20:18:95:E6:E3:6E:E6:0F:FA:FA:B9:12:ED:06:17:8F:39
    Signature Algorithm: sha256WithRSAEncryption
         60:67:28:94:6f:0e:48:63:eb:31:dd:ea:67:18:d5:89:7d:3c:
         c5:8b:4a:7f:e9:be:db:2b:17:df:b0:5f:73:77:2a:32:13:39:
         81:67:42:84:23:f2:45:67:35:ec:88:bf:f8:8f:b0:61:0c:34:
         a4:ae:20:4c:84:c6:db:f8:35:e1:76:d9:df:a6:42:bb:c7:44:
         08:86:7f:36:74:24:5a:da:6c:0d:14:59:35:bd:f2:49:dd:b6:
         1f:c9:b3:0d:47:2a:3d:99:2f:bb:5c:bb:b5:d4:20:e1:99:5f:
         53:46:15:db:68:9b:f0:f3:30:d5:3e:31:e2:8d:84:9e:e3:8a:
         da:da:96:3e:35:13:a5:5f:f0:f9:70:50:70:47:41:11:57:19:
         4e:c0:8f:ae:06:c4:95:13:17:2f:1b:25:9f:75:f2:b1:8e:99:
         a1:6f:13:b1:41:71:fe:88:2a:c8:4f:10:20:55:d7:f3:14:45:
         e5:e0:44:f4:ea:87:95:32:93:0e:fe:53:46:fa:2c:9d:ff:8b:
         22:b9:4b:d9:09:45:a4:de:a4:b8:9a:58:dd:1b:7d:52:9f:8e:
         59:43:88:81:a4:9e:26:d5:6f:ad:dd:0d:c6:37:7d:ed:03:92:
         1b:e5:77:5f:76:ee:3c:8d:c4:5d:56:5b:a2:d9:66:6e:b3:35:
         37:e5:32:b6
-----BEGIN CERTIFICATE-----
MIIDjjCCAnagAwIBAgIQAzrx5qcRqaC7KGSxHQn65TANBgkqhkiG9w0BAQsFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBH
MjAeFw0xMzA4MDExMjAwMDBaFw0zODAxMTUxMjAwMDBaMGExCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j
b20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IEcyMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuzfNNNx7a8myaJCtSnX/RrohCgiN9RlUyfuI
2/Ou8jqJkTx65qsGGmvPrC3oXgkkRLpimn7Wo6h+4FR1IAWsULecYxpsMNzaHxmx
1x7e/dfgy5SDN67sH0NO3Xss0r0upS/kqbitOtSZpLYl6ZtrAGCSYP9PIUkY92eQ
q2EGnI/yuum06ZIya7XzV+hdG82MHauVBJVJ8zUtluNJbd134/tJS7SsVQepj5Wz
tCO7TG1F8PapspUwtP1MVYwnSlcUfIKdzXOS0xZKBgyMUNGPHgm+F6HmIcr9g+UQ
vIOlCsRnKPZzFBQ9RnbDhxSJITRNrw9FDKZJobq7nMWxM4MphQIDAQABo0IwQDAP
BgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNVHQ4EFgQUTiJUIBiV
5uNu5g/6+rkS7QYXjzkwDQYJKoZIhvcNAQELBQADggEBAGBnKJRvDkhj6zHd6mcY
1Yl9PMWLSn/pvtsrF9+wX3N3KjITOYFnQoQj8kVnNeyIv/iPsGEMNKSuIEyExtv4
NeF22d+mQrvHRAiGfzZ0JFrabA0UWTW98kndth/Jsw1HKj2ZL7tcu7XUIOGZX1NG
Fdtom/DzMNU+MeKNhJ7jitralj41E6Vf8PlwUHBHQRFXGU7Aj64GxJUTFy8bJZ91
8rGOmaFvE7FBcf6IKshPECBV1/MUReXgRPTqh5Uykw7+U0b6LJ3/iyK5S9kJRaTe
pLiaWN0bfVKfjllDiIGknibVb63dDcY3fe0Dkhvld1927jyNxF1WW6LZZm6zNTfl
MrY=
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            05:55:56:bc:f2:5e:a4:35:35:c3:a4:0f:d5:ab:45:72
    Signature Algorithm: ecdsa-with-SHA384
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root G3
        Validity
            Not Before: Aug  1 12:00:00 2013 GMT
            Not After : Jan 15 12:00:00 2038 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root G3
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (384 bit)
                pub: 
                    04:dd:a7:d9:bb:8a:b8:0b:fb:0b:7f:21:d2:f0:be:
                    be:73:f3:33:5d:1a:bc:34:ea:de:c6:9b:bc:d0:95:
                    f6:f0:cc:d0:0b:ba:61:5b:51:46:7e:9e:2d:9f:ee:
                    8e:63:0c:17:ec:07:70:f5:cf:84:2e:40:83:9c:e8:
                    3f:41:6d:3b:ad:d3:a4:14:59:36:78:9d:03:43:ee:
                    10:13:6c:72:de:ae:88:a7:a1:6b:b5:43:ce:67:dc:
                    23:ff:03:1c:a3:e2:3e
                ASN1 OID: secp384r1
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Subject Key Identifier: 
                B3:DB:48:A4:F9:A1:C5:D8:AE:36:41:CC:11:63:69:62:29:BC:4B:C6
    Signature Algorithm: ecdsa-with-SHA384
         30:65:02:31:00:ad:bc:f2:6c:3f:12:4a:d1:2d:39:c3:0a:09:
         97:73:f4:88:36:8c:88:27:bb:e6:88:8d:50:85:a7:63:f9:9e:
         32:de:66:93:0f:f1:cc:b1:09:8f:dd:6c:ab:fa:6b:7f:a0:02:
         30:39:66:5b:c2:64:8d:b8:9e:50:dc:a8:d5:49:a2:ed:c7:dc:
         d1:49:7f:17:01:b8:c8:86:8f:4e:8c:88:2b:a8:9a:a9:8a:c5:
         d1:00:bd:f8:54:e2:9a:e5:5b:7c:b3:27:17
-----BEGIN CERTIFICATE-----
MIICPzCCAcWgAwIBAgIQBVVWvPJepDU1w6QP1atFcjAKBggqhkjOPQQDAzBhMQsw
CQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cu
ZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBHMzAe
Fw0xMzA4MDExMjAwMDBaFw0zODAxMTUxMjAwMDBaMGExCzAJBgNVBAYTAlVTMRUw
EwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5jb20x
IDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IEczMHYwEAYHKoZIzj0CAQYF
K4EEACIDYgAE3afZu4q4C/sLfyHS8L6+c/MzXRq8NOrexpu80JX28MzQC7phW1FG
fp4tn+6OYwwX7Adw9c+ELkCDnOg/QW07rdOkFFk2eJ0DQ+4QE2xy3q6Ip6FrtUPO
Z9wj/wMco+I+o0IwQDAPBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAd
BgNVHQ4EFgQUs9tIpPmhxdiuNkHMEWNpYim8S8YwCgYIKoZIzj0EAwMDaAAwZQIx
AK288mw/EkrRLTnDCgmXc/SINoyIJ7vmiI1Qhadj+Z4y3maTD/HMsQmP3Wyr+mt/
oAIwOWZbwmSNuJ5Q3KjVSaLtx9zRSX8XAbjIho9OjIgrqJqpisXRAL34VOKa5Vt8
sycX
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            05:9b:1b:57:9e:8e:21:32:e2:39:07:bd:a7:77:75:5c
    Signature Algorithm: sha384WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Trusted Root G4
        Validity
            Not Before: Aug  1 12:00:00 2013 GMT
            Not After : Jan 15 12:00:00 2038 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Trusted Root G4
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (4096 bit)
                Modulus:
                    00:bf:e6:90:73:68:de:bb:e4:5d:4a:3c:30:22:30:
                    69:33:ec:c2:a7:25:2e:c9:21:3d:f2:8a:d8:59:c2:
                    e1:29:a7:3d:58:ab:76:9a:cd:ae:7b:1b:84:0d:c4:
                    30:1f:f3:1b:a4:38:16:eb:56:c6:97:6d:1d:ab:b2:
                    79:f2:ca:11:d2:e4:5f:d6:05:3c:52:0f:52:1f:c6:
                    9e:15:a5:7e:be:9f:a9:57:16:59:55:72:af:68:93:
                    70:c2:b2:ba:75:99:6a:73:32:94:d1:10:44:10:2e:
                    df:82:f3:07:84:e6:74:3b:6d:71:e2:2d:0c:1b:ee:
                    20:d5:c9:20:1d:63:29:2d:ce:ec:5e:4e:c8:93:f8:
                    21:61:9b:34:eb:05:c6:5e:ec:5b:1a:bc:eb:c9:cf:
                    cd:ac:34:40:5f:b1:7a:66:ee:77:c8:48:a8:66:57:
                    57:9f:54:58:8e:0c:2b:b7:4f:a7:30:d9:56:ee:ca:
                    7b:5d:e3:ad:c9:4f:5e:e5:35:e7:31:cb:da:93:5e:
                    dc:8e:8f:80:da:b6:91:98:40:90:79:c3:78:c7:b6:
                    b1:c4:b5:6a:18:38:03:10:8d:d8:d4:37:a4:2e:05:
                    7d:88:f5:82:3e:10:91:70:ab:55:82:41:32:d7:db:
                    04:73:2a:6e:91:01:7c:21:4c:d4:bc:ae:1b:03:75:
                    5d:78:66:d9:3a:31:44:9a:33:40:bf:08:d7:5a:49:
                    a4:c2:e6:a9:a0:67:dd:a4:27:bc:a1:4f:39:b5:11:
                    58:17:f7:24:5c:46:8f:64:f7:c1:69:88:76:98:76:
                    3d:59:5d:42:76:87:89:97:69:7a:48:f0:e0:a2:12:
                    1b:66:9a:74:ca:de:4b:1e:e7:0e:63:ae:e6:d4:ef:
                    92:92:3a:9e:3d:dc:00:e4:45:25:89:b6:9a:44:19:
                    2b:7e:c0:94:b4:d2:61:6d:eb:33:d9:c5:df:4b:04:
                    00:cc:7d:1c:95:c3:8f:f7:21:b2:b2:11:b7:bb:7f:
                    f2:d5:8c:70:2c:41:60:aa:b1:63:18:44:95:1a:76:
                    62:7e:f6:80:b0:fb:e8:64:a6:33:d1:89:07:e1:bd:
                    b7:e6:43:a4:18:b8:a6:77:01:e1:0f:94:0c:21:1d:
                    b2:54:29:25:89:6c:e5:0e:52:51:47:74:be:26:ac:
                    b6:41:75:de:7a:ac:5f:8d:3f:c9:bc:d3:41:11:12:
                    5b:e5:10:50:eb:31:c5:ca:72:16:22:09:df:7c:4c:
                    75:3f:63:ec:21:5f:c4:20:51:6b:6f:b1:ab:86:8b:
                    4f:c2:d6:45:5f:9d:20:fc:a1:1e:c5:c0:8f:a2:b1:
                    7e:0a:26:99:f5:e4:69:2f:98:1d:2d:f5:d9:a9:b2:
                    1d:e5:1b
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Subject Key Identifier: 
                EC:D7:E3:82:D2:71:5D:64:4C:DF:2E:67:3F:E7:BA:98:AE:1C:0F:4F
    Signature Algorithm: sha384WithRSAEncryption
         bb:61:d9:7d:a9:6c:be:17:c4:91:1b:c3:a1:a2:00:8d:e3:64:
         68:0f:56:cf:77:ae:70:f9:fd:9a:4a:99:b9:c9:78:5c:0c:0c:
         5f:e4:e6:14:29:56:0b:36:49:5d:44:63:e0:ad:9c:96:18:66:
         1b:23:0d:3d:79:e9:6d:6b:d6:54:f8:d2:3c:c1:43:40:ae:1d:
         50:f5:52:fc:90:3b:bb:98:99:69:6b:c7:c1:a7:a8:68:a4:27:
         dc:9d:f9:27:ae:30:85:b9:f6:67:4d:3a:3e:8f:59:39:22:53:
         44:eb:c8:5d:03:ca:ed:50:7a:7d:62:21:0a:80:c8:73:66:d1:
         a0:05:60:5f:e8:a5:b4:a7:af:a8:f7:6d:35:9c:7c:5a:8a:d6:
         a2:38:99:f3:78:8b:f4:4d:d2:20:0b:de:04:ee:8c:9b:47:81:
         72:0d:c0:14:32:ef:30:59:2e:ae:e0:71:f2:56:e4:6a:97:6f:
         92:50:6d:96:8d:68:7a:9a:b2:36:14:7a:06:f2:24:b9:09:11:
         50:d7:08:b1:b8:89:7a:84:23:61:42:29:e5:a3:cd:a2:20:41:
         d7:d1:9c:64:d9:ea:26:a1:8b:14:d7:4c:19:b2:50:41:71:3d:
         3f:4d:70:23:86:0c:4a:dc:81:d2:cc:32:94:84:0d:08:09:97:
         1c:4f:c0:ee:6b:20:74:30:d2:e0:39:34:10:85:21:15:01:08:
         e8:55:32:de:71:49:d9:28:17:50:4d:e6:be:4d:d1:75:ac:d0:
         ca:fb:41:b8:43:a5:aa:d3:c3:05:44:4f:2c:36:9b:e2:fa:e2:
         45:b8:23:53:6c:06:6f:67:55:7f:46:b5:4c:3f:6e:28:5a:79:
         26:d2:a4:a8:62:97:d2:1e:e2:ed:4a:8b:bc:1b:fd:47:4a:0d:
         df:67:66:7e:b2:5b:41:d0:3b:e4:f4:3b:f4:04:63:e9:ef:c2:
         54:00:51:a0:8a:2a:c9:ce:78:cc:d5:ea:87:04:18:b3:ce:af:
         49:88:af:f3:92:99:b6:b3:e6:61:0f:d2:85:00:e7:50:1a:e4:
         1b:95:9d:19:a1:b9:9c:b1:9b:b1:00:1e:ef:d0:0f:4f:42:6c:
         c9:0a:bc:ee:43:fa:3a:71:a5:c8:4d:26:a5:35:fd:89:5d:bc:
         85:62:1d:32:d2:a0:2b:54:ed:9a:57:c1:db:fa:10:cf:19:b7:
         8b:4a:1b:8f:01:b6:27:95:53:e8:b6:89:6d:5b:bc:68:d4:23:
         e8:8b:51:a2:56:f9:f0:a6:80:a0:d6:1e:b3:bc:0f:0f:53:75:
         29:aa:ea:13:77:e4:de:8c:81:21:ad:07:10:47:11:ad:87:3d:
         07:d1:75:bc:cf:f3:66:7e
-----BEGIN CERTIFICATE-----
MIIFkDCCA3igAwIBAgIQBZsbV56OITLiOQe9p3d1XDANBgkqhkiG9w0BAQwFADBi
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSEwHwYDVQQDExhEaWdpQ2VydCBUcnVzdGVkIFJvb3Qg
RzQwHhcNMTMwODAxMTIwMDAwWhcNMzgwMTE1MTIwMDAwWjBiMQswCQYDVQQGEwJV
UzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNlcnQu
Y29tMSEwHwYDVQQDExhEaWdpQ2VydCBUcnVzdGVkIFJvb3QgRzQwggIiMA0GCSqG
SIb3DQEBAQUAA4ICDwAwggIKAoICAQC/5pBzaN675F1KPDAiMGkz7MKnJS7JIT3y
ithZwuEppz1Yq3aaza57G4QNxDAf8xukOBbrVsaXbR2rsnnyyhHS5F/WBTxSD1If
xp4VpX6+n6lXFllVcq9ok3DCsrp1mWpzMpTREEQQLt+C8weE5nQ7bXHiLQwb7iDV
ySAdYyktzuxeTsiT+CFhmzTrBcZe7FsavOvJz82sNEBfsXpm7nfISKhmV1efVFiO
DCu3T6cw2Vbuyntd463JT17lNecxy9qTXtyOj4DatpGYQJB5w3jHtrHEtWoYOAMQ
jdjUN6QuBX2I9YI+EJFwq1WCQTLX2wRzKm6RAXwhTNS8rhsDdV14Ztk6MUSaM0C/
CNdaSaTC5qmgZ92kJ7yhTzm1EVgX9yRcRo9k98FpiHaYdj1ZXUJ2h4mXaXpI8OCi
EhtmmnTK3kse5w5jrubU75KSOp493ADkRSWJtppEGSt+wJS00mFt6zPZxd9LBADM
fRyVw4/3IbKyEbe7f/LVjHAsQWCqsWMYRJUadmJ+9oCw++hkpjPRiQfhvbfmQ6QY
uKZ3AeEPlAwhHbJUKSWJbOUOUlFHdL4mrLZBdd56rF+NP8m800ERElvlEFDrMcXK
chYiCd98THU/Y+whX8QgUWtvsauGi0/C1kVfnSD8oR7FwI+isX4KJpn15GkvmB0t
9dmpsh3lGwIDAQABo0IwQDAPBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIB
hjAdBgNVHQ4EFgQU7NfjgtJxXWRM3y5nP+e6mK4cD08wDQYJKoZIhvcNAQEMBQAD
ggIBALth2X2pbL4XxJEbw6GiAI3jZGgPVs93rnD5/ZpKmbnJeFwMDF/k5hQpVgs2
SV1EY+CtnJYYZhsjDT156W1r1lT40jzBQ0CuHVD1UvyQO7uYmWlrx8GnqGikJ9yd
+SeuMIW59mdNOj6PWTkiU0TryF0Dyu1Qen1iIQqAyHNm0aAFYF/opbSnr6j3bTWc
fFqK1qI4mfN4i/RN0iAL3gTujJtHgXINwBQy7zBZLq7gcfJW5GqXb5JQbZaNaHqa
sjYUegbyJLkJEVDXCLG4iXqEI2FCKeWjzaIgQdfRnGTZ6iahixTXTBmyUEFxPT9N
cCOGDErcgdLMMpSEDQgJlxxPwO5rIHQw0uA5NBCFIRUBCOhVMt5xSdkoF1BN5r5N
0XWs0Mr7QbhDparTwwVETyw2m+L64kW4I1NsBm9nVX9GtUw/bihaeSbSpKhil9Ie
4u1Ki7wb/UdKDd9nZn6yW0HQO+T0O/QEY+nvwlQAUaCKKsnOeMzV6ocEGLPOr0mI
r/OSmbaz5mEP0oUA51Aa5BuVnRmhuZyxm7EAHu/QD09CbMkKvO5D+jpxpchNJqU1
/YldvIViHTLSoCtU7ZpXwdv6EM8Zt4tKG48BtieVU+i2iW1bvGjUI+iLUaJW+fCm
gKDWHrO8Dw9TdSmq6hN35N6MgSGtBxBHEa2HPQfRdbzP82Z+
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            06:9e:1d:b7:7f:cf:1d:fb:a9:7a:f5:e5:c9:a2:40:37
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root CA
        Validity
            Not Before: Mar  8 12:00:00 2013 GMT
            Not After : Mar  8 12:00:00 2023 GMT
        Subject: C=US, O=DigiCert Inc, CN=DigiCert Secure Server CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:bb:57:e4:21:a9:d5:9b:60:37:7e:8e:a1:61:7f:
                    81:e2:1a:c2:75:64:d9:91:50:0b:e4:36:44:24:6e:
                    30:d2:9b:7a:27:fa:c2:6a:ae:6a:70:09:38:b9:20:
                    0a:c8:65:10:4a:88:ac:31:f2:dc:92:f2:63:a1:5d:
                    80:63:59:80:92:23:1c:e6:ef:76:4a:50:35:c9:d8:
                    71:38:b9:ed:f0:e6:42:ae:d3:38:26:79:30:f9:22:
                    94:c6:db:a6:3f:41:78:90:d8:de:5c:7e:69:7d:f8:
                    90:15:3a:d0:a1:a0:be:fa:b2:b2:19:a1:d8:2b:d1:
                    ce:bf:6b:dd:49:ab:a3:92:fe:b5:ab:c8:c1:3e:ee:
                    01:00:d8:a9:44:b8:42:73:88:c3:61:f5:ab:4a:83:
                    28:0a:d2:d4:49:fa:6a:b1:cd:df:57:2c:94:e5:e2:
                    ca:83:5f:b7:ba:62:5c:2f:68:a5:f0:c0:b9:fd:2b:
                    d1:e9:1f:d8:1a:62:15:bd:ff:3d:a6:f7:cb:ef:e6:
                    db:65:2f:25:38:ec:fb:e6:20:66:58:96:34:19:d2:
                    15:ce:21:d3:24:cc:d9:14:6f:d8:fe:55:c7:e7:6f:
                    b6:0f:1a:8c:49:be:29:f2:ba:5a:9a:81:26:37:24:
                    6f:d7:48:12:6c:2e:59:f5:9c:18:bb:d9:f6:68:e2:
                    df:45
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertGlobalRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertGlobalRootCA.crl

            X509v3 Certificate Policies: 
                Policy: X509v3 Any Policy
                  CPS: https://www.digicert.com/CPS

            X509v3 Subject Key Identifier: 
                90:71:DB:37:EB:73:C8:EF:DC:D5:1E:12:B6:34:BA:2B:5A:A0:A6:92
            X509v3 Authority Key Identifier: 
                keyid:03:DE:50:35:56:D1:4C:BB:66:F0:A3:E2:1B:1B:C3:97:B2:3D:D1:55

    Signature Algorithm: sha1WithRSAEncryption
         30:ce:d1:95:51:00:ae:06:0b:a1:0e:02:c0:17:ac:b6:7f:8f:
         20:f6:40:75:74:1c:cc:78:b1:a4:4f:ea:f4:d0:c4:9d:a2:de:
         81:07:26:1f:40:88:51:f0:1f:cf:b7:4c:40:99:d0:f4:3c:71:
         98:73:88:97:2c:19:d7:6e:84:8f:a4:1f:9c:5a:20:e3:51:5c:
         b0:c5:9e:99:6a:4f:c8:69:f7:10:ff:4e:ad:19:d9:c9:58:b3:
         33:ae:0c:d9:96:29:9e:71:b2:70:63:a3:b6:99:16:42:1d:65:
         f3:f7:a0:1e:7d:c5:d4:65:14:b2:62:84:d4:6c:5c:08:0c:d8:
         6c:93:2b:b4:76:59:8a:d1:7f:ff:03:d8:c2:5d:b8:2f:22:d6:
         38:f0:f6:9c:6b:7d:46:eb:99:74:f7:eb:4a:0e:a9:a6:04:eb:
         7b:ce:f0:5c:6b:98:31:5a:98:40:eb:69:c4:05:f4:20:a8:ca:
         08:3a:65:6c:38:15:f5:5c:2c:b2:55:e4:2c:6b:41:f0:be:5c:
         46:ca:4a:29:a0:48:5e:20:d2:45:ff:05:de:34:af:70:4b:81:
         39:e2:ca:07:57:7c:b6:31:dc:21:29:e2:be:97:0e:77:90:14:
         51:40:e1:bf:e3:cc:1b:19:9c:25:ca:a7:06:b2:53:df:23:b2:
         cf:12:19:a3
-----BEGIN CERTIFICATE-----
MIIEjzCCA3egAwIBAgIQBp4dt3/PHfupevXlyaJANzANBgkqhkiG9w0BAQUFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0xMzAzMDgxMjAwMDBaFw0yMzAzMDgxMjAwMDBaMEgxCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxIjAgBgNVBAMTGURpZ2lDZXJ0IFNlY3Vy
ZSBTZXJ2ZXIgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC7V+Qh
qdWbYDd+jqFhf4HiGsJ1ZNmRUAvkNkQkbjDSm3on+sJqrmpwCTi5IArIZRBKiKwx
8tyS8mOhXYBjWYCSIxzm73ZKUDXJ2HE4ue3w5kKu0zgmeTD5IpTG26Y/QXiQ2N5c
fml9+JAVOtChoL76srIZodgr0c6/a91Jq6OS/rWryME+7gEA2KlEuEJziMNh9atK
gygK0tRJ+mqxzd9XLJTl4sqDX7e6YlwvaKXwwLn9K9HpH9gaYhW9/z2m98vv5ttl
LyU47PvmIGZYljQZ0hXOIdMkzNkUb9j+Vcfnb7YPGoxJvinyulqagSY3JG/XSBJs
Lln1nBi72fZo4t9FAgMBAAGjggFaMIIBVjASBgNVHRMBAf8ECDAGAQH/AgEAMA4G
A1UdDwEB/wQEAwIBhjA0BggrBgEFBQcBAQQoMCYwJAYIKwYBBQUHMAGGGGh0dHA6
Ly9vY3NwLmRpZ2ljZXJ0LmNvbTB7BgNVHR8EdDByMDegNaAzhjFodHRwOi8vY3Js
My5kaWdpY2VydC5jb20vRGlnaUNlcnRHbG9iYWxSb290Q0EuY3JsMDegNaAzhjFo
dHRwOi8vY3JsNC5kaWdpY2VydC5jb20vRGlnaUNlcnRHbG9iYWxSb290Q0EuY3Js
MD0GA1UdIAQ2MDQwMgYEVR0gADAqMCgGCCsGAQUFBwIBFhxodHRwczovL3d3dy5k
aWdpY2VydC5jb20vQ1BTMB0GA1UdDgQWBBSQcds363PI79zVHhK2NLorWqCmkjAf
BgNVHSMEGDAWgBQD3lA1VtFMu2bwo+IbG8OXsj3RVTANBgkqhkiG9w0BAQUFAAOC
AQEAMM7RlVEArgYLoQ4CwBestn+PIPZAdXQczHixpE/q9NDEnaLegQcmH0CIUfAf
z7dMQJnQ9DxxmHOIlywZ126Ej6QfnFog41FcsMWemWpPyGn3EP9OrRnZyVizM64M
2ZYpnnGycGOjtpkWQh1l8/egHn3F1GUUsmKE1GxcCAzYbJMrtHZZitF//wPYwl24
LyLWOPD2nGt9RuuZdPfrSg6ppgTre87wXGuYMVqYQOtpxAX0IKjKCDplbDgV9Vws
slXkLGtB8L5cRspKKaBIXiDSRf8F3jSvcEuBOeLKB1d8tjHcISnivpcOd5AUUUDh
v+PMGxmcJcqnBrJT3yOyzxIZow==
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            01:fd:a3:eb:6e:ca:75:c8:88:43:8b:72:4b:cf:bc:91
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root CA
        Validity
            Not Before: Mar  8 12:00:00 2013 GMT
            Not After : Mar  8 12:00:00 2023 GMT
        Subject: C=US, O=DigiCert Inc, CN=DigiCert SHA2 Secure Server CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:dc:ae:58:90:4d:c1:c4:30:15:90:35:5b:6e:3c:
                    82:15:f5:2c:5c:bd:e3:db:ff:71:43:fa:64:25:80:
                    d4:ee:18:a2:4d:f0:66:d0:0a:73:6e:11:98:36:17:
                    64:af:37:9d:fd:fa:41:84:af:c7:af:8c:fe:1a:73:
                    4d:cf:33:97:90:a2:96:87:53:83:2b:b9:a6:75:48:
                    2d:1d:56:37:7b:da:31:32:1a:d7:ac:ab:06:f4:aa:
                    5d:4b:b7:47:46:dd:2a:93:c3:90:2e:79:80:80:ef:
                    13:04:6a:14:3b:b5:9b:92:be:c2:07:65:4e:fc:da:
                    fc:ff:7a:ae:dc:5c:7e:55:31:0c:e8:39:07:a4:d7:
                    be:2f:d3:0b:6a:d2:b1:df:5f:fe:57:74:53:3b:35:
                    80:dd:ae:8e:44:98:b3:9f:0e:d3:da:e0:d7:f4:6b:
                    29:ab:44:a7:4b:58:84:6d:92:4b:81:c3:da:73:8b:
                    12:97:48:90:04:45:75:1a:dd:37:31:97:92:e8:cd:
                    54:0d:3b:e4:c1:3f:39:5e:2e:b8:f3:5c:7e:10:8e:
                    86:41:00:8d:45:66:47:b0:a1:65:ce:a0:aa:29:09:
                    4e:f3:97:eb:e8:2e:ab:0f:72:a7:30:0e:fa:c7:f4:
                    fd:14:77:c3:a4:5b:28:57:c2:b3:f9:82:fd:b7:45:
                    58:9b
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertGlobalRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertGlobalRootCA.crl

            X509v3 Certificate Policies: 
                Policy: X509v3 Any Policy
                  CPS: https://www.digicert.com/CPS

            X509v3 Subject Key Identifier: 
                0F:80:61:1C:82:31:61:D5:2F:28:E7:8D:46:38:B4:2C:E1:C6:D9:E2
            X509v3 Authority Key Identifier: 
                keyid:03:DE:50:35:56:D1:4C:BB:66:F0:A3:E2:1B:1B:C3:97:B2:3D:D1:55

    Signature Algorithm: sha256WithRSAEncryption
         23:3e:df:4b:d2:31:42:a5:b6:7e:42:5c:1a:44:cc:69:d1:68:
         b4:5d:4b:e0:04:21:6c:4b:e2:6d:cc:b1:e0:97:8f:a6:53:09:
         cd:aa:2a:65:e5:39:4f:1e:83:a5:6e:5c:98:a2:24:26:e6:fb:
         a1:ed:93:c7:2e:02:c6:4d:4a:bf:b0:42:df:78:da:b3:a8:f9:
         6d:ff:21:85:53:36:60:4c:76:ce:ec:38:dc:d6:51:80:f0:c5:
         d6:e5:d4:4d:27:64:ab:9b:c7:3e:71:fb:48:97:b8:33:6d:c9:
         13:07:ee:96:a2:1b:18:15:f6:5c:4c:40:ed:b3:c2:ec:ff:71:
         c1:e3:47:ff:d4:b9:00:b4:37:42:da:20:c9:ea:6e:8a:ee:14:
         06:ae:7d:a2:59:98:88:a8:1b:6f:2d:f4:f2:c9:14:5f:26:cf:
         2c:8d:7e:ed:37:c0:a9:d5:39:b9:82:bf:19:0c:ea:34:af:00:
         21:68:f8:ad:73:e2:c9:32:da:38:25:0b:55:d3:9a:1d:f0:68:
         86:ed:2e:41:34:ef:7c:a5:50:1d:bf:3a:f9:d3:c1:08:0c:e6:
         ed:1e:8a:58:25:e4:b8:77:ad:2d:6e:f5:52:dd:b4:74:8f:ab:
         49:2e:9d:3b:93:34:28:1f:78:ce:94:ea:c7:bd:d3:c9:6d:1c:
         de:5c:32:f3
-----BEGIN CERTIFICATE-----
MIIElDCCA3ygAwIBAgIQAf2j627KdciIQ4tyS8+8kTANBgkqhkiG9w0BAQsFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0xMzAzMDgxMjAwMDBaFw0yMzAzMDgxMjAwMDBaME0xCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxJzAlBgNVBAMTHkRpZ2lDZXJ0IFNIQTIg
U2VjdXJlIFNlcnZlciBDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
ANyuWJBNwcQwFZA1W248ghX1LFy949v/cUP6ZCWA1O4Yok3wZtAKc24RmDYXZK83
nf36QYSvx6+M/hpzTc8zl5CilodTgyu5pnVILR1WN3vaMTIa16yrBvSqXUu3R0bd
KpPDkC55gIDvEwRqFDu1m5K+wgdlTvza/P96rtxcflUxDOg5B6TXvi/TC2rSsd9f
/ld0Uzs1gN2ujkSYs58O09rg1/RrKatEp0tYhG2SS4HD2nOLEpdIkARFdRrdNzGX
kujNVA075ME/OV4uuPNcfhCOhkEAjUVmR7ChZc6gqikJTvOX6+guqw9ypzAO+sf0
/RR3w6RbKFfCs/mC/bdFWJsCAwEAAaOCAVowggFWMBIGA1UdEwEB/wQIMAYBAf8C
AQAwDgYDVR0PAQH/BAQDAgGGMDQGCCsGAQUFBwEBBCgwJjAkBggrBgEFBQcwAYYY
aHR0cDovL29jc3AuZGlnaWNlcnQuY29tMHsGA1UdHwR0MHIwN6A1oDOGMWh0dHA6
Ly9jcmwzLmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEdsb2JhbFJvb3RDQS5jcmwwN6A1
oDOGMWh0dHA6Ly9jcmw0LmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEdsb2JhbFJvb3RD
QS5jcmwwPQYDVR0gBDYwNDAyBgRVHSAAMCowKAYIKwYBBQUHAgEWHGh0dHBzOi8v
d3d3LmRpZ2ljZXJ0LmNvbS9DUFMwHQYDVR0OBBYEFA+AYRyCMWHVLyjnjUY4tCzh
xtniMB8GA1UdIwQYMBaAFAPeUDVW0Uy7ZvCj4hsbw5eyPdFVMA0GCSqGSIb3DQEB
CwUAA4IBAQAjPt9L0jFCpbZ+QlwaRMxp0Wi0XUvgBCFsS+JtzLHgl4+mUwnNqipl
5TlPHoOlblyYoiQm5vuh7ZPHLgLGTUq/sELfeNqzqPlt/yGFUzZgTHbO7Djc1lGA
8MXW5dRNJ2Srm8c+cftIl7gzbckTB+6WohsYFfZcTEDts8Ls/3HB40f/1LkAtDdC
2iDJ6m6K7hQGrn2iWZiIqBtvLfTyyRRfJs8sjX7tN8Cp1Tm5gr8ZDOo0rwAhaPit
c+LJMto4JQtV05od8GiG7S5BNO98pVAdvzr508EIDObtHopYJeS4d60tbvVS3bR0
j6tJLp07kzQoH3jOlOrHvdPJbRzeXDLz
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0a:cb:28:ba:46:5e:e5:39:08:76:74:70:f3:cd:c6:12
    Signature Algorithm: sha384WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root CA
        Validity
            Not Before: Mar  8 12:00:00 2013 GMT
            Not After : Mar  8 12:00:00 2023 GMT
        Subject: C=US, O=DigiCert Inc, CN=DigiCert ECC Secure Server CA
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (384 bit)
                pub: 
                    04:e2:08:42:ea:77:d8:24:de:a0:2c:64:a4:13:ce:
                    40:9c:23:72:a9:02:0a:0e:37:3f:21:36:b8:8d:53:
                    14:f6:d5:91:95:4b:f3:96:02:8d:71:1e:c4:d8:cb:
                    a7:9f:5e:ef:a0:e6:7f:5a:92:11:96:53:6f:eb:c0:
                    cb:3f:ae:fd:5b:3f:47:24:e7:9a:07:2e:96:be:a8:
                    2f:bb:57:18:af:71:a4:bd:78:3a:1e:e8:5b:3c:6b:
                    64:11:2b:cc:34:2b:8c
                ASN1 OID: secp384r1
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertGlobalRootCA.crl

            X509v3 Certificate Policies: 
                Policy: X509v3 Any Policy
                  CPS: https://www.digicert.com/CPS

            X509v3 Subject Key Identifier: 
                A3:9D:E6:1F:F9:DA:39:4F:C0:6E:E8:91:CB:95:A5:DA:31:E2:0A:9F
            X509v3 Authority Key Identifier: 
                keyid:03:DE:50:35:56:D1:4C:BB:66:F0:A3:E2:1B:1B:C3:97:B2:3D:D1:55

    Signature Algorithm: sha384WithRSAEncryption
         c7:8a:a0:43:4b:ec:74:c9:c5:ab:d5:1f:30:35:36:6e:98:56:
         7b:48:ac:05:63:ae:7b:9a:57:24:57:cc:6f:fa:de:ab:6d:9c:
         c7:b6:ba:ec:ce:e7:ca:73:64:db:df:04:37:0a:00:49:b3:3f:
         9f:26:ad:91:8c:20:e8:1f:88:0e:2a:fb:66:37:c8:30:e8:d2:
         c2:24:a7:45:48:2d:ea:01:50:4a:31:94:13:df:8d:5f:da:2a:
         bb:49:3c:61:f3:79:c8:9c:66:92:1a:96:2a:f4:7b:36:58:a3:
         2c:41:10:74:1a:d3:ed:48:b6:d2:bb:8a:06:45:71:33:10:30:
         7a:7a:98:21:dd:24:b9:ec:9c:b5:92:07:ad:83:c6:c4:6a:f8:
         77:e6:35:be:13:0f:27:64:b2:43:bf:83:e9:77:56:db:08:87:
         94:47:14:f5:5f:28:af:a3:68:4c:83:8f:60:f7:96:80:79:85:
         6a:76:26:9d:95:0c:20:03:8d:3e:ee:7a:28:65:64:66:a4:d9:
         83:ea:99:74:cd:6e:4d:7d:1c:eb:8d:b2:c5:af:16:1b:4e:c8:
         f3:55:ea:88:38:11:34:1d:11:af:3f:07:a8:4f:6a:d2:74:11:
         2f:2a:fc:73:b7:5f:c2:15:43:05:6c:d6:7d:da:02:bd:22:9b:
         4f:d3:f9:77
-----BEGIN CERTIFICATE-----
MIIDrDCCApSgAwIBAgIQCssoukZe5TkIdnRw883GEjANBgkqhkiG9w0BAQwFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0xMzAzMDgxMjAwMDBaFw0yMzAzMDgxMjAwMDBaMEwxCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxJjAkBgNVBAMTHURpZ2lDZXJ0IEVDQyBT
ZWN1cmUgU2VydmVyIENBMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE4ghC6nfYJN6g
LGSkE85AnCNyqQIKDjc/ITa4jVMU9tWRlUvzlgKNcR7E2Munn17voOZ/WpIRllNv
68DLP679Wz9HJOeaBy6Wvqgvu1cYr3GkvXg6HuhbPGtkESvMNCuMo4IBITCCAR0w
EgYDVR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwNAYIKwYBBQUHAQEE
KDAmMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2VydC5jb20wQgYDVR0f
BDswOTA3oDWgM4YxaHR0cDovL2NybDMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0R2xv
YmFsUm9vdENBLmNybDA9BgNVHSAENjA0MDIGBFUdIAAwKjAoBggrBgEFBQcCARYc
aHR0cHM6Ly93d3cuZGlnaWNlcnQuY29tL0NQUzAdBgNVHQ4EFgQUo53mH/naOU/A
buiRy5Wl2jHiCp8wHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUwDQYJ
KoZIhvcNAQEMBQADggEBAMeKoENL7HTJxavVHzA1Nm6YVntIrAVjrnuaVyRXzG/6
3qttnMe2uuzO58pzZNvfBDcKAEmzP58mrZGMIOgfiA4q+2Y3yDDo0sIkp0VILeoB
UEoxlBPfjV/aKrtJPGHzecicZpIalir0ezZYoyxBEHQa0+1IttK7igZFcTMQMHp6
mCHdJLnsnLWSB62DxsRq+HfmNb4TDydkskO/g+l3VtsIh5RHFPVfKK+jaEyDj2D3
loB5hWp2Jp2VDCADjT7ueihlZGak2YPqmXTNbk19HOuNssWvFhtOyPNV6og4ETQd
Ea8/B6hPatJ0ES8q/HO3X8IVQwVs1n3aAr0im0/T+Xc=
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            03:37:b9:28:34:7c:60:a6:ae:c5:ad:b1:21:7f:38:60
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Validity
            Not Before: Nov  9 12:00:00 2007 GMT
            Not After : Nov 10 00:00:00 2021 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV CA-1
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:f3:96:62:d8:75:6e:19:ff:3f:34:7c:49:4f:31:
                    7e:0d:04:4e:99:81:e2:b3:85:55:91:30:b1:c0:af:
                    70:bb:2c:a8:e7:18:aa:3f:78:f7:90:68:52:86:01:
                    88:97:e2:3b:06:65:90:aa:bd:65:76:c2:ec:be:10:
                    5b:37:78:83:60:75:45:c6:bd:74:aa:b6:9f:a4:3a:
                    01:50:17:c4:39:69:b9:f1:4f:ef:82:c1:ca:f3:4a:
                    db:cc:9e:50:4f:4d:40:a3:3a:90:e7:86:66:bc:f0:
                    3e:76:28:4c:d1:75:80:9e:6a:35:14:35:03:9e:db:
                    0c:8c:c2:28:ad:50:b2:ce:f6:91:a3:c3:a5:0a:58:
                    49:f6:75:44:6c:ba:f9:ce:e9:ab:3a:02:e0:4d:f3:
                    ac:e2:7a:e0:60:22:05:3c:82:d3:52:e2:f3:9c:47:
                    f8:3b:d8:b2:4b:93:56:4a:bf:70:ab:3e:e9:68:c8:
                    1d:8f:58:1d:2a:4d:5e:27:3d:ad:0a:59:2f:5a:11:
                    20:40:d9:68:04:68:2d:f4:c0:84:0b:0a:1b:78:df:
                    ed:1a:58:dc:fb:41:5a:6d:6b:f2:ed:1c:ee:5c:32:
                    b6:5c:ec:d7:a6:03:32:a6:e8:de:b7:28:27:59:88:
                    80:ff:7b:ad:89:58:d5:1e:14:a4:f2:b0:70:d4:a0:
                    3e:a7
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication, TLS Web Client Authentication, Code Signing, E-mail Protection, Time Stamping
            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.2.1
                  CPS: http://www.digicert.com/ssl-cps-repository.htm
                  User Notice:
                    Explicit Text: 

            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com
                CA Issuers - URI:http://www.digicert.com/CACerts/DigiCertHighAssuranceEVRootCA.crt

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertHighAssuranceEVRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertHighAssuranceEVRootCA.crl

            X509v3 Subject Key Identifier: 
                4C:58:CB:25:F0:41:4F:52:F4:28:C8:81:43:9B:A6:A8:A0:E6:92:E5
            X509v3 Authority Key Identifier: 
                keyid:B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3

    Signature Algorithm: sha1WithRSAEncryption
         4c:7a:17:87:28:5d:17:bc:b2:32:73:bf:cd:2e:f5:58:31:1d:
         f0:b1:71:54:9c:d6:9b:67:93:db:2f:03:3e:16:6f:1e:03:c9:
         53:84:a3:56:60:1e:78:94:1b:a2:a8:6f:a3:a4:8b:52:91:d7:
         dd:5c:95:bb:ef:b5:16:49:e9:a5:42:4f:34:f2:47:ff:ae:81:
         7f:13:54:b7:20:c4:70:15:cb:81:0a:81:cb:74:57:dc:9c:df:
         24:a4:29:0c:18:f0:1c:e4:ae:07:33:ec:f1:49:3e:55:cf:6e:
         4f:0d:54:7b:d3:c9:e8:15:48:d4:c5:bb:dc:35:1c:77:45:07:
         48:45:85:bd:d7:7e:53:b8:c0:16:d9:95:cd:8b:8d:7d:c9:60:
         4f:d1:a2:9b:e3:d0:30:d6:b4:73:36:e6:d2:f9:03:b2:e3:a4:
         f5:e5:b8:3e:04:49:00:ba:2e:a6:4a:72:83:72:9d:f7:0b:8c:
         a9:89:e7:b3:d7:64:1f:d6:e3:60:cb:03:c4:dc:88:e9:9d:25:
         01:00:71:cb:03:b4:29:60:25:8f:f9:46:d1:7b:71:ae:cd:53:
         12:5b:84:8e:c2:0f:c7:ed:93:19:d9:c9:fa:8f:58:34:76:32:
         2f:ae:e1:50:14:61:d4:a8:58:a3:c8:30:13:23:ef:c6:25:8c:
         36:8f:1c:80
-----BEGIN CERTIFICATE-----
MIIG5jCCBc6gAwIBAgIQAze5KDR8YKauxa2xIX84YDANBgkqhkiG9w0BAQUFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTA3MTEwOTEyMDAwMFoXDTIxMTExMDAwMDAwMFowaTEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTEoMCYGA1UEAxMfRGlnaUNlcnQgSGlnaCBBc3N1cmFuY2Ug
RVYgQ0EtMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPOWYth1bhn/
PzR8SU8xfg0ETpmB4rOFVZEwscCvcLssqOcYqj9495BoUoYBiJfiOwZlkKq9ZXbC
7L4QWzd4g2B1Rca9dKq2n6Q6AVAXxDlpufFP74LByvNK28yeUE9NQKM6kOeGZrzw
PnYoTNF1gJ5qNRQ1A57bDIzCKK1Qss72kaPDpQpYSfZ1RGy6+c7pqzoC4E3zrOJ6
4GAiBTyC01Li85xH+DvYskuTVkq/cKs+6WjIHY9YHSpNXic9rQpZL1oRIEDZaARo
LfTAhAsKG3jf7RpY3PtBWm1r8u0c7lwytlzs16YDMqbo3rcoJ1mIgP97rYlY1R4U
pPKwcNSgPqcCAwEAAaOCA4UwggOBMA4GA1UdDwEB/wQEAwIBhjA7BgNVHSUENDAy
BggrBgEFBQcDAQYIKwYBBQUHAwIGCCsGAQUFBwMDBggrBgEFBQcDBAYIKwYBBQUH
AwgwggHEBgNVHSAEggG7MIIBtzCCAbMGCWCGSAGG/WwCATCCAaQwOgYIKwYBBQUH
AgEWLmh0dHA6Ly93d3cuZGlnaWNlcnQuY29tL3NzbC1jcHMtcmVwb3NpdG9yeS5o
dG0wggFkBggrBgEFBQcCAjCCAVYeggFSAEEAbgB5ACAAdQBzAGUAIABvAGYAIAB0
AGgAaQBzACAAQwBlAHIAdABpAGYAaQBjAGEAdABlACAAYwBvAG4AcwB0AGkAdAB1
AHQAZQBzACAAYQBjAGMAZQBwAHQAYQBuAGMAZQAgAG8AZgAgAHQAaABlACAARABp
AGcAaQBDAGUAcgB0ACAARQBWACAAQwBQAFMAIABhAG4AZAAgAHQAaABlACAAUgBl
AGwAeQBpAG4AZwAgAFAAYQByAHQAeQAgAEEAZwByAGUAZQBtAGUAbgB0ACAAdwBo
AGkAYwBoACAAbABpAG0AaQB0ACAAbABpAGEAYgBpAGwAaQB0AHkAIABhAG4AZAAg
AGEAcgBlACAAaQBuAGMAbwByAHAAbwByAGEAdABlAGQAIABoAGUAcgBlAGkAbgAg
AGIAeQAgAHIAZQBmAGUAcgBlAG4AYwBlAC4wEgYDVR0TAQH/BAgwBgEB/wIBADCB
gwYIKwYBBQUHAQEEdzB1MCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2Vy
dC5jb20wTQYIKwYBBQUHMAKGQWh0dHA6Ly93d3cuZGlnaWNlcnQuY29tL0NBQ2Vy
dHMvRGlnaUNlcnRIaWdoQXNzdXJhbmNlRVZSb290Q0EuY3J0MIGPBgNVHR8EgYcw
gYQwQKA+oDyGOmh0dHA6Ly9jcmwzLmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEhpZ2hB
c3N1cmFuY2VFVlJvb3RDQS5jcmwwQKA+oDyGOmh0dHA6Ly9jcmw0LmRpZ2ljZXJ0
LmNvbS9EaWdpQ2VydEhpZ2hBc3N1cmFuY2VFVlJvb3RDQS5jcmwwHQYDVR0OBBYE
FExYyyXwQU9S9CjIgUObpqig5pLlMB8GA1UdIwQYMBaAFLE+w2kD+L9HAdSYJhoI
Au9jZCvDMA0GCSqGSIb3DQEBBQUAA4IBAQBMeheHKF0XvLIyc7/NLvVYMR3wsXFU
nNabZ5PbLwM+Fm8eA8lThKNWYB54lBuiqG+jpItSkdfdXJW777UWSemlQk808kf/
roF/E1S3IMRwFcuBCoHLdFfcnN8kpCkMGPAc5K4HM+zxST5Vz25PDVR708noFUjU
xbvcNRx3RQdIRYW9135TuMAW2ZXNi419yWBP0aKb49Aw1rRzNubS+QOy46T15bg+
BEkAui6mSnKDcp33C4ypieez12Qf1uNgywPE3IjpnSUBAHHLA7QpYCWP+UbRe3Gu
zVMSW4SOwg/H7ZMZ2cn6j1g0djIvruFQFGHUqFijyDATI+/GJYw2jxyA
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0a:5f:11:4d:03:5b:17:91:17:d2:ef:d4:03:8c:3f:3b
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Validity
            Not Before: Apr  2 12:00:00 2008 GMT
            Not After : Apr  3 00:00:00 2022 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance CA-3
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:bf:61:0a:29:10:1f:5e:fe:34:37:51:08:f8:1e:
                    fb:22:ed:61:be:0b:0d:70:4c:50:63:26:75:15:b9:
                    41:88:97:b6:f0:a0:15:bb:08:60:e0:42:e8:05:29:
                    10:87:36:8a:28:65:a8:ef:31:07:74:6d:36:97:2f:
                    28:46:66:04:c7:2a:79:26:7a:99:d5:8e:c3:6d:4f:
                    a0:5e:ad:bc:3d:91:c2:59:7b:5e:36:6c:c0:53:cf:
                    00:08:32:3e:10:64:58:10:13:69:c7:0c:ee:9c:42:
                    51:00:f9:05:44:ee:24:ce:7a:1f:ed:8c:11:bd:12:
                    a8:f3:15:f4:1c:7a:31:69:01:1b:a7:e6:5d:c0:9a:
                    6c:7e:09:9e:e7:52:44:4a:10:3a:23:e4:9b:b6:03:
                    af:a8:9c:b4:5b:9f:d4:4b:ad:92:8c:ce:b5:11:2a:
                    aa:37:18:8d:b4:c2:b8:d8:5c:06:8c:f8:ff:23:bd:
                    35:5e:d4:7c:3e:7e:83:0e:91:96:05:98:c3:b2:1f:
                    e3:c8:65:eb:a9:7b:5d:a0:2c:cc:fc:3c:d9:6d:ed:
                    cc:fa:4b:43:8c:c9:d4:b8:a5:61:1c:b2:40:b6:28:
                    12:df:b9:f8:5f:fe:d3:b2:c9:ef:3d:b4:1e:4b:7c:
                    1c:4c:99:36:9e:3d:eb:ec:a7:68:5e:1d:df:67:6e:
                    5e:fb
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.1.3.0.2
                  CPS: http://www.digicert.com/ssl-cps-repository.htm
                  User Notice:
                    Explicit Text: 

            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertHighAssuranceEVRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertHighAssuranceEVRootCA.crl

            X509v3 Authority Key Identifier: 
                keyid:B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3

            X509v3 Subject Key Identifier: 
                50:EA:73:89:DB:29:FB:10:8F:9E:E5:01:20:D4:DE:79:99:48:83:F7
    Signature Algorithm: sha1WithRSAEncryption
         1e:e2:a5:48:9e:6c:db:53:38:0f:ef:a6:1a:2a:ac:e2:03:43:
         ed:9a:bc:3e:8e:75:1b:f0:fd:2e:22:59:ac:13:c0:61:e2:e7:
         fa:e9:99:cd:87:09:75:54:28:bf:46:60:dc:be:51:2c:92:f3:
         1b:91:7c:31:08:70:e2:37:b9:c1:5b:a8:bd:a3:0b:00:fb:1a:
         15:fd:03:ad:58:6a:c5:c7:24:99:48:47:46:31:1e:92:ef:b4:
         5f:4e:34:c7:90:bf:31:c1:f8:b1:84:86:d0:9c:01:aa:df:8a:
         56:06:ce:3a:e9:0e:ae:97:74:5d:d7:71:9a:42:74:5f:de:8d:
         43:7c:de:e9:55:ed:69:00:cb:05:e0:7a:61:61:33:d1:19:4d:
         f9:08:ee:a0:39:c5:25:35:b7:2b:c4:0f:b2:dd:f1:a5:b7:0e:
         24:c4:26:28:8d:79:77:f5:2f:f0:57:ba:7c:07:d4:e1:fc:cd:
         5a:30:57:7e:86:10:47:dd:31:1f:d7:fc:a2:c2:bf:30:7c:5d:
         24:aa:e8:f9:ae:5f:6a:74:c2:ce:6b:b3:46:d8:21:be:29:d4:
         8e:5e:15:d6:42:4a:e7:32:6f:a4:b1:6b:51:83:58:be:3f:6d:
         c7:fb:da:03:21:cb:6a:16:19:4e:0a:f0:ad:84:ca:5d:94:b3:
         5a:76:f7:61
-----BEGIN CERTIFICATE-----
MIIGWDCCBUCgAwIBAgIQCl8RTQNbF5EX0u/UA4w/OzANBgkqhkiG9w0BAQUFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTA4MDQwMjEyMDAwMFoXDTIyMDQwMzAwMDAwMFowZjEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTElMCMGA1UEAxMcRGlnaUNlcnQgSGlnaCBBc3N1cmFuY2Ug
Q0EtMzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL9hCikQH17+NDdR
CPge+yLtYb4LDXBMUGMmdRW5QYiXtvCgFbsIYOBC6AUpEIc2iihlqO8xB3RtNpcv
KEZmBMcqeSZ6mdWOw21PoF6tvD2Rwll7XjZswFPPAAgyPhBkWBATaccM7pxCUQD5
BUTuJM56H+2MEb0SqPMV9Bx6MWkBG6fmXcCabH4JnudSREoQOiPkm7YDr6ictFuf
1EutkozOtREqqjcYjbTCuNhcBoz4/yO9NV7UfD5+gw6RlgWYw7If48hl66l7XaAs
zPw82W3tzPpLQ4zJ1LilYRyyQLYoEt+5+F/+07LJ7z20Hkt8HEyZNp496+ynaF4d
32duXvsCAwEAAaOCAvowggL2MA4GA1UdDwEB/wQEAwIBhjCCAcYGA1UdIASCAb0w
ggG5MIIBtQYLYIZIAYb9bAEDAAIwggGkMDoGCCsGAQUFBwIBFi5odHRwOi8vd3d3
LmRpZ2ljZXJ0LmNvbS9zc2wtY3BzLXJlcG9zaXRvcnkuaHRtMIIBZAYIKwYBBQUH
AgIwggFWHoIBUgBBAG4AeQAgAHUAcwBlACAAbwBmACAAdABoAGkAcwAgAEMAZQBy
AHQAaQBmAGkAYwBhAHQAZQAgAGMAbwBuAHMAdABpAHQAdQB0AGUAcwAgAGEAYwBj
AGUAcAB0AGEAbgBjAGUAIABvAGYAIAB0AGgAZQAgAEQAaQBnAGkAQwBlAHIAdAAg
AEMAUAAvAEMAUABTACAAYQBuAGQAIAB0AGgAZQAgAFIAZQBsAHkAaQBuAGcAIABQ
AGEAcgB0AHkAIABBAGcAcgBlAGUAbQBlAG4AdAAgAHcAaABpAGMAaAAgAGwAaQBt
AGkAdAAgAGwAaQBhAGIAaQBsAGkAdAB5ACAAYQBuAGQAIABhAHIAZQAgAGkAbgBj
AG8AcgBwAG8AcgBhAHQAZQBkACAAaABlAHIAZQBpAG4AIABiAHkAIAByAGUAZgBl
AHIAZQBuAGMAZQAuMBIGA1UdEwEB/wQIMAYBAf8CAQAwNAYIKwYBBQUHAQEEKDAm
MCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2VydC5jb20wgY8GA1UdHwSB
hzCBhDBAoD6gPIY6aHR0cDovL2NybDMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0SGln
aEFzc3VyYW5jZUVWUm9vdENBLmNybDBAoD6gPIY6aHR0cDovL2NybDQuZGlnaWNl
cnQuY29tL0RpZ2lDZXJ0SGlnaEFzc3VyYW5jZUVWUm9vdENBLmNybDAfBgNVHSME
GDAWgBSxPsNpA/i/RwHUmCYaCALvY2QrwzAdBgNVHQ4EFgQUUOpzidsp+xCPnuUB
INTeeZlIg/cwDQYJKoZIhvcNAQEFBQADggEBAB7ipUiebNtTOA/vphoqrOIDQ+2a
vD6OdRvw/S4iWawTwGHi5/rpmc2HCXVUKL9GYNy+USyS8xuRfDEIcOI3ucFbqL2j
CwD7GhX9A61YasXHJJlIR0YxHpLvtF9ONMeQvzHB+LGEhtCcAarfilYGzjrpDq6X
dF3XcZpCdF/ejUN83ulV7WkAywXgemFhM9EZTfkI7qA5xSU1tyvED7Ld8aW3DiTE
JiiNeXf1L/BXunwH1OH8zVowV36GEEfdMR/X/KLCvzB8XSSq6PmuX2p0ws5rs0bY
Ib4p1I5eFdZCSucyb6Sxa1GDWL4/bcf72gMhy2oWGU4K8K2Eyl2Us1p292E=
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            06:fd:f9:03:96:03:ad:ea:00:0a:eb:3f:27:bb:ba:1b
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root CA
        Validity
            Not Before: Nov 10 00:00:00 2006 GMT
            Not After : Nov 10 00:00:00 2021 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID CA-1
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:e8:82:2d:99:f9:ca:c2:42:95:a5:80:73:40:70:
                    d2:9e:56:54:5c:a9:c4:d2:41:06:8f:c9:33:fc:4d:
                    45:91:5c:81:9f:ed:2c:9c:f8:16:59:df:9e:b5:24:
                    15:c2:98:b9:b4:77:49:dc:89:c4:0a:da:af:cb:5e:
                    6b:ed:ad:b0:71:31:eb:cf:3a:40:0c:46:4d:93:ec:
                    8b:7a:36:08:03:ab:0c:34:fe:18:49:82:fe:c7:c7:
                    31:48:80:7c:1e:a2:0f:92:0e:50:c9:c6:87:eb:36:
                    3f:d8:30:c3:ff:a6:f7:fb:a2:cd:6f:73:23:fe:ac:
                    56:05:90:f0:32:21:16:89:c6:70:88:f9:05:97:7d:
                    a3:c7:43:dd:02:e8:3b:3d:ed:b1:41:a3:ed:3f:be:
                    db:95:48:c4:ee:1e:b3:f2:bc:0c:2b:99:d0:c6:5d:
                    12:42:81:e1:83:6e:82:73:3f:26:4b:14:90:ae:59:
                    66:0a:c4:8d:be:d2:ce:06:ae:ad:84:6f:48:84:9b:
                    4f:40:b9:f1:4c:f2:af:98:fb:f6:ce:40:5d:5c:f6:
                    a8:f1:2f:af:ec:89:22:f2:6b:18:65:b1:c1:73:ad:
                    d7:f1:d8:cf:1e:0a:74:5c:42:b8:68:7e:b7:d5:77:
                    0a:27:56:7c:0f:62:a4:3f:32:14:60:95:fd:07:04:
                    a2:09
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication, TLS Web Client Authentication, Code Signing, E-mail Protection, Time Stamping
            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.0.1.4
                  CPS: http://www.digicert.com/ssl-cps-repository.htm
                  User Notice:
                    Explicit Text: 
                Policy: 2.16.840.1.114412.3.21

            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com
                CA Issuers - URI:http://cacerts.digicert.com/DigiCertAssuredIDRootCA.crt

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertAssuredIDRootCA.crl

            X509v3 Subject Key Identifier: 
                15:00:12:2B:13:98:B2:99:07:ED:1E:DF:A2:BE:57:0D:2B:67:02:CD
            X509v3 Authority Key Identifier: 
                keyid:45:EB:A2:AF:F4:92:CB:82:31:2D:51:8B:A7:A7:21:9D:F3:6D:C8:0F

    Signature Algorithm: sha1WithRSAEncryption
         46:50:3e:c9:b7:28:24:a7:38:1d:b6:5b:29:af:52:cf:52:e9:
         31:47:ab:56:5c:7b:d5:0d:0b:41:b3:ef:ec:75:1f:74:38:f2:
         b2:5c:61:a2:9c:95:c3:50:e4:82:b9:23:d1:ba:3a:86:72:ad:
         38:78:ac:75:5d:17:17:34:72:47:85:94:56:d1:eb:bb:36:84:
         77:cc:24:a5:f3:04:19:55:a9:e7:e3:e7:ab:62:cd:fb:8b:2d:
         90:c2:c0:d2:b5:94:bd:5e:4f:b1:05:d2:0e:3d:1a:a9:14:5b:
         a6:86:31:62:a8:a8:33:e4:9b:39:a7:c4:f5:ce:1d:78:76:94:
         25:73:e4:2a:ab:cf:9c:76:4b:ed:5f:c2:4b:16:e4:4b:70:4c:
         00:89:1e:fc:c5:79:bc:4c:12:57:fe:5f:e1:1e:bc:02:5d:a8:
         fe:fb:07:38:4f:0d:c6:5d:91:b9:0f:67:45:cd:d6:83:ed:e7:
         92:0d:8d:b1:69:8c:4f:fb:59:e0:23:0f:d2:aa:ae:00:7c:ee:
         9c:42:0e:cf:91:d7:27:b7:16:ee:0f:c3:bd:7c:0a:a0:ee:2c:
         08:55:85:22:b8:eb:18:1a:4d:fc:2a:21:ad:49:31:83:47:95:
         77:71:dc:b1:1b:4b:4b:1c:10:9c:77:14:c1:9d:4f:2f:5a:95:
         08:29:10:26
-----BEGIN CERTIFICATE-----
MIIGzTCCBbWgAwIBAgIQBv35A5YDreoACus/J7u6GzANBgkqhkiG9w0BAQUFADBl
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJv
b3QgQ0EwHhcNMDYxMTEwMDAwMDAwWhcNMjExMTEwMDAwMDAwWjBiMQswCQYDVQQG
EwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNl
cnQuY29tMSEwHwYDVQQDExhEaWdpQ2VydCBBc3N1cmVkIElEIENBLTEwggEiMA0G
CSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDogi2Z+crCQpWlgHNAcNKeVlRcqcTS
QQaPyTP8TUWRXIGf7Syc+BZZ3561JBXCmLm0d0ncicQK2q/LXmvtrbBxMevPOkAM
Rk2T7It6NggDqww0/hhJgv7HxzFIgHweog+SDlDJxofrNj/YMMP/pvf7os1vcyP+
rFYFkPAyIRaJxnCI+QWXfaPHQ90C6Ds97bFBo+0/vtuVSMTuHrPyvAwrmdDGXRJC
geGDboJzPyZLFJCuWWYKxI2+0s4Grq2Eb0iEm09AufFM8q+Y+/bOQF1c9qjxL6/s
iSLyaxhlscFzrdfx2M8eCnRcQrhofrfVdwonVnwPYqQ/MhRglf0HBKIJAgMBAAGj
ggN6MIIDdjAOBgNVHQ8BAf8EBAMCAYYwOwYDVR0lBDQwMgYIKwYBBQUHAwEGCCsG
AQUFBwMCBggrBgEFBQcDAwYIKwYBBQUHAwQGCCsGAQUFBwMIMIIB0gYDVR0gBIIB
yTCCAcUwggG0BgpghkgBhv1sAAEEMIIBpDA6BggrBgEFBQcCARYuaHR0cDovL3d3
dy5kaWdpY2VydC5jb20vc3NsLWNwcy1yZXBvc2l0b3J5Lmh0bTCCAWQGCCsGAQUF
BwICMIIBVh6CAVIAQQBuAHkAIAB1AHMAZQAgAG8AZgAgAHQAaABpAHMAIABDAGUA
cgB0AGkAZgBpAGMAYQB0AGUAIABjAG8AbgBzAHQAaQB0AHUAdABlAHMAIABhAGMA
YwBlAHAAdABhAG4AYwBlACAAbwBmACAAdABoAGUAIABEAGkAZwBpAEMAZQByAHQA
IABDAFAALwBDAFAAUwAgAGEAbgBkACAAdABoAGUAIABSAGUAbAB5AGkAbgBnACAA
UABhAHIAdAB5ACAAQQBnAHIAZQBlAG0AZQBuAHQAIAB3AGgAaQBjAGgAIABsAGkA
bQBpAHQAIABsAGkAYQBiAGkAbABpAHQAeQAgAGEAbgBkACAAYQByAGUAIABpAG4A
YwBvAHIAcABvAHIAYQB0AGUAZAAgAGgAZQByAGUAaQBuACAAYgB5ACAAcgBlAGYA
ZQByAGUAbgBjAGUALjALBglghkgBhv1sAxUwEgYDVR0TAQH/BAgwBgEB/wIBADB5
BggrBgEFBQcBAQRtMGswJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRpZ2ljZXJ0
LmNvbTBDBggrBgEFBQcwAoY3aHR0cDovL2NhY2VydHMuZGlnaWNlcnQuY29tL0Rp
Z2lDZXJ0QXNzdXJlZElEUm9vdENBLmNydDCBgQYDVR0fBHoweDA6oDigNoY0aHR0
cDovL2NybDMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0QXNzdXJlZElEUm9vdENBLmNy
bDA6oDigNoY0aHR0cDovL2NybDQuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0QXNzdXJl
ZElEUm9vdENBLmNybDAdBgNVHQ4EFgQUFQASKxOYspkH7R7for5XDStnAs0wHwYD
VR0jBBgwFoAUReuir/SSy4IxLVGLp6chnfNtyA8wDQYJKoZIhvcNAQEFBQADggEB
AEZQPsm3KCSnOB22WymvUs9S6TFHq1Zce9UNC0Gz7+x1H3Q48rJcYaKclcNQ5IK5
I9G6OoZyrTh4rHVdFxc0ckeFlFbR67s2hHfMJKXzBBlVqefj56tizfuLLZDCwNK1
lL1eT7EF0g49GqkUW6aGMWKoqDPkmzmnxPXOHXh2lCVz5Cqrz5x2S+1fwksW5Etw
TACJHvzFebxMElf+X+EevAJdqP77BzhPDcZdkbkPZ0XN1oPt55INjbFpjE/7WeAj
D9KqrgB87pxCDs+R1ye3Fu4Pw718CqDuLAhVhSK46xgaTfwqIa1JMYNHlXdx3LEb
S0scEJx3FMGdTy9alQgpECY=
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            01:65:cf:cc:19:17:4b:69:87:26:41:59:ca:e0:c7:9e
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global Root CA
        Validity
            Not Before: Nov 10 00:00:00 2006 GMT
            Not After : Nov 10 00:00:00 2021 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Global CA-1
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:f2:25:1c:31:67:b7:53:eb:a8:4f:a5:fd:e6:cf:
                    30:52:af:10:35:02:0c:6c:c4:ff:f4:3c:a2:8a:19:
                    73:00:14:da:2b:6a:c1:68:30:4b:ff:1b:36:d4:c8:
                    dd:26:6e:2c:90:16:71:6a:3e:8f:67:b0:92:96:ac:
                    8b:de:28:97:a3:fb:b0:a5:2c:e2:cc:a7:8b:cb:b9:
                    18:51:aa:dd:06:7f:9b:8f:f6:6c:8c:be:04:0c:77:
                    9b:59:50:7b:50:0e:05:5c:25:9c:95:7e:39:87:09:
                    0f:43:4e:0e:50:b7:6f:27:6e:c4:99:0d:f4:33:40:
                    0f:77:ce:e6:4a:54:9e:7d:8c:3c:d3:c1:c0:4b:b4:
                    bb:21:4e:fa:7a:0a:ef:9d:ed:84:17:8a:28:a5:de:
                    38:2e:8b:d4:62:8e:2f:65:af:ec:85:ca:6c:a9:0b:
                    c0:5e:41:eb:a6:b9:ba:79:31:58:09:53:1c:f9:dd:
                    a1:84:53:31:db:19:e6:11:9e:07:74:b1:07:8e:1c:
                    9e:45:8d:77:9c:4a:58:f3:89:8d:3f:62:60:fa:63:
                    a2:e6:ce:e0:3a:ef:e5:3f:95:49:7f:9b:61:fd:76:
                    df:c0:7d:50:d7:d8:5b:53:e3:0f:ee:56:15:87:7c:
                    24:62:3b:69:c4:02:4d:78:b6:19:1e:ee:e0:e9:83:
                    e4:19
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication, TLS Web Client Authentication, Code Signing, E-mail Protection, Time Stamping
            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.1.3.0.3
                  CPS: http://www.digicert.com/ssl-cps-repository.htm
                  User Notice:
                    Explicit Text: 

            X509v3 Basic Constraints: critical
                CA:TRUE
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com
                CA Issuers - URI:http://www.digicert.com/CACerts/DigiCertGlobalRootCA.crt

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertGlobalRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertGlobalRootCA.crl

            X509v3 Subject Key Identifier: 
                1E:1C:88:15:AA:F2:46:D0:05:DA:E9:1E:DC:22:BD:A8:97:DE:0F:B2
            X509v3 Authority Key Identifier: 
                keyid:03:DE:50:35:56:D1:4C:BB:66:F0:A3:E2:1B:1B:C3:97:B2:3D:D1:55

    Signature Algorithm: sha1WithRSAEncryption
         64:d7:0f:15:aa:44:48:2a:5b:11:5e:6d:d6:43:5f:e0:6c:8c:
         af:14:d7:9d:5b:fe:4f:0f:eb:a2:94:21:06:83:73:bd:7f:05:
         78:d1:90:6e:33:0c:c8:ec:77:fa:dd:08:35:4c:12:6b:fd:12:
         05:66:e5:c5:0e:64:30:4e:ce:b9:03:7e:9e:c1:db:43:a7:60:
         fe:7e:3a:53:6f:cd:e5:97:e7:d4:05:94:b1:3b:5f:35:0a:c7:
         3b:03:44:fa:dd:ce:b5:83:f9:4f:ba:bf:b2:67:6c:57:37:d7:
         3d:ab:18:51:c6:6a:14:50:3e:95:0c:35:12:75:35:42:95:63:
         76:a6:f6:f6:39:2b:78:1e:21:60:11:37:a2:73:72:c0:4a:78:
         6f:96:44:12:62:6c:38:f7:ea:d3:df:74:11:1b:2e:55:a3:c7:
         bb:8e:7c:2c:58:a4:3c:7a:2f:2c:c3:19:03:ce:ef:bc:56:6b:
         81:af:4a:00:ab:10:53:43:4e:17:fc:ae:78:60:2e:67:00:b0:
         06:37:b7:1a:fc:85:d9:56:2c:5f:6d:e5:b0:7d:46:c2:da:26:
         2a:0a:45:7b:b3:4b:32:ca:87:19:0a:2b:6f:c5:3a:a9:a2:d0:
         0c:21:31:13:14:95:b6:fa:3c:e7:e0:62:61:6e:6d:b7:bb:3c:
         54:b5:96:88
-----BEGIN CERTIFICATE-----
MIIGsDCCBZigAwIBAgIQAWXPzBkXS2mHJkFZyuDHnjANBgkqhkiG9w0BAQUFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0wNjExMTAwMDAwMDBaFw0yMTExMTAwMDAwMDBaMF4xCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j
b20xHTAbBgNVBAMTFERpZ2lDZXJ0IEdsb2JhbCBDQS0xMIIBIjANBgkqhkiG9w0B
AQEFAAOCAQ8AMIIBCgKCAQEA8iUcMWe3U+uoT6X95s8wUq8QNQIMbMT/9Dyiihlz
ABTaK2rBaDBL/xs21MjdJm4skBZxaj6PZ7CSlqyL3iiXo/uwpSzizKeLy7kYUard
Bn+bj/ZsjL4EDHebWVB7UA4FXCWclX45hwkPQ04OULdvJ27EmQ30M0APd87mSlSe
fYw808HAS7S7IU76egrvne2EF4oopd44LovUYo4vZa/shcpsqQvAXkHrprm6eTFY
CVMc+d2hhFMx2xnmEZ4HdLEHjhyeRY13nEpY84mNP2Jg+mOi5s7gOu/lP5VJf5th
/XbfwH1Q19hbU+MP7lYVh3wkYjtpxAJNeLYZHu7g6YPkGQIDAQABo4IDZTCCA2Ew
DgYDVR0PAQH/BAQDAgGGMDsGA1UdJQQ0MDIGCCsGAQUFBwMBBggrBgEFBQcDAgYI
KwYBBQUHAwMGCCsGAQUFBwMEBggrBgEFBQcDCDCCAcYGA1UdIASCAb0wggG5MIIB
tQYLYIZIAYb9bAEDAAMwggGkMDoGCCsGAQUFBwIBFi5odHRwOi8vd3d3LmRpZ2lj
ZXJ0LmNvbS9zc2wtY3BzLXJlcG9zaXRvcnkuaHRtMIIBZAYIKwYBBQUHAgIwggFW
HoIBUgBBAG4AeQAgAHUAcwBlACAAbwBmACAAdABoAGkAcwAgAEMAZQByAHQAaQBm
AGkAYwBhAHQAZQAgAGMAbwBuAHMAdABpAHQAdQB0AGUAcwAgAGEAYwBjAGUAcAB0
AGEAbgBjAGUAIABvAGYAIAB0AGgAZQAgAEQAaQBnAGkAQwBlAHIAdAAgAEMAUAAv
AEMAUABTACAAYQBuAGQAIAB0AGgAZQAgAFIAZQBsAHkAaQBuAGcAIABQAGEAcgB0
AHkAIABBAGcAcgBlAGUAbQBlAG4AdAAgAHcAaABpAGMAaAAgAGwAaQBtAGkAdAAg
AGwAaQBhAGIAaQBsAGkAdAB5ACAAYQBuAGQAIABhAHIAZQAgAGkAbgBjAG8AcgBw
AG8AcgBhAHQAZQBkACAAaABlAHIAZQBpAG4AIABiAHkAIAByAGUAZgBlAHIAZQBu
AGMAZQAuMA8GA1UdEwEB/wQFMAMBAf8wegYIKwYBBQUHAQEEbjBsMCQGCCsGAQUF
BzABhhhodHRwOi8vb2NzcC5kaWdpY2VydC5jb20wRAYIKwYBBQUHMAKGOGh0dHA6
Ly93d3cuZGlnaWNlcnQuY29tL0NBQ2VydHMvRGlnaUNlcnRHbG9iYWxSb290Q0Eu
Y3J0MHsGA1UdHwR0MHIwN6A1oDOGMWh0dHA6Ly9jcmwzLmRpZ2ljZXJ0LmNvbS9E
aWdpQ2VydEdsb2JhbFJvb3RDQS5jcmwwN6A1oDOGMWh0dHA6Ly9jcmw0LmRpZ2lj
ZXJ0LmNvbS9EaWdpQ2VydEdsb2JhbFJvb3RDQS5jcmwwHQYDVR0OBBYEFB4ciBWq
8kbQBdrpHtwivaiX3g+yMB8GA1UdIwQYMBaAFAPeUDVW0Uy7ZvCj4hsbw5eyPdFV
MA0GCSqGSIb3DQEBBQUAA4IBAQBk1w8VqkRIKlsRXm3WQ1/gbIyvFNedW/5PD+ui
lCEGg3O9fwV40ZBuMwzI7Hf63Qg1TBJr/RIFZuXFDmQwTs65A36ewdtDp2D+fjpT
b83ll+fUBZSxO181Csc7A0T63c61g/lPur+yZ2xXN9c9qxhRxmoUUD6VDDUSdTVC
lWN2pvb2OSt4HiFgETeic3LASnhvlkQSYmw49+rT33QRGy5Vo8e7jnwsWKQ8ei8s
wxkDzu+8VmuBr0oAqxBTQ04X/K54YC5nALAGN7ca/IXZVixfbeWwfUbC2iYqCkV7
s0syyocZCitvxTqpotAMITETFJW2+jzn4GJhbm23uzxUtZaI
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0f:a8:49:06:15:d7:00:a0:be:21:76:fd:c5:ec:6d:bd
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root CA
        Validity
            Not Before: Feb 11 12:00:00 2011 GMT
            Not After : Feb 10 12:00:00 2026 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Code Signing CA-1
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:9c:7c:f9:a0:8f:0a:ca:89:4b:53:9a:3c:ec:19:
                    22:7f:0c:cb:f7:48:44:d0:3f:22:6e:9a:4f:fa:ce:
                    df:c6:d3:24:91:ff:a8:52:93:e7:72:f8:f1:46:86:
                    94:c5:ab:17:f4:78:7f:cb:7a:be:90:26:1c:7c:53:
                    2f:e5:83:e7:bb:6a:05:28:cc:a0:17:11:4b:18:20:
                    f3:b7:bd:f2:dc:cb:b8:86:40:45:1f:4f:f9:39:98:
                    d3:3b:eb:6f:69:a4:5c:01:2e:b1:66:e9:a6:b8:dc:
                    06:de:ae:8d:62:fe:44:27:82:a9:03:c3:fe:e6:26:
                    0b:93:49:71:e5:38:6d:4e:fe:80:5c:67:77:b5:d5:
                    6f:66:64:94:6b:bb:88:49:68:d8:d1:f4:7e:21:c1:
                    e6:c5:1e:c9:e0:96:b8:09:95:f9:4b:9e:ae:cc:31:
                    bd:6a:5a:42:a7:c4:05:47:7d:60:2b:d9:a1:5d:bd:
                    fb:56:93:ef:1b:fc:72:f9:df:3e:32:fd:dd:2a:2c:
                    32:46:da:5e:63:bc:d0:b8:cb:4c:10:35:fc:50:5f:
                    ac:43:44:7c:fe:84:44:50:b4:8c:2a:3d:74:0c:f4:
                    7a:8b:72:4b:e0:56:b5:13:b6:63:e9:3f:4e:49:23:
                    e7:06:16:f0:25:ab:01:f8:28:41:e6:c7:d4:64:3b:
                    5d:ab
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                Code Signing
            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.3
                  CPS: http://www.digicert.com/ssl-cps-repository.htm
                  User Notice:
                    Explicit Text: 

            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com
                CA Issuers - URI:http://cacerts.digicert.com/DigiCertAssuredIDRootCA.crt

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertAssuredIDRootCA.crl

            X509v3 Subject Key Identifier: 
                7B:68:CE:29:AA:C0:17:BE:49:7A:E1:E5:3F:D6:A7:F7:45:8F:35:32
            X509v3 Authority Key Identifier: 
                keyid:45:EB:A2:AF:F4:92:CB:82:31:2D:51:8B:A7:A7:21:9D:F3:6D:C8:0F

    Signature Algorithm: sha1WithRSAEncryption
         7b:72:1d:64:ff:88:c8:3a:c1:b7:e9:e7:a9:c4:87:bb:db:94:
         92:d7:90:59:33:fa:2b:87:de:a8:5b:80:25:3f:13:8f:9b:83:
         1b:7c:43:c4:e6:8c:df:39:3e:c3:15:ec:b0:da:3b:21:25:7b:
         24:c1:72:5d:b8:47:91:81:13:46:fa:9c:3f:6a:51:38:de:b4:
         25:cb:f0:ab:df:c5:28:01:54:79:10:46:24:d1:38:0f:26:a1:
         61:90:4d:ba:bd:28:e6:3f:f1:c4:aa:9b:f6:da:35:53:4f:c9:
         f2:3d:d3:6c:dc:23:ed:aa:a0:4d:67:09:f3:3a:80:3d:3c:fb:
         36:4c:90:e7:76:a4:dd:f2:3a:bf:56:35:2f:a2:4c:65:e8:e0:
         d4:da:d1:c7:c8:91:6a:2d:23:4f:37:3b:19:94:18:d4:d5:9c:
         10:3c:d5:b1:1c:19:ff:8f:c8:6b:9b:9e:f8:ae:9c:99:96:78:
         d1:cd:9c:51:15:5b:42:26:72:5a:8d:0a:4a:23:92:40:e8:86:
         de:22:c2:93:3a:d4:9b:68:a6:df:29:7f:06:b9:3c:0e:bd:9f:
         c4:86:9c:82:47:42:71:32:86:09:99:72:09:79:4b:9d:71:69:
         f5:41:ff:7f:39:77:64:f1:84:8d:be:8b:1e:b2:7d:68:a3:a5:
         90:b1:0c:ff
-----BEGIN CERTIFICATE-----
MIIGozCCBYugAwIBAgIQD6hJBhXXAKC+IXb9xextvTANBgkqhkiG9w0BAQUFADBl
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJv
b3QgQ0EwHhcNMTEwMjExMTIwMDAwWhcNMjYwMjEwMTIwMDAwWjBvMQswCQYDVQQG
EwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNl
cnQuY29tMS4wLAYDVQQDEyVEaWdpQ2VydCBBc3N1cmVkIElEIENvZGUgU2lnbmlu
ZyBDQS0xMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnHz5oI8KyolL
U5o87BkifwzL90hE0D8ibppP+s7fxtMkkf+oUpPncvjxRoaUxasX9Hh/y3q+kCYc
fFMv5YPnu2oFKMygFxFLGCDzt73y3Mu4hkBFH0/5OZjTO+tvaaRcAS6xZummuNwG
3q6NYv5EJ4KpA8P+5iYLk0lx5ThtTv6AXGd3tdVvZmSUa7uISWjY0fR+IcHmxR7J
4Ja4CZX5S56uzDG9alpCp8QFR31gK9mhXb37VpPvG/xy+d8+Mv3dKiwyRtpeY7zQ
uMtMEDX8UF+sQ0R8/oREULSMKj10DPR6i3JL4Fa1E7Zj6T9OSSPnBhbwJasB+ChB
5sfUZDtdqwIDAQABo4IDQzCCAz8wDgYDVR0PAQH/BAQDAgGGMBMGA1UdJQQMMAoG
CCsGAQUFBwMDMIIBwwYDVR0gBIIBujCCAbYwggGyBghghkgBhv1sAzCCAaQwOgYI
KwYBBQUHAgEWLmh0dHA6Ly93d3cuZGlnaWNlcnQuY29tL3NzbC1jcHMtcmVwb3Np
dG9yeS5odG0wggFkBggrBgEFBQcCAjCCAVYeggFSAEEAbgB5ACAAdQBzAGUAIABv
AGYAIAB0AGgAaQBzACAAQwBlAHIAdABpAGYAaQBjAGEAdABlACAAYwBvAG4AcwB0
AGkAdAB1AHQAZQBzACAAYQBjAGMAZQBwAHQAYQBuAGMAZQAgAG8AZgAgAHQAaABl
ACAARABpAGcAaQBDAGUAcgB0ACAAQwBQAC8AQwBQAFMAIABhAG4AZAAgAHQAaABl
ACAAUgBlAGwAeQBpAG4AZwAgAFAAYQByAHQAeQAgAEEAZwByAGUAZQBtAGUAbgB0
ACAAdwBoAGkAYwBoACAAbABpAG0AaQB0ACAAbABpAGEAYgBpAGwAaQB0AHkAIABh
AG4AZAAgAGEAcgBlACAAaQBuAGMAbwByAHAAbwByAGEAdABlAGQAIABoAGUAcgBl
AGkAbgAgAGIAeQAgAHIAZQBmAGUAcgBlAG4AYwBlAC4wEgYDVR0TAQH/BAgwBgEB
/wIBADB5BggrBgEFBQcBAQRtMGswJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRp
Z2ljZXJ0LmNvbTBDBggrBgEFBQcwAoY3aHR0cDovL2NhY2VydHMuZGlnaWNlcnQu
Y29tL0RpZ2lDZXJ0QXNzdXJlZElEUm9vdENBLmNydDCBgQYDVR0fBHoweDA6oDig
NoY0aHR0cDovL2NybDMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0QXNzdXJlZElEUm9v
dENBLmNybDA6oDigNoY0aHR0cDovL2NybDQuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0
QXNzdXJlZElEUm9vdENBLmNybDAdBgNVHQ4EFgQUe2jOKarAF75JeuHlP9an90WP
NTIwHwYDVR0jBBgwFoAUReuir/SSy4IxLVGLp6chnfNtyA8wDQYJKoZIhvcNAQEF
BQADggEBAHtyHWT/iMg6wbfp56nEh7vblJLXkFkz+iuH3qhbgCU/E4+bgxt8Q8Tm
jN85PsMV7LDaOyEleyTBcl24R5GBE0b6nD9qUTjetCXL8KvfxSgBVHkQRiTROA8m
oWGQTbq9KOY/8cSqm/baNVNPyfI902zcI+2qoE1nCfM6gD08+zZMkOd2pN3yOr9W
NS+iTGXo4NTa0cfIkWotI083OxmUGNTVnBA81bEcGf+PyGubnviunJmWeNHNnFEV
W0ImclqNCkojkkDoht4iwpM61Jtopt8pfwa5PA69n8SGnIJHQnEyhgmZcgl5S51x
afVB/385d2TxhI2+ix6yfWijpZCxDP8=
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            02:c4:d1:e5:8a:4a:68:0c:56:8d:a3:04:7e:7e:4d:5f
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Validity
            Not Before: Feb 11 12:00:00 2011 GMT
            Not After : Feb 10 12:00:00 2026 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance Code Signing CA-1
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:c5:f9:23:e6:94:27:c4:80:14:a4:80:32:5f:40:
                    a3:8d:6f:70:c0:e5:36:71:71:3a:75:a4:aa:1a:92:
                    94:89:5e:ac:23:71:cb:4e:67:7d:41:3f:aa:e3:4b:
                    b7:7b:be:9d:c1:a8:38:8f:69:2f:3a:24:e9:77:59:
                    12:c7:66:04:43:c2:0d:26:82:89:40:19:f2:2c:ea:
                    e7:4c:e7:7c:05:1a:b8:ff:88:09:4f:26:37:ef:3a:
                    a4:fa:22:6c:88:c9:4a:1b:61:f2:ae:10:5e:6f:bc:
                    d1:79:9b:59:18:60:e5:ee:29:b5:03:2a:a4:ce:f1:
                    83:19:4f:69:05:73:28:09:fb:22:10:93:22:a0:90:
                    19:1a:4c:31:f2:d3:2b:d8:84:43:af:3c:63:ff:98:
                    db:20:d2:09:2b:54:c1:ea:fd:6a:83:e7:10:a3:12:
                    71:f5:d6:d7:e1:12:7a:d5:e0:56:5a:ce:ea:01:5b:
                    68:65:5b:c1:3f:58:52:33:a9:35:61:4e:22:cb:81:
                    ca:36:a3:12:cb:06:d6:cf:1b:4d:18:7e:b9:92:b9:
                    12:cf:40:26:d8:9a:36:85:b3:15:aa:47:93:84:6b:
                    07:bb:bc:d5:b3:de:25:00:11:89:00:68:c1:29:3c:
                    ea:3e:2d:ee:50:ab:d7:1c:30:06:78:3c:a5:10:23:
                    67:91
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                Code Signing
            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.3
                  CPS: http://www.digicert.com/ssl-cps-repository.htm
                  User Notice:
                    Explicit Text: 

            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com
                CA Issuers - URI:http://cacerts.digicert.com/DigiCertHighAssuranceEVRootCA.crt

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertHighAssuranceEVRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertHighAssuranceEVRootCA.crl

            X509v3 Subject Key Identifier: 
                97:48:03:EB:15:08:6B:B9:B2:58:23:CC:94:2E:F1:C6:65:D2:64:8E
            X509v3 Authority Key Identifier: 
                keyid:B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3

    Signature Algorithm: sha1WithRSAEncryption
         49:eb:7c:60:be:ae:ef:c9:7c:b3:c5:ba:4b:64:df:16:69:e2:
         86:fa:29:d9:de:98:85:7d:40:66:26:33:2f:44:55:aa:aa:90:
         e9:35:70:0a:34:be:d3:ae:54:2e:8e:65:00:d6:7a:32:20:3e:
         6c:26:b8:98:a9:39:b1:bc:95:c7:aa:e9:f5:ee:46:66:c6:b3:
         e8:12:f8:b3:97:9d:ff:74:58:82:34:99:75:50:ac:44:8f:e8:
         92:ce:7d:8b:0f:31:96:c7:dc:d3:11:30:98:74:16:c6:e5:6b:
         45:76:a3:94:01:cd:33:00:7a:48:f6:6f:86:31:c9:56:2b:33:
         22:d5:f8:01:b6:44:ce:8c:b4:ca:88:d2:e4:16:e3:e7:f6:e2:
         3e:e1:09:c0:9d:79:43:43:7f:55:5c:05:ad:93:10:c6:2c:0d:
         6b:c0:9e:ea:78:e5:d2:77:d6:b8:da:9a:98:7f:ba:4c:92:2b:
         9d:bd:a4:88:b1:dd:af:c3:4c:d2:97:9b:03:c6:ae:5f:1b:44:
         0f:33:37:15:e3:cb:ff:2f:56:d3:16:a4:5b:55:67:9d:a2:ca:
         db:34:6c:0c:73:4a:b5:7b:a4:b6:b3:e9:35:02:78:70:ec:00:
         7a:cb:fc:4b:4f:22:36:bb:14:84:c9:8f:91:dd:0f:3c:75:8c:
         ca:0b:88:e7
-----BEGIN CERTIFICATE-----
MIIGwjCCBaqgAwIBAgIQAsTR5YpKaAxWjaMEfn5NXzANBgkqhkiG9w0BAQUFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTExMDIxMTEyMDAwMFoXDTI2MDIxMDEyMDAwMFowczEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTEyMDAGA1UEAxMpRGlnaUNlcnQgSGlnaCBBc3N1cmFuY2Ug
Q29kZSBTaWduaW5nIENBLTEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQDF+SPmlCfEgBSkgDJfQKONb3DA5TZxcTp1pKoakpSJXqwjcctOZ31BP6rjS7d7
vp3BqDiPaS86JOl3WRLHZgRDwg0mgolAGfIs6udM53wFGrj/iAlPJjfvOqT6ImyI
yUobYfKuEF5vvNF5m1kYYOXuKbUDKqTO8YMZT2kFcygJ+yIQkyKgkBkaTDHy0yvY
hEOvPGP/mNsg0gkrVMHq/WqD5xCjEnH11tfhEnrV4FZazuoBW2hlW8E/WFIzqTVh
TiLLgco2oxLLBtbPG00YfrmSuRLPQCbYmjaFsxWqR5OEawe7vNWz3iUAEYkAaMEp
POo+Le5Qq9ccMAZ4PKUQI2eRAgMBAAGjggNXMIIDUzAOBgNVHQ8BAf8EBAMCAYYw
EwYDVR0lBAwwCgYIKwYBBQUHAwMwggHDBgNVHSAEggG6MIIBtjCCAbIGCGCGSAGG
/WwDMIIBpDA6BggrBgEFBQcCARYuaHR0cDovL3d3dy5kaWdpY2VydC5jb20vc3Ns
LWNwcy1yZXBvc2l0b3J5Lmh0bTCCAWQGCCsGAQUFBwICMIIBVh6CAVIAQQBuAHkA
IAB1AHMAZQAgAG8AZgAgAHQAaABpAHMAIABDAGUAcgB0AGkAZgBpAGMAYQB0AGUA
IABjAG8AbgBzAHQAaQB0AHUAdABlAHMAIABhAGMAYwBlAHAAdABhAG4AYwBlACAA
bwBmACAAdABoAGUAIABEAGkAZwBpAEMAZQByAHQAIABDAFAALwBDAFAAUwAgAGEA
bgBkACAAdABoAGUAIABSAGUAbAB5AGkAbgBnACAAUABhAHIAdAB5ACAAQQBnAHIA
ZQBlAG0AZQBuAHQAIAB3AGgAaQBjAGgAIABsAGkAbQBpAHQAIABsAGkAYQBiAGkA
bABpAHQAeQAgAGEAbgBkACAAYQByAGUAIABpAG4AYwBvAHIAcABvAHIAYQB0AGUA
ZAAgAGgAZQByAGUAaQBuACAAYgB5ACAAcgBlAGYAZQByAGUAbgBjAGUALjASBgNV
HRMBAf8ECDAGAQH/AgEAMH8GCCsGAQUFBwEBBHMwcTAkBggrBgEFBQcwAYYYaHR0
cDovL29jc3AuZGlnaWNlcnQuY29tMEkGCCsGAQUFBzAChj1odHRwOi8vY2FjZXJ0
cy5kaWdpY2VydC5jb20vRGlnaUNlcnRIaWdoQXNzdXJhbmNlRVZSb290Q0EuY3J0
MIGPBgNVHR8EgYcwgYQwQKA+oDyGOmh0dHA6Ly9jcmwzLmRpZ2ljZXJ0LmNvbS9E
aWdpQ2VydEhpZ2hBc3N1cmFuY2VFVlJvb3RDQS5jcmwwQKA+oDyGOmh0dHA6Ly9j
cmw0LmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEhpZ2hBc3N1cmFuY2VFVlJvb3RDQS5j
cmwwHQYDVR0OBBYEFJdIA+sVCGu5slgjzJQu8cZl0mSOMB8GA1UdIwQYMBaAFLE+
w2kD+L9HAdSYJhoIAu9jZCvDMA0GCSqGSIb3DQEBBQUAA4IBAQBJ63xgvq7vyXyz
xbpLZN8WaeKG+inZ3piFfUBmJjMvRFWqqpDpNXAKNL7TrlQujmUA1noyID5sJriY
qTmxvJXHqun17kZmxrPoEvizl53/dFiCNJl1UKxEj+iSzn2LDzGWx9zTETCYdBbG
5WtFdqOUAc0zAHpI9m+GMclWKzMi1fgBtkTOjLTKiNLkFuPn9uI+4QnAnXlDQ39V
XAWtkxDGLA1rwJ7qeOXSd9a42pqYf7pMkiudvaSIsd2vw0zSl5sDxq5fG0QPMzcV
48v/L1bTFqRbVWedosrbNGwMc0q1e6S2s+k1Anhw7AB6y/xLTyI2uxSEyY+R3Q88
dYzKC4jn
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0d:d0:e3:37:4a:c9:5b:db:fa:6b:43:4b:2a:48:ec:06
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Validity
            Not Before: Apr 18 12:00:00 2012 GMT
            Not After : Apr 18 12:00:00 2027 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert EV Code Signing CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:b9:06:74:1c:5d:b4:20:aa:a9:21:a8:2a:42:46:
                    ab:25:20:17:25:cb:22:8f:90:a2:a0:31:6b:83:05:
                    75:af:b2:0e:7c:12:49:7b:6a:86:64:84:0f:83:dc:
                    64:b9:b1:6e:16:05:3e:1c:95:b9:e7:e7:88:6d:b8:
                    62:81:90:79:d4:dd:f5:e2:96:f9:c3:b5:88:23:57:
                    4a:1a:cf:71:29:e9:08:00:8f:b5:98:e3:a7:32:fd:
                    ac:2e:b8:f4:93:53:f4:0a:39:43:91:af:d5:6b:e8:
                    d4:9f:46:bd:8e:3d:ab:e2:f9:2b:d4:ea:00:40:66:
                    24:b7:e8:7f:b4:44:75:8d:78:9a:ae:31:c1:37:cf:
                    4e:1f:5b:f8:45:4a:d7:3f:c2:c9:92:06:64:be:de:
                    06:8a:af:d0:e8:8a:b1:f0:2c:88:00:6f:0b:dc:85:
                    a7:4c:cb:06:bf:d6:2e:2a:32:6e:29:71:af:8e:22:
                    f3:0f:d0:d8:98:48:2d:a8:08:cb:b6:8b:23:c2:63:
                    e0:b6:73:eb:6f:7d:26:4f:8b:f7:34:3d:37:86:0c:
                    b7:78:27:f4:c2:86:db:43:6b:5a:f8:3d:3d:f4:e8:
                    b0:62:56:c6:e7:ed:78:a1:fb:fd:7a:72:4f:32:65:
                    c4:7c:c3:c4:77:a0:04:32:32:ed:8f:3f:af:86:dd:
                    7e:d1
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                Code Signing
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com
                CA Issuers - URI:http://cacerts.digicert.com/DigiCertHighAssuranceEVRootCA.crt

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertHighAssuranceEVRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertHighAssuranceEVRootCA.crl

            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.3.2
                  CPS: http://www.digicert.com/ssl-cps-repository.htm
                  User Notice:
                    Explicit Text: 

            X509v3 Subject Key Identifier: 
                AD:69:06:70:FC:80:1B:16:B3:A9:18:94:6B:94:02:86:5E:F7:27:8C
            X509v3 Authority Key Identifier: 
                keyid:B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3

    Signature Algorithm: sha1WithRSAEncryption
         9e:5b:96:3a:2e:12:88:ac:ab:01:6d:a4:9f:75:e4:01:87:a3:
         a5:32:d7:bc:ba:a9:7e:a3:d6:14:17:f7:c2:13:6b:7c:73:8f:
         2b:6a:e5:0f:26:59:68:b0:8e:25:9b:6c:ef:fa:6c:93:92:08:
         c1:4d:cf:45:9e:9c:46:d6:1e:74:a1:9b:14:a3:fa:01:2f:4a:
         b1:01:e1:72:40:48:11:13:68:b9:36:9d:91:4b:d7:c2:39:12:
         10:c1:c4:dc:bb:62:14:14:2a:61:5d:4f:38:7c:66:1f:c6:1b:
         ff:ad:be:4f:7f:94:5b:73:43:00:0f:4d:73:b7:51:cf:0e:f6:
         77:c0:5b:cd:34:8c:d9:63:13:aa:0e:61:11:d6:f2:8e:27:fc:
         b4:7b:b8:b9:11:20:91:86:78:ea:0e:d4:28:ff:2a:d5:24:38:
         e8:37:b2:ec:96:bb:9f:bc:4a:16:50:e1:5e:bf:51:7d:23:a0:
         32:c7:c1:94:9e:7a:c9:c0:26:a2:cc:25:87:a0:12:7e:74:9f:
         2d:8d:b1:c8:e7:84:be:b9:d1:e9:de:bb:6a:4e:88:73:71:e1:
         22:38:cb:24:87:e9:73:7e:51:b2:ff:98:eb:4e:7e:2f:e0:ca:
         0e:fa:b3:5e:d1:ba:05:42:a8:48:9f:83:f6:3f:c4:ca:a8:df:
         68:a0:50:61
-----BEGIN CERTIFICATE-----
MIIGtTCCBZ2gAwIBAgIQDdDjN0rJW9v6a0NLKkjsBjANBgkqhkiG9w0BAQUFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTEyMDQxODEyMDAwMFoXDTI3MDQxODEyMDAwMFowZTEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTEkMCIGA1UEAxMbRGlnaUNlcnQgRVYgQ29kZSBTaWduaW5n
IENBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuQZ0HF20IKqpIagq
QkarJSAXJcsij5CioDFrgwV1r7IOfBJJe2qGZIQPg9xkubFuFgU+HJW55+eIbbhi
gZB51N314pb5w7WII1dKGs9xKekIAI+1mOOnMv2sLrj0k1P0CjlDka/Va+jUn0a9
jj2r4vkr1OoAQGYkt+h/tER1jXiarjHBN89OH1v4RUrXP8LJkgZkvt4Giq/Q6Iqx
8CyIAG8L3IWnTMsGv9YuKjJuKXGvjiLzD9DYmEgtqAjLtosjwmPgtnPrb30mT4v3
ND03hgy3eCf0wobbQ2ta+D099OiwYlbG5+14ofv9enJPMmXEfMPEd6AEMjLtjz+v
ht1+0QIDAQABo4IDWDCCA1QwEgYDVR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8E
BAMCAYYwEwYDVR0lBAwwCgYIKwYBBQUHAwMwfwYIKwYBBQUHAQEEczBxMCQGCCsG
AQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2VydC5jb20wSQYIKwYBBQUHMAKGPWh0
dHA6Ly9jYWNlcnRzLmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEhpZ2hBc3N1cmFuY2VF
VlJvb3RDQS5jcnQwgY8GA1UdHwSBhzCBhDBAoD6gPIY6aHR0cDovL2NybDMuZGln
aWNlcnQuY29tL0RpZ2lDZXJ0SGlnaEFzc3VyYW5jZUVWUm9vdENBLmNybDBAoD6g
PIY6aHR0cDovL2NybDQuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0SGlnaEFzc3VyYW5j
ZUVWUm9vdENBLmNybDCCAcQGA1UdIASCAbswggG3MIIBswYJYIZIAYb9bAMCMIIB
pDA6BggrBgEFBQcCARYuaHR0cDovL3d3dy5kaWdpY2VydC5jb20vc3NsLWNwcy1y
ZXBvc2l0b3J5Lmh0bTCCAWQGCCsGAQUFBwICMIIBVh6CAVIAQQBuAHkAIAB1AHMA
ZQAgAG8AZgAgAHQAaABpAHMAIABDAGUAcgB0AGkAZgBpAGMAYQB0AGUAIABjAG8A
bgBzAHQAaQB0AHUAdABlAHMAIABhAGMAYwBlAHAAdABhAG4AYwBlACAAbwBmACAA
dABoAGUAIABEAGkAZwBpAEMAZQByAHQAIABDAFAALwBDAFAAUwAgAGEAbgBkACAA
dABoAGUAIABSAGUAbAB5AGkAbgBnACAAUABhAHIAdAB5ACAAQQBnAHIAZQBlAG0A
ZQBuAHQAIAB3AGgAaQBjAGgAIABsAGkAbQBpAHQAIABsAGkAYQBiAGkAbABpAHQA
eQAgAGEAbgBkACAAYQByAGUAIABpAG4AYwBvAHIAcABvAHIAYQB0AGUAZAAgAGgA
ZQByAGUAaQBuACAAYgB5ACAAcgBlAGYAZQByAGUAbgBjAGUALjAdBgNVHQ4EFgQU
rWkGcPyAGxazqRiUa5QChl73J4wwHwYDVR0jBBgwFoAUsT7DaQP4v0cB1JgmGggC
72NkK8MwDQYJKoZIhvcNAQEFBQADggEBAJ5bljouEoisqwFtpJ915AGHo6Uy17y6
qX6j1hQX98ITa3xzjytq5Q8mWWiwjiWbbO/6bJOSCMFNz0WenEbWHnShmxSj+gEv
SrEB4XJASBETaLk2nZFL18I5EhDBxNy7YhQUKmFdTzh8Zh/GG/+tvk9/lFtzQwAP
TXO3Uc8O9nfAW800jNljE6oOYRHW8o4n/LR7uLkRIJGGeOoO1Cj/KtUkOOg3suyW
u5+8ShZQ4V6/UX0joDLHwZSeesnAJqLMJYegEn50ny2NscjnhL650eneu2pOiHNx
4SI4yySH6XN+UbL/mOtOfi/gyg76s17RugVCqEifg/Y/xMqo32igUGE=
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            03:f1:b4:e1:5f:3a:82:f1:14:96:78:b3:d7:d8:47:5c
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Validity
            Not Before: Apr 18 12:00:00 2012 GMT
            Not After : Apr 18 12:00:00 2027 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert EV Code Signing CA (SHA2)
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:a7:53:fa:0f:b2:b5:13:f1:64:cf:84:80:fc:ae:
                    80:35:d1:b6:d7:c7:a3:2c:ac:1a:2c:ac:f1:84:ac:
                    3a:35:12:3a:92:91:ba:57:e4:c4:c9:f3:2f:a8:48:
                    3c:b7:d6:6e:dc:97:22:ba:51:79:61:af:43:2f:0d:
                    b7:9b:b4:49:31:ae:44:58:3e:a4:a1:96:a7:87:4f:
                    23:7e:c3:6c:65:24:90:55:3e:a1:ca:23:7c:c5:42:
                    e9:c4:7a:62:45:9b:7d:de:63:74:cb:9e:63:25:f8:
                    84:9a:9a:ad:45:4f:ae:7d:1f:c8:13:cb:75:9b:c9:
                    e1:e1:8a:f8:0b:0c:98:f4:ca:3e:d0:45:aa:7a:1e:
                    a5:58:93:36:34:be:2b:2e:2b:31:58:66:b4:32:10:
                    9f:9d:f0:52:a1:ef:e8:3e:d3:76:f2:40:5a:dc:fa:
                    6a:3d:1b:4b:ad:76:b0:8c:5c:ee:36:ba:83:ea:30:
                    a8:4c:de:f1:0b:2a:58:41:88:ae:00:89:ab:03:d1:
                    16:82:20:22:76:eb:5e:54:38:12:62:e1:d2:70:24:
                    db:ed:1f:70:d2:64:09:80:2d:e2:b6:9d:ce:1f:f2:
                    bb:21:f3:6c:db:d8:b3:19:7b:8a:50:9f:ef:ec:36:
                    0a:5c:9a:b7:4a:d3:08:a0:39:79:fd:dd:bf:3d:3a:
                    09:25
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                Code Signing
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com
                CA Issuers - URI:http://cacerts.digicert.com/DigiCertHighAssuranceEVRootCA.crt

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertHighAssuranceEVRootCA.crl

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertHighAssuranceEVRootCA.crl

            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.3.2
                  CPS: http://www.digicert.com/ssl-cps-repository.htm
                  User Notice:
                    Explicit Text: 

            X509v3 Subject Key Identifier: 
                8F:E8:7E:F0:6D:32:6A:00:05:23:C7:70:97:6A:3A:90:FF:6B:EA:D4
            X509v3 Authority Key Identifier: 
                keyid:B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3

    Signature Algorithm: sha256WithRSAEncryption
         19:33:4a:0c:81:33:37:db:ad:36:c9:e4:c9:3a:bb:b5:1b:2e:
         7a:a2:e2:f4:43:42:17:9e:bf:4e:a1:4d:e1:b1:db:e9:81:dd:
         9f:01:f2:e4:88:d5:e9:fe:09:fd:21:c1:ec:5d:80:d2:f0:d6:
         c1:43:c2:fe:77:2b:db:f9:d7:91:33:ce:6c:d5:b2:19:3b:e6:
         2e:d6:c9:93:4f:88:40:8e:cd:e1:f5:7e:f1:0f:c6:59:56:72:
         e8:eb:6a:41:bd:1c:d5:46:d5:7c:49:ca:66:38:15:c1:bf:e0:
         91:70:77:87:dc:c9:8d:31:c9:0c:29:a2:33:ed:8d:e2:87:cd:
         89:8d:3f:1b:ff:d5:e0:1a:97:8b:7c:da:6d:fb:a8:c6:b2:3a:
         66:6b:7b:01:b3:cd:d8:a6:34:ec:12:01:ab:95:58:a5:c4:53:
         57:a8:60:e6:e7:02:12:a0:b9:23:64:a2:4d:bb:7c:81:25:64:
         21:be:cf:ee:42:18:43:97:bb:a5:37:06:af:4d:ff:26:a5:4d:
         61:4b:ec:46:41:b8:65:ce:b8:79:9e:08:96:0b:81:8c:8a:3b:
         8f:c7:99:8c:a3:2a:6e:98:6d:5e:61:c6:96:b7:8a:b9:61:2d:
         93:b8:eb:0e:04:43:d7:f5:fe:a6:f0:62:d4:99:6a:a5:c1:c1:
         f0:64:94:80
-----BEGIN CERTIFICATE-----
MIIGvDCCBaSgAwIBAgIQA/G04V86gvEUlniz19hHXDANBgkqhkiG9w0BAQsFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTEyMDQxODEyMDAwMFoXDTI3MDQxODEyMDAwMFowbDEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTErMCkGA1UEAxMiRGlnaUNlcnQgRVYgQ29kZSBTaWduaW5n
IENBIChTSEEyKTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKdT+g+y
tRPxZM+EgPyugDXRttfHoyysGiys8YSsOjUSOpKRulfkxMnzL6hIPLfWbtyXIrpR
eWGvQy8Nt5u0STGuRFg+pKGWp4dPI37DbGUkkFU+ocojfMVC6cR6YkWbfd5jdMue
YyX4hJqarUVPrn0fyBPLdZvJ4eGK+AsMmPTKPtBFqnoepViTNjS+Ky4rMVhmtDIQ
n53wUqHv6D7TdvJAWtz6aj0bS612sIxc7ja6g+owqEze8QsqWEGIrgCJqwPRFoIg
InbrXlQ4EmLh0nAk2+0fcNJkCYAt4radzh/yuyHzbNvYsxl7ilCf7+w2Clyat0rT
CKA5ef3dvz06CSUCAwEAAaOCA1gwggNUMBIGA1UdEwEB/wQIMAYBAf8CAQAwDgYD
VR0PAQH/BAQDAgGGMBMGA1UdJQQMMAoGCCsGAQUFBwMDMH8GCCsGAQUFBwEBBHMw
cTAkBggrBgEFBQcwAYYYaHR0cDovL29jc3AuZGlnaWNlcnQuY29tMEkGCCsGAQUF
BzAChj1odHRwOi8vY2FjZXJ0cy5kaWdpY2VydC5jb20vRGlnaUNlcnRIaWdoQXNz
dXJhbmNlRVZSb290Q0EuY3J0MIGPBgNVHR8EgYcwgYQwQKA+oDyGOmh0dHA6Ly9j
cmwzLmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEhpZ2hBc3N1cmFuY2VFVlJvb3RDQS5j
cmwwQKA+oDyGOmh0dHA6Ly9jcmw0LmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEhpZ2hB
c3N1cmFuY2VFVlJvb3RDQS5jcmwwggHEBgNVHSAEggG7MIIBtzCCAbMGCWCGSAGG
/WwDAjCCAaQwOgYIKwYBBQUHAgEWLmh0dHA6Ly93d3cuZGlnaWNlcnQuY29tL3Nz
bC1jcHMtcmVwb3NpdG9yeS5odG0wggFkBggrBgEFBQcCAjCCAVYeggFSAEEAbgB5
ACAAdQBzAGUAIABvAGYAIAB0AGgAaQBzACAAQwBlAHIAdABpAGYAaQBjAGEAdABl
ACAAYwBvAG4AcwB0AGkAdAB1AHQAZQBzACAAYQBjAGMAZQBwAHQAYQBuAGMAZQAg
AG8AZgAgAHQAaABlACAARABpAGcAaQBDAGUAcgB0ACAAQwBQAC8AQwBQAFMAIABh
AG4AZAAgAHQAaABlACAAUgBlAGwAeQBpAG4AZwAgAFAAYQByAHQAeQAgAEEAZwBy
AGUAZQBtAGUAbgB0ACAAdwBoAGkAYwBoACAAbABpAG0AaQB0ACAAbABpAGEAYgBp
AGwAaQB0AHkAIABhAG4AZAAgAGEAcgBlACAAaQBuAGMAbwByAHAAbwByAGEAdABl
AGQAIABoAGUAcgBlAGkAbgAgAGIAeQAgAHIAZQBmAGUAcgBlAG4AYwBlAC4wHQYD
VR0OBBYEFI/ofvBtMmoABSPHcJdqOpD/a+rUMB8GA1UdIwQYMBaAFLE+w2kD+L9H
AdSYJhoIAu9jZCvDMA0GCSqGSIb3DQEBCwUAA4IBAQAZM0oMgTM32602yeTJOru1
Gy56ouL0Q0IXnr9OoU3hsdvpgd2fAfLkiNXp/gn9IcHsXYDS8NbBQ8L+dyvb+deR
M85s1bIZO+Yu1smTT4hAjs3h9X7xD8ZZVnLo62pBvRzVRtV8ScpmOBXBv+CRcHeH
3MmNMckMKaIz7Y3ih82JjT8b/9XgGpeLfNpt+6jGsjpma3sBs83YpjTsEgGrlVil
xFNXqGDm5wISoLkjZKJNu3yBJWQhvs/uQhhDl7ulNwavTf8mpU1hS+xGQbhlzrh5
ngiWC4GMijuPx5mMoypumG1eYcaWt4q5YS2TuOsOBEPX9f6m8GLUmWqlwcHwZJSA
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            04:e1:e7:a4:dc:5c:f2:f3:6d:c0:2b:42:b8:5d:15:9f
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Validity
            Not Before: Oct 22 12:00:00 2013 GMT
            Not After : Oct 22 12:00:00 2028 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 High Assurance Server CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:b6:e0:2f:c2:24:06:c8:6d:04:5f:d7:ef:0a:64:
                    06:b2:7d:22:26:65:16:ae:42:40:9b:ce:dc:9f:9f:
                    76:07:3e:c3:30:55:87:19:b9:4f:94:0e:5a:94:1f:
                    55:56:b4:c2:02:2a:af:d0:98:ee:0b:40:d7:c4:d0:
                    3b:72:c8:14:9e:ef:90:b1:11:a9:ae:d2:c8:b8:43:
                    3a:d9:0b:0b:d5:d5:95:f5:40:af:c8:1d:ed:4d:9c:
                    5f:57:b7:86:50:68:99:f5:8a:da:d2:c7:05:1f:a8:
                    97:c9:dc:a4:b1:82:84:2d:c6:ad:a5:9c:c7:19:82:
                    a6:85:0f:5e:44:58:2a:37:8f:fd:35:f1:0b:08:27:
                    32:5a:f5:bb:8b:9e:a4:bd:51:d0:27:e2:dd:3b:42:
                    33:a3:05:28:c4:bb:28:cc:9a:ac:2b:23:0d:78:c6:
                    7b:e6:5e:71:b7:4a:3e:08:fb:81:b7:16:16:a1:9d:
                    23:12:4d:e5:d7:92:08:ac:75:a4:9c:ba:cd:17:b2:
                    1e:44:35:65:7f:53:25:39:d1:1c:0a:9a:63:1b:19:
                    92:74:68:0a:37:c2:c2:52:48:cb:39:5a:a2:b6:e1:
                    5d:c1:dd:a0:20:b8:21:a2:93:26:6f:14:4a:21:41:
                    c7:ed:6d:9b:f2:48:2f:f3:03:f5:a2:68:92:53:2f:
                    5e:e3
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication, TLS Web Client Authentication
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertHighAssuranceEVRootCA.crl

            X509v3 Certificate Policies: 
                Policy: X509v3 Any Policy
                  CPS: https://www.digicert.com/CPS

            X509v3 Subject Key Identifier: 
                51:68:FF:90:AF:02:07:75:3C:CC:D9:65:64:62:A2:12:B8:59:72:3B
            X509v3 Authority Key Identifier: 
                keyid:B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3

    Signature Algorithm: sha256WithRSAEncryption
         18:8a:95:89:03:e6:6d:df:5c:fc:1d:68:ea:4a:8f:83:d6:51:
         2f:8d:6b:44:16:9e:ac:63:f5:d2:6e:6c:84:99:8b:aa:81:71:
         84:5b:ed:34:4e:b0:b7:79:92:29:cc:2d:80:6a:f0:8e:20:e1:
         79:a4:fe:03:47:13:ea:f5:86:ca:59:71:7d:f4:04:96:6b:d3:
         59:58:3d:fe:d3:31:25:5c:18:38:84:a3:e6:9f:82:fd:8c:5b:
         98:31:4e:cd:78:9e:1a:fd:85:cb:49:aa:f2:27:8b:99:72:fc:
         3e:aa:d5:41:0b:da:d5:36:a1:bf:1c:6e:47:49:7f:5e:d9:48:
         7c:03:d9:fd:8b:49:a0:98:26:42:40:eb:d6:92:11:a4:64:0a:
         57:54:c4:f5:1d:d6:02:5e:6b:ac:ee:c4:80:9a:12:72:fa:56:
         93:d7:ff:bf:30:85:06:30:bf:0b:7f:4e:ff:57:05:9d:24:ed:
         85:c3:2b:fb:a6:75:a8:ac:2d:16:ef:7d:79:27:b2:eb:c2:9d:
         0b:07:ea:aa:85:d3:01:a3:20:28:41:59:43:28:d2:81:e3:aa:
         f6:ec:7b:3b:77:b6:40:62:80:05:41:45:01:ef:17:06:3e:de:
         c0:33:9b:67:d3:61:2e:72:87:e4:69:fc:12:00:57:40:1e:70:
         f5:1e:c9:b4
-----BEGIN CERTIFICATE-----
MIIEsTCCA5mgAwIBAgIQBOHnpNxc8vNtwCtCuF0VnzANBgkqhkiG9w0BAQsFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTEzMTAyMjEyMDAwMFoXDTI4MTAyMjEyMDAwMFowcDEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTEvMC0GA1UEAxMmRGlnaUNlcnQgU0hBMiBIaWdoIEFzc3Vy
YW5jZSBTZXJ2ZXIgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC2
4C/CJAbIbQRf1+8KZAayfSImZRauQkCbztyfn3YHPsMwVYcZuU+UDlqUH1VWtMIC
Kq/QmO4LQNfE0DtyyBSe75CxEamu0si4QzrZCwvV1ZX1QK/IHe1NnF9Xt4ZQaJn1
itrSxwUfqJfJ3KSxgoQtxq2lnMcZgqaFD15EWCo3j/018QsIJzJa9buLnqS9UdAn
4t07QjOjBSjEuyjMmqwrIw14xnvmXnG3Sj4I+4G3FhahnSMSTeXXkgisdaScus0X
sh5ENWV/UyU50RwKmmMbGZJ0aAo3wsJSSMs5WqK24V3B3aAguCGikyZvFEohQcft
bZvySC/zA/WiaJJTL17jAgMBAAGjggFJMIIBRTASBgNVHRMBAf8ECDAGAQH/AgEA
MA4GA1UdDwEB/wQEAwIBhjAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIw
NAYIKwYBBQUHAQEEKDAmMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2Vy
dC5jb20wSwYDVR0fBEQwQjBAoD6gPIY6aHR0cDovL2NybDQuZGlnaWNlcnQuY29t
L0RpZ2lDZXJ0SGlnaEFzc3VyYW5jZUVWUm9vdENBLmNybDA9BgNVHSAENjA0MDIG
BFUdIAAwKjAoBggrBgEFBQcCARYcaHR0cHM6Ly93d3cuZGlnaWNlcnQuY29tL0NQ
UzAdBgNVHQ4EFgQUUWj/kK8CB3U8zNllZGKiErhZcjswHwYDVR0jBBgwFoAUsT7D
aQP4v0cB1JgmGggC72NkK8MwDQYJKoZIhvcNAQELBQADggEBABiKlYkD5m3fXPwd
aOpKj4PWUS+Na0QWnqxj9dJubISZi6qBcYRb7TROsLd5kinMLYBq8I4g4Xmk/gNH
E+r1hspZcX30BJZr01lYPf7TMSVcGDiEo+afgv2MW5gxTs14nhr9hctJqvIni5ly
/D6q1UEL2tU2ob8cbkdJf17ZSHwD2f2LSaCYJkJA69aSEaRkCldUxPUd1gJea6zu
xICaEnL6VpPX/78whQYwvwt/Tv9XBZ0k7YXDK/umdaisLRbvfXknsuvCnQsH6qqF
0wGjIChBWUMo0oHjqvbsezt3tkBigAVBRQHvFwY+3sAzm2fTYS5yh+Rp/BIAV0Ae
cPUeybQ=
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0c:79:a9:44:b0:8c:11:95:20:92:61:5f:e2:6b:1d:83
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Validity
            Not Before: Oct 22 12:00:00 2013 GMT
            Not After : Oct 22 12:00:00 2028 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 Extended Validation Server CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:d7:53:a4:04:51:f8:99:a6:16:48:4b:67:27:aa:
                    93:49:d0:39:ed:0c:b0:b0:00:87:f1:67:28:86:85:
                    8c:8e:63:da:bc:b1:40:38:e2:d3:f5:ec:a5:05:18:
                    b8:3d:3e:c5:99:17:32:ec:18:8c:fa:f1:0c:a6:64:
                    21:85:cb:07:10:34:b0:52:88:2b:1f:68:9b:d2:b1:
                    8f:12:b0:b3:d2:e7:88:1f:1f:ef:38:77:54:53:5f:
                    80:79:3f:2e:1a:aa:a8:1e:4b:2b:0d:ab:b7:63:b9:
                    35:b7:7d:14:bc:59:4b:df:51:4a:d2:a1:e2:0c:e2:
                    90:82:87:6a:ae:ea:d7:64:d6:98:55:e8:fd:af:1a:
                    50:6c:54:bc:11:f2:fd:4a:f2:9d:bb:7f:0e:f4:d5:
                    be:8e:16:89:12:55:d8:c0:71:34:ee:f6:dc:2d:ec:
                    c4:87:25:86:8d:d8:21:e4:b0:4d:0c:89:dc:39:26:
                    17:dd:f6:d7:94:85:d8:04:21:70:9d:6f:6f:ff:5c:
                    ba:19:e1:45:cb:56:57:28:7e:1c:0d:41:57:aa:b7:
                    b8:27:bb:b1:e4:fa:2a:ef:21:23:75:1a:ad:2d:9b:
                    86:35:8c:9c:77:b5:73:ad:d8:94:2d:e4:f3:0c:9d:
                    ee:c1:4e:62:7e:17:c0:71:9e:2c:de:f1:f9:10:28:
                    19:33
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication, TLS Web Client Authentication
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertHighAssuranceEVRootCA.crl

            X509v3 Certificate Policies: 
                Policy: X509v3 Any Policy
                  CPS: https://www.digicert.com/CPS

            X509v3 Subject Key Identifier: 
                3D:D3:50:A5:D6:A0:AD:EE:F3:4A:60:0A:65:D3:21:D4:F8:F8:D6:0F
            X509v3 Authority Key Identifier: 
                keyid:B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3

    Signature Algorithm: sha256WithRSAEncryption
         9d:b6:d0:90:86:e1:86:02:ed:c5:a0:f0:34:1c:74:c1:8d:76:
         cc:86:0a:a8:f0:4a:8a:42:d6:3f:c8:a9:4d:ad:7c:08:ad:e6:
         b6:50:b8:a2:1a:4d:88:07:b1:29:21:dc:e7:da:c6:3c:21:e0:
         e3:11:49:70:ac:7a:1d:01:a4:ca:11:3a:57:ab:7d:57:2a:40:
         74:fd:d3:1d:85:18:50:df:57:47:75:a1:7d:55:20:2e:47:37:
         50:72:8c:7f:82:1b:d2:62:8f:2d:03:5a:da:c3:c8:a1:ce:2c:
         52:a2:00:63:eb:73:ba:71:c8:49:27:23:97:64:85:9e:38:0e:
         ad:63:68:3c:ba:52:81:58:79:a3:2c:0c:df:de:6d:eb:31:f2:
         ba:a0:7c:6c:f1:2c:d4:e1:bd:77:84:37:03:ce:32:b5:c8:9a:
         81:1a:4a:92:4e:3b:46:9a:85:fe:83:a2:f9:9e:8c:a3:cc:0d:
         5e:b3:3d:cf:04:78:8f:14:14:7b:32:9c:c7:00:a6:5c:c4:b5:
         a1:55:8d:5a:56:68:a4:22:70:aa:3c:81:71:d9:9d:a8:45:3b:
         f4:e5:f6:a2:51:dd:c7:7b:62:e8:6f:0c:74:eb:b8:da:f8:bf:
         87:0d:79:50:91:90:9b:18:3b:91:59:27:f1:35:28:13:ab:26:
         7e:d5:f7:7a
-----BEGIN CERTIFICATE-----
MIIEtjCCA56gAwIBAgIQDHmpRLCMEZUgkmFf4msdgzANBgkqhkiG9w0BAQsFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTEzMTAyMjEyMDAwMFoXDTI4MTAyMjEyMDAwMFowdTEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTE0MDIGA1UEAxMrRGlnaUNlcnQgU0hBMiBFeHRlbmRlZCBW
YWxpZGF0aW9uIFNlcnZlciBDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBANdTpARR+JmmFkhLZyeqk0nQOe0MsLAAh/FnKIaFjI5j2ryxQDji0/XspQUY
uD0+xZkXMuwYjPrxDKZkIYXLBxA0sFKIKx9om9KxjxKws9LniB8f7zh3VFNfgHk/
LhqqqB5LKw2rt2O5Nbd9FLxZS99RStKh4gzikIKHaq7q12TWmFXo/a8aUGxUvBHy
/Urynbt/DvTVvo4WiRJV2MBxNO723C3sxIclho3YIeSwTQyJ3DkmF93215SF2AQh
cJ1vb/9cuhnhRctWVyh+HA1BV6q3uCe7seT6Ku8hI3UarS2bhjWMnHe1c63YlC3k
8wyd7sFOYn4XwHGeLN7x+RAoGTMCAwEAAaOCAUkwggFFMBIGA1UdEwEB/wQIMAYB
Af8CAQAwDgYDVR0PAQH/BAQDAgGGMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEF
BQcDAjA0BggrBgEFBQcBAQQoMCYwJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRp
Z2ljZXJ0LmNvbTBLBgNVHR8ERDBCMECgPqA8hjpodHRwOi8vY3JsNC5kaWdpY2Vy
dC5jb20vRGlnaUNlcnRIaWdoQXNzdXJhbmNlRVZSb290Q0EuY3JsMD0GA1UdIAQ2
MDQwMgYEVR0gADAqMCgGCCsGAQUFBwIBFhxodHRwczovL3d3dy5kaWdpY2VydC5j
b20vQ1BTMB0GA1UdDgQWBBQ901Cl1qCt7vNKYApl0yHU+PjWDzAfBgNVHSMEGDAW
gBSxPsNpA/i/RwHUmCYaCALvY2QrwzANBgkqhkiG9w0BAQsFAAOCAQEAnbbQkIbh
hgLtxaDwNBx0wY12zIYKqPBKikLWP8ipTa18CK3mtlC4ohpNiAexKSHc59rGPCHg
4xFJcKx6HQGkyhE6V6t9VypAdP3THYUYUN9XR3WhfVUgLkc3UHKMf4Ib0mKPLQNa
2sPIoc4sUqIAY+tzunHISScjl2SFnjgOrWNoPLpSgVh5oywM395t6zHyuqB8bPEs
1OG9d4Q3A84ytciagRpKkk47RpqF/oOi+Z6Mo8wNXrM9zwR4jxQUezKcxwCmXMS1
oVWNWlZopCJwqjyBcdmdqEU79OX2olHdx3ti6G8MdOu42vi/hw15UJGQmxg7kVkn
8TUoE6smftX3eg==
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            04:09:18:1b:5f:d5:bb:66:75:53:43:b5:6f:95:50:08
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root CA
        Validity
            Not Before: Oct 22 12:00:00 2013 GMT
            Not After : Oct 22 12:00:00 2028 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 Assured ID Code Signing CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:f8:d3:b3:1c:7f:0e:11:af:67:77:07:d3:0b:31:
                    49:19:cf:d0:fb:45:99:b1:3a:db:44:f5:7f:e5:a8:
                    9d:db:32:d7:71:ea:76:9d:05:2e:b7:8f:fa:92:43:
                    c0:a5:f9:89:d4:37:19:d7:b6:aa:f0:9c:86:a5:d8:
                    25:ac:0e:79:28:3a:7e:e9:d1:67:d3:c6:fb:29:27:
                    c7:d3:7b:23:94:e4:91:23:96:90:77:82:f9:a1:84:
                    23:66:12:54:33:50:74:b1:28:26:bb:24:69:c2:c2:
                    52:f2:14:67:8a:89:45:d4:2d:a1:a3:e9:88:2c:20:
                    95:ae:1c:4a:87:08:df:0c:f5:e2:4d:60:18:be:aa:
                    c4:b2:ae:70:31:66:33:71:3e:ac:70:a2:ab:ce:7f:
                    e9:7c:cb:92:a1:e5:3b:31:1c:cf:ea:f2:0a:e4:57:
                    bb:4a:b5:e9:74:e6:2b:fe:6c:cb:7e:74:39:36:0d:
                    90:ef:e4:b5:4e:a4:a9:ea:6a:0a:ab:84:f3:ac:67:
                    4e:b5:c4:f7:8c:d1:20:25:23:eb:08:64:3e:52:96:
                    c1:f2:0f:12:f4:c5:8e:0f:c1:a2:e8:2c:51:f7:73:
                    bc:bd:85:b1:62:83:73:41:82:07:e4:38:8b:6a:73:
                    20:d0:0f:64:73:3c:9e:9f:a6:33:a9:fd:19:df:25:
                    93:d1
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                Code Signing
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com
                CA Issuers - URI:http://cacerts.digicert.com/DigiCertAssuredIDRootCA.crt

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertAssuredIDRootCA.crl

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl

            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.0.2.4
                  CPS: https://www.digicert.com/CPS
                Policy: 2.16.840.1.114412.3

            X509v3 Subject Key Identifier: 
                5A:C4:B9:7B:2A:0A:A3:A5:EA:71:03:C0:60:F9:2D:F6:65:75:0E:58
            X509v3 Authority Key Identifier: 
                keyid:45:EB:A2:AF:F4:92:CB:82:31:2D:51:8B:A7:A7:21:9D:F3:6D:C8:0F

    Signature Algorithm: sha256WithRSAEncryption
         3e:ec:0d:5a:24:b3:f3:22:d1:15:c8:2c:7c:25:29:76:a8:1d:
         5d:1c:2d:3a:1a:c4:ef:30:61:d7:7e:0b:60:fd:c3:3d:0f:c4:
         af:8b:fd:ef:2a:df:20:55:37:b0:e1:f6:d1:92:75:0f:51:b4:
         6e:a5:8e:5a:e2:5e:24:81:4e:10:a4:ee:3f:71:8e:63:0e:13:
         4b:ad:d7:5f:44:79:f3:36:14:06:8a:f7:9c:46:4e:5c:ff:90:
         b1:1b:07:0e:91:15:fb:ba:af:b5:51:c2:8d:24:ae:24:c6:c7:
         27:2a:a1:29:28:1a:3a:71:28:02:3c:2e:91:a3:c0:25:11:e2:
         9c:14:47:a1:7a:68:68:af:9b:a7:5c:20:5c:d9:71:b1:0c:8f:
         bb:a8:f8:c5:12:68:9f:cf:40:cb:40:44:a5:13:f0:e6:64:0c:
         25:08:42:32:b2:36:8a:24:02:fe:2f:72:7e:1c:d7:49:45:96:
         e8:59:1d:e9:fa:74:64:6b:b2:eb:66:43:da:b3:b0:8c:d5:e9:
         0d:dd:f6:01:20:ce:99:31:63:3d:08:1a:18:b3:81:9b:4f:c6:
         93:10:06:fc:07:81:fa:8b:da:f9:82:49:f7:62:6e:a1:53:fa:
         12:94:18:85:2e:92:91:ea:68:6c:44:32:b2:66:a1:e7:18:a4:
         9a:64:51:ef
-----BEGIN CERTIFICATE-----
MIIFMDCCBBigAwIBAgIQBAkYG1/Vu2Z1U0O1b5VQCDANBgkqhkiG9w0BAQsFADBl
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJv
b3QgQ0EwHhcNMTMxMDIyMTIwMDAwWhcNMjgxMDIyMTIwMDAwWjByMQswCQYDVQQG
EwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNl
cnQuY29tMTEwLwYDVQQDEyhEaWdpQ2VydCBTSEEyIEFzc3VyZWQgSUQgQ29kZSBT
aWduaW5nIENBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+NOzHH8O
Ea9ndwfTCzFJGc/Q+0WZsTrbRPV/5aid2zLXcep2nQUut4/6kkPApfmJ1DcZ17aq
8JyGpdglrA55KDp+6dFn08b7KSfH03sjlOSRI5aQd4L5oYQjZhJUM1B0sSgmuyRp
wsJS8hRniolF1C2ho+mILCCVrhxKhwjfDPXiTWAYvqrEsq5wMWYzcT6scKKrzn/p
fMuSoeU7MRzP6vIK5Fe7SrXpdOYr/mzLfnQ5Ng2Q7+S1TqSp6moKq4TzrGdOtcT3
jNEgJSPrCGQ+UpbB8g8S9MWOD8Gi6CxR93O8vYWxYoNzQYIH5DiLanMg0A9kczye
n6Yzqf0Z3yWT0QIDAQABo4IBzTCCAckwEgYDVR0TAQH/BAgwBgEB/wIBADAOBgNV
HQ8BAf8EBAMCAYYwEwYDVR0lBAwwCgYIKwYBBQUHAwMweQYIKwYBBQUHAQEEbTBr
MCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2VydC5jb20wQwYIKwYBBQUH
MAKGN2h0dHA6Ly9jYWNlcnRzLmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEFzc3VyZWRJ
RFJvb3RDQS5jcnQwgYEGA1UdHwR6MHgwOqA4oDaGNGh0dHA6Ly9jcmw0LmRpZ2lj
ZXJ0LmNvbS9EaWdpQ2VydEFzc3VyZWRJRFJvb3RDQS5jcmwwOqA4oDaGNGh0dHA6
Ly9jcmwzLmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEFzc3VyZWRJRFJvb3RDQS5jcmww
TwYDVR0gBEgwRjA4BgpghkgBhv1sAAIEMCowKAYIKwYBBQUHAgEWHGh0dHBzOi8v
d3d3LmRpZ2ljZXJ0LmNvbS9DUFMwCgYIYIZIAYb9bAMwHQYDVR0OBBYEFFrEuXsq
CqOl6nEDwGD5LfZldQ5YMB8GA1UdIwQYMBaAFEXroq/0ksuCMS1Ri6enIZ3zbcgP
MA0GCSqGSIb3DQEBCwUAA4IBAQA+7A1aJLPzItEVyCx8JSl2qB1dHC06GsTvMGHX
fgtg/cM9D8Svi/3vKt8gVTew4fbRknUPUbRupY5a4l4kgU4QpO4/cY5jDhNLrddf
RHnzNhQGivecRk5c/5CxGwcOkRX7uq+1UcKNJK4kxscnKqEpKBo6cSgCPC6Ro8Al
EeKcFEehemhor5unXCBc2XGxDI+7qPjFEmifz0DLQESlE/DmZAwlCEIysjaKJAL+
L3J+HNdJRZboWR3p+nRka7LrZkPas7CM1ekN3fYBIM6ZMWM9CBoYs4GbT8aTEAb8
B4H6i9r5gkn3Ym6hU/oSlBiFLpKR6mhsRDKyZqHnGKSaZFHv
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0b:7e:10:90:3c:38:49:0f:fa:2f:67:9a:87:a1:a7:b9
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance EV Root CA
        Validity
            Not Before: Oct 22 12:00:00 2013 GMT
            Not After : Oct 22 12:00:00 2028 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 High Assurance Code Signing CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:b4:4a:5e:7d:07:0f:41:de:c4:f5:76:16:36:bd:
                    71:ff:cf:3f:4f:73:4b:9c:d1:0d:fe:4a:cb:57:58:
                    5e:85:16:dd:02:15:54:99:f0:8f:3c:2f:4d:02:78:
                    10:68:c8:d8:35:4b:3f:c1:f7:67:ce:98:1c:ae:33:
                    b9:2d:1d:a4:0a:54:93:c4:85:a2:df:35:b1:f5:f1:
                    3c:a7:b3:34:fb:5d:48:c9:46:c9:62:44:bc:48:99:
                    eb:28:49:53:c3:3d:8f:c0:0e:de:35:98:e9:62:51:
                    df:3d:6b:40:61:ee:04:41:da:cf:a7:5c:56:96:d1:
                    f9:4c:b7:44:84:87:98:69:e5:82:b9:13:e6:55:bf:
                    c8:92:70:92:0a:31:6f:7f:8b:32:ab:cf:6b:5a:9f:
                    62:c4:3e:ee:be:ed:59:a4:53:7f:0b:f1:52:88:8a:
                    7b:0a:67:24:cb:90:cd:ec:d2:4d:34:4c:b0:e1:b5:
                    9f:9c:c6:f6:6f:2c:cd:e6:ca:53:74:01:9f:67:35:
                    de:38:49:2d:ce:ed:39:44:82:19:79:4e:1a:b2:b5:
                    fb:bb:78:f0:49:66:a7:cf:fa:5c:96:75:92:8b:1a:
                    72:d9:ff:50:92:53:cc:3e:c2:43:32:09:1a:86:13:
                    69:3c:fb:81:32:33:32:64:75:73:28:26:1d:08:30:
                    3b:07
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                Code Signing
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com
                CA Issuers - URI:http://cacerts.digicert.com/DigiCertHighAssuranceEVRootCA.crt

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertHighAssuranceEVRootCA.crl

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertHighAssuranceEVRootCA.crl

            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.0.2.4
                  CPS: https://www.digicert.com/CPS
                Policy: 2.16.840.1.114412.3

            X509v3 Subject Key Identifier: 
                67:9D:0F:20:09:0C:CC:8A:3A:E5:82:46:72:62:FC:F1:CC:90:E5:40
            X509v3 Authority Key Identifier: 
                keyid:B1:3E:C3:69:03:F8:BF:47:01:D4:98:26:1A:08:02:EF:63:64:2B:C3

    Signature Algorithm: sha256WithRSAEncryption
         6a:0e:ff:7e:13:7c:06:a5:4b:c0:2e:8c:f9:53:64:09:e2:ba:
         58:91:30:50:ec:cc:9f:e1:d3:a8:2f:48:46:36:18:29:d0:78:
         28:5f:98:56:40:0f:1e:ba:bd:b1:3b:87:5c:dc:5b:d8:20:0d:
         ed:1a:16:4d:d5:11:24:21:4b:f1:27:69:90:13:eb:11:a1:01:
         da:fd:b5:4e:79:59:75:bd:38:2a:6a:c3:f6:8e:41:2b:8a:a2:
         8b:d7:2c:51:51:d9:9c:a0:c8:e3:4e:ba:6c:a8:47:d2:4e:d1:
         68:1f:8c:02:57:3b:b3:29:6a:8e:6a:20:2a:b9:f2:00:62:64:
         ba:c8:e9:00:f9:cc:a4:d4:ba:9a:35:d8:af:2c:65:6c:16:7c:
         58:21:de:4a:30:d0:fa:eb:24:5d:06:c9:9d:16:b7:ad:4a:45:
         d3:25:e2:0c:f0:40:aa:5c:4d:ac:7e:cd:06:82:b9:76:46:69:
         08:d8:32:b6:82:fe:e3:a9:58:34:43:1b:8e:67:67:97:3f:68:
         31:16:36:38:95:3e:87:f7:c7:c3:af:9d:7a:77:19:d9:de:93:
         b5:fd:6e:2b:fc:94:f9:3d:b7:4c:12:35:2c:30:be:e8:8d:9e:
         05:70:9a:48:13:f4:8c:d6:e7:1e:ac:38:e7:a8:f3:ad:0c:b7:
         7a:ec:67:ed
-----BEGIN CERTIFICATE-----
MIIFTzCCBDegAwIBAgIQC34QkDw4SQ/6L2eah6GnuTANBgkqhkiG9w0BAQsFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTEzMTAyMjEyMDAwMFoXDTI4MTAyMjEyMDAwMFowdjEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTE1MDMGA1UEAxMsRGlnaUNlcnQgU0hBMiBIaWdoIEFzc3Vy
YW5jZSBDb2RlIFNpZ25pbmcgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
AoIBAQC0Sl59Bw9B3sT1dhY2vXH/zz9Pc0uc0Q3+SstXWF6FFt0CFVSZ8I88L00C
eBBoyNg1Sz/B92fOmByuM7ktHaQKVJPEhaLfNbH18TynszT7XUjJRsliRLxImeso
SVPDPY/ADt41mOliUd89a0Bh7gRB2s+nXFaW0flMt0SEh5hp5YK5E+ZVv8iScJIK
MW9/izKrz2tan2LEPu6+7VmkU38L8VKIinsKZyTLkM3s0k00TLDhtZ+cxvZvLM3m
ylN0AZ9nNd44SS3O7TlEghl5Thqytfu7ePBJZqfP+lyWdZKLGnLZ/1CSU8w+wkMy
CRqGE2k8+4EyMzJkdXMoJh0IMDsHAgMBAAGjggHhMIIB3TASBgNVHRMBAf8ECDAG
AQH/AgEAMA4GA1UdDwEB/wQEAwIBhjATBgNVHSUEDDAKBggrBgEFBQcDAzB/Bggr
BgEFBQcBAQRzMHEwJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRpZ2ljZXJ0LmNv
bTBJBggrBgEFBQcwAoY9aHR0cDovL2NhY2VydHMuZGlnaWNlcnQuY29tL0RpZ2lD
ZXJ0SGlnaEFzc3VyYW5jZUVWUm9vdENBLmNydDCBjwYDVR0fBIGHMIGEMECgPqA8
hjpodHRwOi8vY3JsNC5kaWdpY2VydC5jb20vRGlnaUNlcnRIaWdoQXNzdXJhbmNl
RVZSb290Q0EuY3JsMECgPqA8hjpodHRwOi8vY3JsMy5kaWdpY2VydC5jb20vRGln
aUNlcnRIaWdoQXNzdXJhbmNlRVZSb290Q0EuY3JsME8GA1UdIARIMEYwOAYKYIZI
AYb9bAACBDAqMCgGCCsGAQUFBwIBFhxodHRwczovL3d3dy5kaWdpY2VydC5jb20v
Q1BTMAoGCGCGSAGG/WwDMB0GA1UdDgQWBBRnnQ8gCQzMijrlgkZyYvzxzJDlQDAf
BgNVHSMEGDAWgBSxPsNpA/i/RwHUmCYaCALvY2QrwzANBgkqhkiG9w0BAQsFAAOC
AQEAag7/fhN8BqVLwC6M+VNkCeK6WJEwUOzMn+HTqC9IRjYYKdB4KF+YVkAPHrq9
sTuHXNxb2CAN7RoWTdURJCFL8SdpkBPrEaEB2v21TnlZdb04KmrD9o5BK4qii9cs
UVHZnKDI4066bKhH0k7RaB+MAlc7sylqjmogKrnyAGJkusjpAPnMpNS6mjXYryxl
bBZ8WCHeSjDQ+uskXQbJnRa3rUpF0yXiDPBAqlxNrH7NBoK5dkZpCNgytoL+46lY
NEMbjmdnlz9oMRY2OJU+h/fHw6+dencZ2d6Ttf1uK/yU+T23TBI1LDC+6I2eBXCa
SBP0jNbnHqw456jzrQy3euxn7Q==
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            04:ae:79:60:66:66:90:1a:b9:c5:7f:a6:6c:5b:dc:cd
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root CA
        Validity
            Not Before: Nov  5 12:00:00 2013 GMT
            Not After : Nov  5 12:00:00 2028 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 Assured ID CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:dc:f8:11:23:3f:6a:b5:ef:c0:27:79:1b:2d:05:
                    87:a2:10:43:31:df:0e:d4:15:1b:4f:77:a4:22:ce:
                    3e:8c:70:f0:be:07:8e:dd:27:2a:bc:01:1d:b6:2c:
                    0a:ca:dd:69:58:1f:41:ed:6a:05:1f:da:63:78:59:
                    1e:22:2c:2b:f8:ba:7e:c9:35:3b:56:f1:1f:7c:42:
                    7e:25:b0:23:19:c6:45:38:d7:3d:44:f9:20:7c:60:
                    ae:1c:b0:5a:18:04:be:8f:3b:f7:a9:f0:94:9a:a6:
                    0c:63:49:b6:41:17:53:40:32:bd:4f:e6:50:7a:50:
                    1f:25:45:e8:f1:89:af:cd:ff:5e:ff:50:f3:f0:17:
                    11:ff:bd:c5:89:f5:bd:62:b9:d8:fb:8e:45:04:85:
                    6d:99:cd:c5:48:4d:fd:26:ab:02:36:45:ea:36:d7:
                    5f:6a:e2:1a:82:41:0a:dc:e9:d9:f5:91:17:8e:c6:
                    21:ad:38:3c:13:e6:1e:7f:66:9c:d5:4d:d9:46:da:
                    f6:cf:52:a7:7d:3b:24:c3:b4:c1:51:35:c6:9b:eb:
                    d4:ef:7e:e4:7b:03:e1:44:bd:7a:0b:37:e0:4f:cb:
                    82:8a:0d:71:18:e4:a6:e0:89:2c:1b:f3:b2:73:3c:
                    c1:b0:92:9e:18:7d:bd:ab:7d:c5:7d:08:a1:2c:c2:
                    9f:73
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertAssuredIDRootCA.crl

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl

            X509v3 Extended Key Usage: 
                TLS Web Client Authentication, E-mail Protection
            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.0.2.4
                  CPS: https://www.digicert.com/CPS
                  User Notice:
                    Explicit Text: 

            X509v3 Subject Key Identifier: 
                E7:02:23:80:00:4F:D8:D7:BC:94:0B:D9:3F:74:39:49:32:3C:8A:79
            X509v3 Authority Key Identifier: 
                keyid:45:EB:A2:AF:F4:92:CB:82:31:2D:51:8B:A7:A7:21:9D:F3:6D:C8:0F

    Signature Algorithm: sha256WithRSAEncryption
         4e:d4:89:27:b9:fd:1d:87:77:7e:0f:28:05:90:f1:0a:2f:c9:
         3b:3e:bd:93:9c:90:c6:af:fa:91:51:87:32:54:6b:e8:ca:c1:
         71:51:5c:99:8e:b5:fa:e0:62:19:99:a6:07:9a:7c:13:27:db:
         e5:02:a2:84:12:d9:15:f6:0a:44:57:c4:34:8f:6d:73:1f:3a:
         d6:a0:01:3d:de:e3:82:e4:45:6b:b2:eb:25:fd:80:5e:1d:39:
         90:be:5b:42:f6:ce:91:60:e3:f9:a0:56:0c:48:10:ce:33:68:
         7b:cd:93:ab:f0:6a:cb:4f:58:0f:b9:7b:5d:1b:04:81:9e:a7:
         22:e2:57:49:27:ee:92:93:08:be:69:08:53:02:3a:2e:5d:c8:
         19:f4:50:80:4e:02:d9:b0:5d:91:b4:93:87:f9:96:96:81:55:
         40:8e:66:ed:82:d8:bb:28:5f:ed:4c:61:ed:06:58:28:19:53:
         11:44:bc:47:20:29:b2:04:6c:d8:89:54:99:3b:75:db:67:78:
         2c:1e:92:78:c6:55:3d:58:12:11:32:6e:bd:43:0c:4d:34:db:
         7b:c6:42:db:ec:be:ab:9b:61:4c:06:92:e3:8e:df:21:2a:50:
         57:04:c0:60:a2:35:d5:24:9e:66:37:09:49:cf:d5:f8:65:b3:
         e0:e2:6e:c2
-----BEGIN CERTIFICATE-----
MIIGTjCCBTagAwIBAgIQBK55YGZmkBq5xX+mbFvczTANBgkqhkiG9w0BAQsFADBl
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJv
b3QgQ0EwHhcNMTMxMTA1MTIwMDAwWhcNMjgxMTA1MTIwMDAwWjBlMQswCQYDVQQG
EwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNl
cnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBTSEEyIEFzc3VyZWQgSUQgQ0EwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDc+BEjP2q178AneRstBYeiEEMx
3w7UFRtPd6Qizj6McPC+B47dJyq8AR22LArK3WlYH0HtagUf2mN4WR4iLCv4un7J
NTtW8R98Qn4lsCMZxkU41z1E+SB8YK4csFoYBL6PO/ep8JSapgxjSbZBF1NAMr1P
5lB6UB8lRejxia/N/17/UPPwFxH/vcWJ9b1iudj7jkUEhW2ZzcVITf0mqwI2Reo2
119q4hqCQQrc6dn1kReOxiGtODwT5h5/ZpzVTdlG2vbPUqd9OyTDtMFRNcab69Tv
fuR7A+FEvXoLN+BPy4KKDXEY5KbgiSwb87JzPMGwkp4Yfb2rfcV9CKEswp9zAgMB
AAGjggL4MIIC9DASBgNVHRMBAf8ECDAGAQH/AgEAMA4GA1UdDwEB/wQEAwIBhjA0
BggrBgEFBQcBAQQoMCYwJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRpZ2ljZXJ0
LmNvbTCBgQYDVR0fBHoweDA6oDigNoY0aHR0cDovL2NybDQuZGlnaWNlcnQuY29t
L0RpZ2lDZXJ0QXNzdXJlZElEUm9vdENBLmNybDA6oDigNoY0aHR0cDovL2NybDMu
ZGlnaWNlcnQuY29tL0RpZ2lDZXJ0QXNzdXJlZElEUm9vdENBLmNybDAdBgNVHSUE
FjAUBggrBgEFBQcDAgYIKwYBBQUHAwQwggGzBgNVHSAEggGqMIIBpjCCAaIGCmCG
SAGG/WwAAgQwggGSMCgGCCsGAQUFBwIBFhxodHRwczovL3d3dy5kaWdpY2VydC5j
b20vQ1BTMIIBZAYIKwYBBQUHAgIwggFWHoIBUgBBAG4AeQAgAHUAcwBlACAAbwBm
ACAAdABoAGkAcwAgAEMAZQByAHQAaQBmAGkAYwBhAHQAZQAgAGMAbwBuAHMAdABp
AHQAdQB0AGUAcwAgAGEAYwBjAGUAcAB0AGEAbgBjAGUAIABvAGYAIAB0AGgAZQAg
AEQAaQBnAGkAQwBlAHIAdAAgAEMAUAAvAEMAUABTACAAYQBuAGQAIAB0AGgAZQAg
AFIAZQBsAHkAaQBuAGcAIABQAGEAcgB0AHkAIABBAGcAcgBlAGUAbQBlAG4AdAAg
AHcAaABpAGMAaAAgAGwAaQBtAGkAdAAgAGwAaQBhAGIAaQBsAGkAdAB5ACAAYQBu
AGQAIABhAHIAZQAgAGkAbgBjAG8AcgBwAG8AcgBhAHQAZQBkACAAaABlAHIAZQBp
AG4AIABiAHkAIAByAGUAZgBlAHIAZQBuAGMAZQAuMB0GA1UdDgQWBBTnAiOAAE/Y
17yUC9k/dDlJMjyKeTAfBgNVHSMEGDAWgBRF66Kv9JLLgjEtUYunpyGd823IDzAN
BgkqhkiG9w0BAQsFAAOCAQEATtSJJ7n9HYd3fg8oBZDxCi/JOz69k5yQxq/6kVGH
MlRr6MrBcVFcmY61+uBiGZmmB5p8Eyfb5QKihBLZFfYKRFfENI9tcx861qABPd7j
guRFa7LrJf2AXh05kL5bQvbOkWDj+aBWDEgQzjNoe82Tq/Bqy09YD7l7XRsEgZ6n
IuJXSSfukpMIvmkIUwI6Ll3IGfRQgE4C2bBdkbSTh/mWloFVQI5m7YLYuyhf7Uxh
7QZYKBlTEUS8RyApsgRs2IlUmTt122d4LB6SeMZVPVgSETJuvUMMTTTbe8ZC2+y+
q5thTAaS447fISpQVwTAYKI11SSeZjcJSc/V+GWz4OJuwg==
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            09:c7:b7:db:b2:78:24:37:91:19:5e:6a:b1:33:87:10
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Assured ID Root CA
        Validity
            Not Before: Nov  5 12:00:00 2013 GMT
            Not After : Nov  5 12:00:00 2028 GMT
        Subject: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert Document Signing CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:ca:51:16:3a:77:99:f5:16:24:a9:b5:e4:e6:de:
                    95:7a:31:0d:4a:9d:6a:cb:7d:25:7c:a4:9b:9c:84:
                    fa:08:34:68:f5:09:26:3e:31:45:b0:21:48:4c:d5:
                    99:95:86:9f:1a:9d:77:f6:bb:67:0f:a5:5a:9c:a5:
                    0f:fb:9b:e5:d2:18:2a:eb:c6:54:0c:62:88:23:62:
                    ab:61:96:6e:4f:17:62:1e:81:ac:e9:10:e6:c3:9d:
                    b9:b6:f6:e5:e1:ec:c6:ae:4b:2d:9b:48:87:63:bc:
                    1d:9f:60:a5:16:8f:fd:92:eb:71:34:c8:5e:c6:15:
                    d8:8a:4a:fc:13:36:30:aa:0d:36:1c:2d:d0:0e:4f:
                    b5:31:e4:59:4e:f4:0b:32:b9:62:6c:75:65:f8:7a:
                    89:b3:93:5b:34:48:b0:f8:4b:f5:20:56:83:a5:48:
                    f5:8d:94:63:4c:03:27:a9:30:55:67:95:15:1b:29:
                    40:23:9d:b4:a8:c7:0d:65:00:b3:11:e8:22:4f:6f:
                    44:98:e4:fe:72:75:e2:ba:ac:31:4d:5d:e1:0f:a6:
                    94:2c:a0:c8:28:dd:a8:46:05:38:8e:10:14:9a:aa:
                    3f:84:16:e2:4e:5a:92:5f:85:b3:46:cd:41:14:6f:
                    fe:ca:fa:af:bd:ad:35:0a:2d:81:2c:9f:b6:fc:cd:
                    0e:d3
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            Authority Information Access: 
                OCSP - URI:http://ocsp.digicert.com

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl4.digicert.com/DigiCertAssuredIDRootCA.crl

                Full Name:
                  URI:http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl

            X509v3 Extended Key Usage: 
                TLS Web Client Authentication, E-mail Protection
            X509v3 Certificate Policies: 
                Policy: 2.16.840.1.114412.0.2.4
                  CPS: https://www.digicert.com/CPS
                  User Notice:
                    Explicit Text: 
                Policy: 2.16.840.1.114412.3.21

            X509v3 Subject Key Identifier: 
                EF:CE:35:93:CE:F6:86:C5:F8:84:F5:0C:E7:5A:6F:D9:2F:4B:E3:64
            X509v3 Authority Key Identifier: 
                keyid:45:EB:A2:AF:F4:92:CB:82:31:2D:51:8B:A7:A7:21:9D:F3:6D:C8:0F

    Signature Algorithm: sha256WithRSAEncryption
         59:70:cd:f7:30:e3:b4:fa:53:cb:7c:01:84:f2:14:06:09:c2:
         fb:64:4b:65:5e:a2:5e:f9:0c:54:bd:dd:66:ad:5c:84:f8:df:
         8f:7a:6a:77:b3:e8:2c:92:0f:45:99:72:a1:f6:75:38:92:25:
         29:85:b2:6e:41:6b:05:dd:ff:fa:e1:9e:b6:1c:c4:28:5c:21:
         7f:a3:e7:d1:a1:31:e6:06:8a:9f:52:8b:8c:e3:4c:fa:21:ea:
         f5:ad:cf:1f:d8:55:8d:dc:d4:2b:b1:77:e7:83:10:a9:b7:5c:
         c2:31:39:98:98:0a:79:77:21:f5:f7:d9:70:5d:c5:66:02:11:
         9b:95:c0:cb:ad:89:36:0a:c0:83:eb:0a:81:9d:bd:dd:b8:f0:
         d3:07:51:12:8b:9b:6d:ca:ee:44:b3:08:f4:91:79:fe:f5:e0:
         62:11:f2:f7:79:69:fb:9d:99:c9:c7:af:01:e7:4e:83:ed:be:
         1d:cc:10:80:c5:6f:36:cf:35:54:d1:8b:47:8b:76:59:c0:0c:
         09:c9:6a:b1:64:cb:2d:f7:4f:e9:3b:a7:36:20:49:e4:f7:17:
         f0:7b:0f:21:4f:47:d3:8e:d2:90:a4:69:6d:4f:d2:14:e2:c7:
         65:d2:3a:8c:b1:4f:95:14:1f:9c:71:33:49:f6:55:ac:be:a0:
         b0:11:ea:80
-----BEGIN CERTIFICATE-----
MIIGXDCCBUSgAwIBAgIQCce327J4JDeRGV5qsTOHEDANBgkqhkiG9w0BAQsFADBl
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSQwIgYDVQQDExtEaWdpQ2VydCBBc3N1cmVkIElEIFJv
b3QgQ0EwHhcNMTMxMTA1MTIwMDAwWhcNMjgxMTA1MTIwMDAwWjBmMQswCQYDVQQG
EwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNl
cnQuY29tMSUwIwYDVQQDExxEaWdpQ2VydCBEb2N1bWVudCBTaWduaW5nIENBMIIB
IjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAylEWOneZ9RYkqbXk5t6VejEN
Sp1qy30lfKSbnIT6CDRo9QkmPjFFsCFITNWZlYafGp139rtnD6VanKUP+5vl0hgq
68ZUDGKII2KrYZZuTxdiHoGs6RDmw525tvbl4ezGrkstm0iHY7wdn2ClFo/9kutx
NMhexhXYikr8EzYwqg02HC3QDk+1MeRZTvQLMrlibHVl+HqJs5NbNEiw+Ev1IFaD
pUj1jZRjTAMnqTBVZ5UVGylAI520qMcNZQCzEegiT29EmOT+cnXiuqwxTV3hD6aU
LKDIKN2oRgU4jhAUmqo/hBbiTlqSX4WzRs1BFG/+yvqvva01Ci2BLJ+2/M0O0wID
AQABo4IDBTCCAwEwEgYDVR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYw
NAYIKwYBBQUHAQEEKDAmMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2Vy
dC5jb20wgYEGA1UdHwR6MHgwOqA4oDaGNGh0dHA6Ly9jcmw0LmRpZ2ljZXJ0LmNv
bS9EaWdpQ2VydEFzc3VyZWRJRFJvb3RDQS5jcmwwOqA4oDaGNGh0dHA6Ly9jcmwz
LmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEFzc3VyZWRJRFJvb3RDQS5jcmwwHQYDVR0l
BBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMEMIIBwAYDVR0gBIIBtzCCAbMwggGiBgpg
hkgBhv1sAAIEMIIBkjAoBggrBgEFBQcCARYcaHR0cHM6Ly93d3cuZGlnaWNlcnQu
Y29tL0NQUzCCAWQGCCsGAQUFBwICMIIBVh6CAVIAQQBuAHkAIAB1AHMAZQAgAG8A
ZgAgAHQAaABpAHMAIABDAGUAcgB0AGkAZgBpAGMAYQB0AGUAIABjAG8AbgBzAHQA
aQB0AHUAdABlAHMAIABhAGMAYwBlAHAAdABhAG4AYwBlACAAbwBmACAAdABoAGUA
IABEAGkAZwBpAEMAZQByAHQAIABDAFAALwBDAFAAUwAgAGEAbgBkACAAdABoAGUA
IABSAGUAbAB5AGkAbgBnACAAUABhAHIAdAB5ACAAQQBnAHIAZQBlAG0AZQBuAHQA
IAB3AGgAaQBjAGgAIABsAGkAbQBpAHQAIABsAGkAYQBiAGkAbABpAHQAeQAgAGEA
bgBkACAAYQByAGUAIABpAG4AYwBvAHIAcABvAHIAYQB0AGUAZAAgAGgAZQByAGUA
aQBuACAAYgB5ACAAcgBlAGYAZQByAGUAbgBjAGUALjALBglghkgBhv1sAxUwHQYD
VR0OBBYEFO/ONZPO9obF+IT1DOdab9kvS+NkMB8GA1UdIwQYMBaAFEXroq/0ksuC
MS1Ri6enIZ3zbcgPMA0GCSqGSIb3DQEBCwUAA4IBAQBZcM33MOO0+lPLfAGE8hQG
CcL7ZEtlXqJe+QxUvd1mrVyE+N+Pemp3s+gskg9FmXKh9nU4kiUphbJuQWsF3f/6
4Z62HMQoXCF/o+fRoTHmBoqfUouM40z6Ier1rc8f2FWN3NQrsXfngxCpt1zCMTmY
mAp5dyH199lwXcVmAhGblcDLrYk2CsCD6wqBnb3duPDTB1ESi5ttyu5Eswj0kXn+
9eBiEfL3eWn7nZnJx68B506D7b4dzBCAxW82zzVU0YtHi3ZZwAwJyWqxZMst90/p
O6c2IEnk9xfwew8hT0fTjtKQpGltT9IU4sdl0jqMsU+VFB+ccTNJ9lWsvqCwEeqA
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            04:00:00:00:00:01:08:d9:61:1e:1e
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=BE, O=GlobalSign nv-sa, OU=Root CA, CN=GlobalSign Root CA
        Validity
            Not Before: Jan 28 12:00:00 1999 GMT
            Not After : Jan 27 11:00:00 2014 GMT
        Subject: C=BE, O=GlobalSign nv-sa, OU=Primary Secure Server CA, CN=GlobalSign Primary Secure Server CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:f6:6a:ed:a8:6b:30:a3:2d:ac:e9:42:9e:18:35:
                    c0:1e:f7:6f:74:cb:b7:42:24:53:ad:31:cb:ef:a5:
                    c9:c5:3d:03:5e:a5:9d:76:cd:19:e2:e1:16:2d:a4:
                    2d:44:20:f1:1a:1f:f7:7d:60:cd:a6:c7:15:a9:ab:
                    8a:a2:c9:66:6c:dd:10:a3:d8:9b:77:29:ee:a6:40:
                    cd:2f:34:36:7f:a3:17:05:0b:cb:58:a5:22:a6:7c:
                    35:e6:8d:5d:a1:53:c2:9a:c5:da:5d:fe:d8:0e:7d:
                    3b:22:97:52:2c:dd:b2:3c:0b:90:dc:05:fd:b2:e5:
                    0a:55:1e:5d:9e:62:fb:7f:e3:b8:96:f4:9f:26:ac:
                    a2:5c:84:d9:82:ba:e0:e8:f5:95:6e:04:0a:96:64:
                    49:a3:0f:9e:83:a9:63:e7:c9:21:99:6b:a0:16:91:
                    25:c8:14:d9:bd:dc:ec:3c:77:53:47:56:43:84:7e:
                    d6:63:e5:e3:28:af:3c:4f:c0:7d:b4:18:f6:d7:be:
                    57:0b:89:db:d6:c1:83:92:92:e3:9c:30:d1:59:4c:
                    a5:71:90:5f:86:07:70:e8:4e:94:14:c9:f2:4e:a3:
                    80:c2:5a:11:a9:e8:e8:e2:bc:02:9c:bf:38:4d:7a:
                    da:3c:51:63:ee:bc:f8:7c:51:7e:a0:b8:e0:48:a9:
                    af:ad
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier: 
                D1:66:83:F5:89:5E:D0:BE:7F:61:D2:DD:A8:CA:FA:2B:7A:4A:7F:31
            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl.globalsign.net/Root.crl

            X509v3 Authority Key Identifier: 
                keyid:60:7B:66:1A:45:0D:97:CA:89:50:2F:7D:04:CD:34:A8:FF:FC:FD:4B

    Signature Algorithm: sha1WithRSAEncryption
         9c:ff:14:d4:0b:bc:9f:a3:3b:6c:b9:2c:21:c8:37:c3:45:4d:
         30:fb:4b:92:90:e0:58:05:bd:f4:6e:49:d9:94:b5:09:88:68:
         6e:53:dc:ce:2c:76:a3:83:4a:28:c5:8a:3f:45:d0:f1:2f:00:
         65:a0:ae:17:07:38:00:d4:28:e2:c3:e2:49:6c:b7:02:e1:29:
         3d:94:dc:d4:c7:da:63:01:4f:3c:ab:e1:e7:84:61:44:e6:fe:
         2d:7c:ff:e2:5c:19:bc:4c:59:cc:be:9a:b4:71:48:22:ce:2d:
         9b:b4:7e:c1:da:e7:be:8d:67:88:76:65:ea:91:25:03:56:ea:
         04:8e:25:65:68:54:23:19:91:26:df:b2:46:76:31:58:ab:0a:
         fd:a8:35:e2:a0:c4:25:af:7a:30:04:87:1d:ff:c0:7b:6b:80:
         1e:4a:77:7b:d2:2c:14:d2:08:99:06:5c:02:8d:d5:49:a3:5b:
         00:14:aa:33:9c:49:91:6c:2a:01:8b:e2:17:9d:cc:74:c3:6a:
         0a:e6:eb:81:38:a5:fc:27:a9:85:20:96:6d:5c:e0:8b:32:18:
         61:f9:c3:bc:d3:55:e8:86:2c:4b:7d:09:37:f2:f8:d1:91:76:
         6b:bc:fa:90:70:ba:38:5a:24:29:7d:ff:cd:d2:7e:35:6d:d9:
         ee:9a:f8:c5
-----BEGIN CERTIFICATE-----
MIID7zCCAtegAwIBAgILBAAAAAABCNlhHh4wDQYJKoZIhvcNAQEFBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw05OTAxMjgxMjAw
MDBaFw0xNDAxMjcxMTAwMDBaMHkxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMSEwHwYDVQQLExhQcmltYXJ5IFNlY3VyZSBTZXJ2ZXIgQ0Ex
LDAqBgNVBAMTI0dsb2JhbFNpZ24gUHJpbWFyeSBTZWN1cmUgU2VydmVyIENBMIIB
IjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9mrtqGswoy2s6UKeGDXAHvdv
dMu3QiRTrTHL76XJxT0DXqWdds0Z4uEWLaQtRCDxGh/3fWDNpscVqauKoslmbN0Q
o9ibdynupkDNLzQ2f6MXBQvLWKUipnw15o1doVPCmsXaXf7YDn07IpdSLN2yPAuQ
3AX9suUKVR5dnmL7f+O4lvSfJqyiXITZgrrg6PWVbgQKlmRJow+eg6lj58khmWug
FpElyBTZvdzsPHdTR1ZDhH7WY+XjKK88T8B9tBj2175XC4nb1sGDkpLjnDDRWUyl
cZBfhgdw6E6UFMnyTqOAwloRqejo4rwCnL84TXraPFFj7rz4fFF+oLjgSKmvrQID
AQABo4GZMIGWMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8EBTADAQH/MB0GA1Ud
DgQWBBTRZoP1iV7Qvn9h0t2oyvorekp/MTAzBgNVHR8ELDAqMCigJqAkhiJodHRw
Oi8vY3JsLmdsb2JhbHNpZ24ubmV0L1Jvb3QuY3JsMB8GA1UdIwQYMBaAFGB7ZhpF
DZfKiVAvfQTNNKj//P1LMA0GCSqGSIb3DQEBBQUAA4IBAQCc/xTUC7yfoztsuSwh
yDfDRU0w+0uSkOBYBb30bknZlLUJiGhuU9zOLHajg0ooxYo/RdDxLwBloK4XBzgA
1Cjiw+JJbLcC4Sk9lNzUx9pjAU88q+HnhGFE5v4tfP/iXBm8TFnMvpq0cUgizi2b
tH7B2ue+jWeIdmXqkSUDVuoEjiVlaFQjGZEm37JGdjFYqwr9qDXioMQlr3owBIcd
/8B7a4AeSnd70iwU0giZBlwCjdVJo1sAFKoznEmRbCoBi+IXncx0w2oK5uuBOKX8
J6mFIJZtXOCLMhhh+cO801XohixLfQk38vjRkXZrvPqQcLo4WiQpff/N0n41bdnu
mvjF
-----END CERTIFICATE-----

Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            04:00:00:00:00:01:08:d9:61:25:cf
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=BE, O=GlobalSign nv-sa, OU=Primary Secure Server CA, CN=GlobalSign Primary Secure Server CA
        Validity
            Not Before: Jan 22 09:00:00 2004 GMT
            Not After : Jan 27 10:00:00 2014 GMT
        Subject: C=BE, O=GlobalSign nv-sa, OU=ServerSign CA, CN=GlobalSign ServerSign CA
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:d1:d9:0d:a9:e8:bf:c9:6f:fa:24:ee:56:73:b6:
                    13:d0:2e:0d:9d:f1:bf:ed:3c:56:9f:18:c5:8c:c9:
                    1d:8b:66:9d:22:24:98:44:99:98:90:ac:98:ed:52:
                    95:60:08:d2:21:da:4e:7a:56:8e:72:37:e4:5a:22:
                    ca:61:45:5a:f8:1c:e7:8a:4d:04:01:21:70:19:c5:
                    a1:d7:66:20:b1:c2:da:17:06:8f:20:db:0e:d0:46:
                    d4:ef:b0:24:c7:e3:9b:76:f5:c9:c2:d2:04:09:22:
                    fc:b0:65:ff:6c:f6:d6:29:97:63:25:08:ed:90:d6:
                    8a:fc:89:07:bd:ac:eb:03:16:63:7a:ea:a3:f8:82:
                    86:74:54:f7:75:a6:d6:5c:b2:a4:aa:ac:7f:b3:70:
                    bc:4e:25:c7:5b:c6:03:5d:aa:b1:b5:e5:da:f8:2c:
                    d7:b1:32:3d:74:f9:ba:0b:d3:d0:e7:03:1a:64:57:
                    3b:6c:c3:b1:8b:78:e2:03:d0:95:b7:84:9c:6f:f5:
                    51:59:15:81:f3:a5:31:06:90:f9:03:0d:16:83:d8:
                    38:46:b7:c6:de:f2:5b:79:ce:6a:e8:d5:38:80:62:
                    0d:1c:34:09:3a:3d:d9:a1:99:9b:5a:4e:84:56:6d:
                    b6:8c:7d:ec:1d:ae:48:69:6f:38:8c:7e:f4:cb:e9:
                    d7:ed
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Subject Key Identifier: 
                48:BB:A8:BF:5B:84:D2:57:48:E4:61:E9:91:20:91:D8:1D:25:DF:7F
            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl.globalsign.net/primserver.crl

            Netscape Cert Type: 
                SSL CA
            X509v3 Authority Key Identifier: 
                keyid:D1:66:83:F5:89:5E:D0:BE:7F:61:D2:DD:A8:CA:FA:2B:7A:4A:7F:31

    Signature Algorithm: sha1WithRSAEncryption
         3f:f8:95:69:f2:af:5b:4c:ee:a4:63:d0:ff:82:9c:d8:ea:99:
         48:22:e8:97:09:04:d3:eb:6c:67:31:fb:ab:fa:08:2f:9a:35:
         1c:ac:c4:03:86:9c:1c:e1:cc:a3:16:7b:3e:b5:0b:db:8b:19:
         cf:ef:4a:a0:6c:6c:65:ac:f7:fa:98:fb:25:28:a1:fe:4d:24:
         15:a2:e4:27:e8:dc:0e:a1:b3:e3:8d:6e:85:15:91:7c:fd:0d:
         bf:ca:3f:1a:f3:2c:1b:10:30:64:75:1a:dd:c5:b5:ca:d1:0a:
         1d:4b:f3:81:29:ca:aa:bd:28:ac:66:c9:4a:80:1b:e7:da:1e:
         cd:79:c1:e9:85:a4:7a:5f:b8:f3:8d:2d:72:21:3f:0f:0c:18:
         7c:11:4b:ec:aa:ba:48:57:52:33:c8:a8:0e:c3:9f:cf:b6:ec:
         0d:b9:ed:9e:4e:bb:bc:c5:20:06:28:a2:14:5a:10:ca:66:fe:
         fa:93:a9:20:c0:58:57:90:6f:aa:d7:39:bd:41:a5:6b:a7:ce:
         2f:9d:e2:53:81:a2:c0:f9:d4:10:ca:95:c2:9a:da:1c:1b:45:
         3f:e0:c8:a4:d0:ab:d0:86:a2:e4:6d:70:63:9a:df:76:67:73:
         f5:36:89:3e:2d:d8:b7:05:cb:74:d7:d9:01:ff:4a:7c:9b:f9:
         ab:9e:3b:ad
-----BEGIN CERTIFICATE-----
MIIEFzCCAv+gAwIBAgILBAAAAAABCNlhJc8wDQYJKoZIhvcNAQEFBQAweTELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExITAfBgNVBAsTGFBy
aW1hcnkgU2VjdXJlIFNlcnZlciBDQTEsMCoGA1UEAxMjR2xvYmFsU2lnbiBQcmlt
YXJ5IFNlY3VyZSBTZXJ2ZXIgQ0EwHhcNMDQwMTIyMDkwMDAwWhcNMTQwMTI3MTAw
MDAwWjBjMQswCQYDVQQGEwJCRTEZMBcGA1UEChMQR2xvYmFsU2lnbiBudi1zYTEW
MBQGA1UECxMNU2VydmVyU2lnbiBDQTEhMB8GA1UEAxMYR2xvYmFsU2lnbiBTZXJ2
ZXJTaWduIENBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0dkNqei/
yW/6JO5Wc7YT0C4NnfG/7TxWnxjFjMkdi2adIiSYRJmYkKyY7VKVYAjSIdpOelaO
cjfkWiLKYUVa+Bznik0EASFwGcWh12YgscLaFwaPINsO0EbU77Akx+ObdvXJwtIE
CSL8sGX/bPbWKZdjJQjtkNaK/IkHvazrAxZjeuqj+IKGdFT3dabWXLKkqqx/s3C8
TiXHW8YDXaqxteXa+CzXsTI9dPm6C9PQ5wMaZFc7bMOxi3jiA9CVt4Scb/VRWRWB
86UxBpD5Aw0Wg9g4RrfG3vJbec5q6NU4gGINHDQJOj3ZoZmbWk6EVm22jH3sHa5I
aW84jH70y+nX7QIDAQABo4G1MIGyMA4GA1UdDwEB/wQEAwIBBjASBgNVHRMBAf8E
CDAGAQH/AgEAMB0GA1UdDgQWBBRIu6i/W4TSV0jkYemRIJHYHSXffzA5BgNVHR8E
MjAwMC6gLKAqhihodHRwOi8vY3JsLmdsb2JhbHNpZ24ubmV0L3ByaW1zZXJ2ZXIu
Y3JsMBEGCWCGSAGG+EIBAQQEAwICBDAfBgNVHSMEGDAWgBTRZoP1iV7Qvn9h0t2o
yvorekp/MTANBgkqhkiG9w0BAQUFAAOCAQEAP/iVafKvW0zupGPQ/4Kc2OqZSCLo
lwkE0+tsZzH7q/oIL5o1HKzEA4acHOHMoxZ7PrUL24sZz+9KoGxsZaz3+pj7JSih
/k0kFaLkJ+jcDqGz441uhRWRfP0Nv8o/GvMsGxAwZHUa3cW1ytEKHUvzgSnKqr0o
rGbJSoAb59oezXnB6YWkel+4840tciE/DwwYfBFL7Kq6SFdSM8ioDsOfz7bsDbnt
nk67vMUgBiiiFFoQymb++pOpIMBYV5Bvqtc5vUGla6fOL53iU4GiwPnUEMqVwpra
HBtFP+DIpNCr0Iai5G1wY5rfdmdz9TaJPi3YtwXLdNfZAf9KfJv5q547rQ==
-----END CERTIFICATE-----

	`
)
