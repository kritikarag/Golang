package main

import (
	"fmt"
)

func main(){
	 
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
	mydefer()
}

func mydefer(){
	for i:= 0; i<5 ;i++{
		defer fmt.Print(i)
	}
}