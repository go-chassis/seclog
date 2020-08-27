#!/bin/sh
set -e
#Start unit test
for d in $(go list ./... | grep -v third_party| grep -v example); do
    if [ $(ls | grep _test.go | wc -l) -gt 0 ]; then
        go test $d -cover
    fi
done
