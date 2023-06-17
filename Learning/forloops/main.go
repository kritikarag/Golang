package main

import (
	"fmt"
)

func main(){
	var fruit = [5]string{"apple","peach","guava","mango","banana"}

	// for i:=0 ; i<len(fruit);i++{
	// 	fmt.Println(fruit[i])
	// }

	// for i:=range(fruit){
	// 	fmt.Println(fruit[i])
	// }

	for i,j:=range(fruit){
		fmt.Printf("%v: %v ",i,j)
	}
}