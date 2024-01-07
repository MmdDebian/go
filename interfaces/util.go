package main 


type calc interface {
	add() float64 
	multiplication() float64
}



type input struct {
	oneNumber , twoNumber float64 
}


func add(i input) float64 {
	return i.oneNumber + i.twoNumber 
}


func multiplication(i input) float64 {
	return i.oneNumber * i.twoNumber 
}

