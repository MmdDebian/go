package server

import (
	"fmt"
	"log"
	"net/http"
)


type server struct{}


func New() *server {
	return &server{}
}


func (s server) Serve(port int ){

	http.Handle("/" , &sample{})
	http.HandleFunc("/hello" , s.hello)

	addr := fmt.Sprintf(":%d" , port)
	fmt.Println("app running on port ",addr)
	log.Fatal(http.ListenAndServe(addr , nil))
}


type sample struct {}

func (s *sample) ServeHTTP(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w , "this is the sample handler")
}