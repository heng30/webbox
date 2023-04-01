#!/bin/bash

APP_NAME=webbox
DATE=`date "+%Y_%m_%d"`
VERSION=`git tag | tail -n 1`
BUILD_PLATFORM=`uname -s`_`uname -m`
PASSWORD=`echo -n "123456" | md5sum | awk '{print $$1}'`
Auth_Token="Auth-Token: your-token"

all: build-release

run: main.go version.go
	go run $^

build:
	rm -f ./webbox_*
	echo "package main" > version.go
	echo "const VERSION string = \"${VERSION}\"" >> version.go
	go build -o ${APP_NAME}_${VERSION}_${BUILD_PLATFORM}_${DATE}

build-release:
	rm -f ./webbox_*
	echo "package main" > version.go
	echo "const VERSION string = \"${VERSION}\"" >> version.go
	go build -ldflags "-s -w" -o ${APP_NAME}_${VERSION}_${BUILD_PLATFORM}_${DATE}

build-arm:
	echo "package main" > version.go
	echo "const VERSION string = \"${VERSION}\"" >> version.go
	env GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ${APP_NAME}_arm_${VERSION}_${DATE}

gen_denpendence:
	go mod tidy

download-dependence:
	go mod download

import-dependence:
	go mod vendor

gen-cert:
	cd cert && mkcert -install && mkcert localhost 127.0.0.1 ::1&& mv localhost+2-key.pem key.pem && mv localhost+2.pem cert.pem

ping:
	curl "http://localhost:8002/ping"

ping-tls:
	curl "https://localhost:8002/ping"

login:
	curl -v "http://localhost:8002/login?username=test&&password=${PASSWORD}"

catalog:
	curl -v -H ${Auth_Token} "http://localhost:8002/catalog?path=/"

download:
	curl -v -H ${Auth_Token} "http://localhost:8002/download?filename=/a.c" -o /tmp/a.c

upload:
	curl -v -H ${Auth_Token} -X POST -F 'file=@main.go' "http://localhost:8002/upload?filename=/main.go"

mkdir:
	curl -v -H ${Auth_Token} "http://localhost:8002/mkdir?dirname=/foo/bar"

deldir:
	curl -v -H ${Auth_Token} "http://localhost:8002/deldir?dirname=/foo/bar"

delfile:
	curl -v -H ${Auth_Token} "http://localhost:8002/delfile?filename=/main.go"

