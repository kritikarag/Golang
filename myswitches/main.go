package main

import (
	"fmt"
	"math/rand"
)

func main(){
	dicenum := rand.Intn(6)+1
	fmt.Println("Dice Number :",dicenum)
	switch dicenum{
	case 1: 
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6: 
		fmt.Println("Six")
	default:
		fmt.Println("Invalid")
	}
}