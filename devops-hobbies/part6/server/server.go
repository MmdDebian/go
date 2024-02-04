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

func (s *server) Server(port int) {

	http.HandleFunc("/hello" , s.hello)
	http.HandleFunc("/bmi" , s.bmi)
	http.Handle("/sample" , &sample{})

	addr := fmt.Sprintf(":%d",port)
	fmt.Println("Listening on" , addr)
	log.Fatal(http.ListenAndServe(addr,nil))
}


type sample struct {}


func (s *sample) ServeHTTP(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w , "sample handler")
}