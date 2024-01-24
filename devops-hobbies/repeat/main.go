package main 


type HttpException interface {
	OK() string 
	BADREQUEST() string 
}

