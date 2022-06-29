package main

import "fmt"

/*
go get example.com/pkg
- downloads the package into global cache

go get -u example.com/pkg
- gets and upgrades the packages

go get go list -u -m all
- view available dependency upgrades(with indirect)

go get -u ./...
- upgrades to the latest or minor patch release

go get -t -u ./..
- same as above - with test packages

go build [-race]
- build a ninary

go test [-race] -bench

go vet

go tool (pprof) - profiling

go get
- managing the project dependencies

Functions and variables that start with a Capittal letter means that it was exported. 
If it starts with the small letter it can't be exported.

*/

func main() {
	fmt.Println("hello world")
}