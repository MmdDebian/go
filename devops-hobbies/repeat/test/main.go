package main

import "github.com/mohamadmirzaeidev/test1/server"

func main() {
	server := server.New()
	server.Serve(3000)
}