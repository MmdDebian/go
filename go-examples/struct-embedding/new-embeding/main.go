package main

import (
	"fmt"
	"time"
)

type detail struct {
	ram int64 
	rom int64 
	storage int64
}


func (d *detail) getTime() time.Time {
	return time.Now()
}

func (d *detail) getInfo() {
	fmt.Printf("Rom is %d and Ram is %d [storage is %d] \n" , d.ram , d.rom , d.storage)
}

type os struct {
	detail 
	name string 
	id string 
}

func (o *os) printOsName() {
	fmt.Printf("The you os name is : %s \n" , o.name)
}

func main() {
	newOs := os{
		detail: detail{
			ram:  50 ,
			rom: 20,
			storage: 500,
		},
		name: "kali linux",
		id: "da545aed4ae",
	}	

	newOs.getTime()
	newOs.getInfo()
	newOs.printOsName()
}
