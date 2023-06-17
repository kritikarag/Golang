package main

import (
	"fmt"
)

func main(){
	fmt.Println("Structures in Golang")
	student := user{"Kritika","kritika@gmail.com",true,21}

	fmt.Println(student)
	fmt.Printf("%+v\n",student)
	student.GetStatus()
	student.NewMail()
}

type user struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u user) GetStatus(){
	fmt.Println("Is user active: ",u.Status)
}

func (u user) NewMail(){
	u.Email= "kritikarag.kr@gmail.com"
	fmt.Println("Email of this user is: ",u.Email)
}