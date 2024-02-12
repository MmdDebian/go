package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s Server) Serve(port int) {
	addr := fmt.Sprintf(":%d", port)
	fmt.Println("server running on port " , addr)
	log.Fatal(http.ListenAndServe(addr , nil))
}