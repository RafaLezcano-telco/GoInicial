#!/bin/sh  

for d in $(go list ./...); do
    echo "Testing $d"
    go test -v $d
done


