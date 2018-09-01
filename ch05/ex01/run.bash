#!/bin/bash

go build gopl.io/ch1/fetch
go build findlinks.go
./fetch https://golang.org | ./findlinks