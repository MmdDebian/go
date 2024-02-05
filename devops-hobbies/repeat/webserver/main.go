package main

import "server.test/webserver/server"

func main() {
	server := server.New()
	server.Serve(4000)
}