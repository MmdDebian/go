package json 

import (
	"encoding/json"
)

type Person struct {
	Name string 
	Age int 
}


func createJson() ([]byte , error){
	person := Person{Name:"Alice" , Age : 30}


	jsonByte ,err := json.Marshal(person)

	if err != nil {
		return nil , err
	}

	return jsonByte , nil 
}