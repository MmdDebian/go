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
	addr := fmt.Sprintf(": %d",port)
	log.Fatal(http.ListenAndServe(addr , nil))
	fmt.Printf("server running on port %d" , addr)
}


func (s Server) test(){
	
}