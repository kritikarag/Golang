package main

import (
	"fmt"
)

func main(){
	hello()
	res := sum(1,2)

	res2:= add(1,2,3,4)
	fmt.Println(res)
	fmt.Print(res2)
}

func hello(){ 
	fmt.Println("Hello World!")
}

func sum(val1 int, val2 int) int{
	return val1+val2
}

func add(values...int )int{
	result:=0
	for _,j:=range(values){
		result+= j
	}
	return result
}