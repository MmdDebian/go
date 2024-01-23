package main

import (
	"fmt"
	"math"
)

// package Redhat
type DNF struct {
	version string 
	num int 
}


func (dnf *DNF) update(pkg string ){
	dnf.num = 2 
	fmt.Printf("%p \n" , &dnf)
	fmt.Printf("updating package %s\n" , pkg)
}


func runStruct(){
	fedora := DNF{version: "1.1.1" , num: 1} 
	fmt.Printf("%p \n" , &fedora)

	fedora.update("curl")
	fmt.Println(fedora.num)
}


type geometry interface {
	area() float64
}


type rectangle struct {
	with , height float64
}


func (r rectangle) area() float64 {
	return r.height * r.with
}

type circle struct {
	radius float64 
}


func (c circle) area()float64 {
	return math.Pi * c.radius * c.radius
}


func printArea(g geometry){
	fmt.Printf("The geomtry area is: %f\n" , g.area())
}


func printResult(){
	r1 := &rectangle{with: 5 , height: 3}
	c1 := &circle{radius: 4}

	printArea(r1)
	printArea(c1)
}



type animal struct {
	age int 
}

func (a animal) walk(){
	fmt.Println("I can walk")
}


type person struct {
	animal 
	name string 
}

func (p person) talk(){
	fmt.Println("i can talk to you")
}


func main(){
	mohamad := person{
		animal: animal{
			age: 23,
		},
		name : "mohamad" ,
	}

	mohamad.talk()
	mohamad.walk()
}