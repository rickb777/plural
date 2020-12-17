#!/bin/bash -e
cd "$(dirname $0)"
PATH=$HOME/go/bin:$PATH
unset GOPATH
export GO111MODULE=on

if ! type -p goveralls; then
  echo go install github.com/mattn/goveralls
  go getinstallgithub.com/mattn/goveralls
fi

echo date...
go test -v -covermode=count -coverprofile=date.out .
go tool cover -func=date.out
[ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=date.out -service=travis-ci -repotoken $COVERALLS_TOKEN
