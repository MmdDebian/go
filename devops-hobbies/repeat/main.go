package main

import "fmt"

type Auth interface {
	Register(username string) string
	Login(username string) res
}

type response struct {
	message string
	status  int
}

type res struct {
	response
	id int
}

func (r response) Register(username string) string {
	return r.message
}

func (r res) Login(username string) res {
	return res{
		response: response{
			message: r.message,
			status: r.status,
		},
		id: r.id,
	}
}

func printAuth(a Auth) {
	l := a.Login("milad")

	fmt.Printf("the id of %d \n" , l.id)
	fmt.Printf("the message of %s \n" , l.message)
	fmt.Printf("the status of %d \n" , l.status)
}


func main(){
	// r1 := response{message: "hello" , status: 200}
	r2 := res{response: response{message: "hello" , status: 200} , id: 11}

	// printAuth(r1)
	printAuth(r2)
}