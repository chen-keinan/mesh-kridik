#!/usr/bin/env bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0
export PATH=$GOPATH/bin:$PATH
export PATH=$PATH:/root/go/bin
golangci-lint run -v  > lint.xml