package main

import (
	"fmt"
)

func main(){
	fmt.Println("Structures in Golang")
	student := user{"Kritika","kritika@gmail.com",true,21}

	fmt.Println(student)
	fmt.Printf("%+v\n",student)
}

type user struct{
	Name string
	Email string
	Status bool
	Age int
}