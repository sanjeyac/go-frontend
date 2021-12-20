#!/bin/sh

OUTPUT=./build/main.wasm

GOOS=js GOARCH=wasm go build -o $OUTPUT ./src/main.go