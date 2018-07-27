#!/bin/sh
set -e

go get -d -u github.com/stretchr/testify/assert

cd $GOPATH/src/github.com/go-chassis
rm -rf paas-lager
git clone https://github.com/go-chassis/paas-lager.git

cd $GOPATH/src/github.com/go-chassis/paas-lager
#Start unit test
for d in $(go list ./... | grep -v third_party); do
    echo $d
    echo $GOPATH
    cd $GOPATH/src/$d
    if [ $(ls | grep _test.go | wc -l) -gt 0 ]; then
        go test 
    fi
done
