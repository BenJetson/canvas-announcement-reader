#!/usr/bin/env bash

TARGET="./cmd/reader/main.go"

GOOS=darwin GOARCH=amd64 go build -o reader-darwin-amd64 $TARGET
GOOS=windows GOARCH=amd64 go build -o reader-windows-amd64.exe $TARGET
GOOS=linux GOARCH=amd64 go build -o reader-linux-amd64 $TARGET
