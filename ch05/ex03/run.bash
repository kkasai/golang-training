#!/bin/bash

go build gopl.io/ch1/fetch
go build text_node.go
./fetch https://golang.org | ./text_node