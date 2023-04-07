#!/bin/bash

DATE=`date "+%Y_%m_%d"`
VERSION=`git tag | tail -n 1`

all: build-release

run:
	npm run dev

format:
	npm run format

build-release:
	npm run build

download-dependence:
	npm install

clean:
	- rm -rf ./node_modules ./package-lock.json ./dist
