package main

import (
	"fmt"
)

func main(){
	var ptr *int 
	fmt.Println("Value of pointer is: ",ptr)

	num:= 123
	ptr = &num 

	fmt.Println("Value of pointer is: ",ptr)
	fmt.Println("Value of pointer is: ",*ptr)
	
	*ptr = *ptr + 2

	fmt.Println("Value of pointer is: ",num)
}