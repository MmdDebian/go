package server

import (
	"fmt"
	"net/http"
)

func (s server) hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w , "hello this one")
}