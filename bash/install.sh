#!/bin/bash

# Build to support 'replace' in modules since `go install` isn't supported.
# Issue: https://github.com/golang/go/issues/40276
git clone --depth 1 https://github.com/ambientkit/amb.git ambrepo
chmod -R +w ambrepo
cd ambrepo
go build -o ../amb cmd/amb/main.go
cd ..
chmod +x amb
rm -r ambrepo