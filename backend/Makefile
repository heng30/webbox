#!/bin/bash

APP_NAME=webbox
DATE=`date "+%Y_%m_%d"`
VERSION=`git tag | tail -n 1`
BUILD_PLATFORM=`uname -s`_`uname -m`
PASSWORD=`echo -n "123456" | md5sum | awk '{print $$1}'`
AuthToken="your-token"

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
	curl -v "http://localhost:8002/catalog?authtoken=${AuthToken}&&path=/"

download:
	curl -v "http://localhost:8002/download?authtoken=${AuthToken}&&filename=/a.c" -o /tmp/a.c

upload:
	curl -v "http://localhost:8002/upload?authtoken=${AuthToken}&&filename=/main.go"

readdir:
	curl -v "http://localhost:8002/readdir?authtoken=${AuthToken}&&dirname=/"

mkdir:
	curl -v "http://localhost:8002/mkdir?authtoken=${AuthToken}&&dirname=/foo/bar"

deldir:
	curl -v "http://localhost:8002/deldir?authtoken=${AuthToken}&&dirname=/foo/bar"

delfile:
	curl -v "http://localhost:8002/delfile?authtoken=${AuthToken}&&filename=/main.go"

configure:
	curl -v "http://localhost:8002/configure?authtoken=${AuthToken}"


