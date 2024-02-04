package main

import "github.com/mohamadmirzaeidev/go-cource/server"

func main() {
	server := server.New()

	server.Server(8080)
}