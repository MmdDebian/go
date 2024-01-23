package main

import "fmt"

type person struct {
	id         int
	first_name string
	last_name  string
}

func addPerson(name string) *person {
	p := person{first_name: name}
	p.last_name = "mirzaei"

	return &p
}

func main() {
	fmt.Println(person{1 , "mamad" , "mirzaei"})
	
	fmt.Println(person{2 , "milad" , "mirzaei"})

	fmt.Println(person{first_name: "mamad"})

	fmt.Println(&person{first_name:"milad" , last_name: "mirzaei"})

	fmt.Println(addPerson("ali"))

	s := person{first_name: "milad" , id: 50}

	fmt.Println(s.first_name)

	sp := &s 
	
	fmt.Println(sp)


	sp.id = 51 
	fmt.Println(sp.id)


	dog := struct {
		age int 
		name string 
	} {
		age : 15 ,
		name : "rex" ,
	}

	fmt.Println(dog)
}