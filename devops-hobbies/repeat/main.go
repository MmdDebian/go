package main

import (
	"fmt"
)

type ErrorMessage interface {
	initial() string
}

type response struct {
	message string
	status  int
}

type res struct {
	message string
	status  int
}

func (r response) initial() string {
	return r.message
}

func (r res) initial() string {
	switch r.status {
		case 200:
			return "request is ok"
		case 201 :
			return "created successfully"
		case 500 :
			return "Internal server error"
		default :
			return r.message 
	}
}

func initialError(e ErrorMessage) {
	fmt.Printf("Message : %s \n " , e.initial())
}



type UserService interface  {
	create()
	update()
}


type Users struct {
	name string 
	lastName  string
}


func (u Users) update(){
	fmt.Println(u.lastName)
}


func (u Users) create(){
	fmt.Println(u.name)
}



func main() {
	r1 := res{message: "hello world" , status: 201}
	r2 := res{message: "hello world" , status: 500}
	r3 := res{message: "hello world" , status: 200}
	r4 := res{message: "hello world" , status: 205}

	initialError(r1)
	initialError(r2)
	initialError(r3)
	initialError(r4)


	
	var u UserService 
	u = Users{name: "mamad"}
	u = Users{lastName: "mirzaei"}
	u.update()
	u.create()
}