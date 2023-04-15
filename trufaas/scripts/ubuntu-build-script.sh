#!/bin/bash

make skaffold-prebuild
skaffold run -p kind
go build -o fission.exe cmd/fission-cli/main.go
GOOS=linux GOARCH=amd64 go build -o fission cmd/fission-cli/main.go
sudo mv ./fission /usr/local/bin/fission