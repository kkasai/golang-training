#!/bin/bash

go build gopl.io/ch1/fetch
go build element_count.go
./fetch https://golang.org | ./element_count