#!/bin/bash

[ ! -e data ] && mkdir data;
curl http://curl.haxx.se/ca/cacert.pem > data/ca.pem

# needs go-bindata from github.com/jteeuwen/go-bindata
go get github.com/jteeuwen/go-bindata/go-bindata

go-bindata -pkg=osxcrossfix -nomemcopy -tags="darwin,!cgo" data
