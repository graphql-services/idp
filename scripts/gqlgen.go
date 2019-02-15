package main

// go run scripts/gqlgen.go init
// go run scripts/gqlgen.go -v
// go run server/server.go
// go generate ./...
// go build -o binary ./server/server.go

import "github.com/99designs/gqlgen/cmd"

func main() {
	cmd.Execute()
}
