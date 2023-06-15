package main

import (
	"fmt"
	"sort"
)

func main(){
	// var fruit =  []string{}
	// var fruit = []string{"apple","peach","guava","mango","banana","litchi"}
	// fruit = append(fruit, "banana")
	// fruit = append(fruit[1:])
	// fmt.Println(fruit)
	// fmt.Println(len(fruit))
	// fmt.Println(fruit[2])

	score := make([]int, 3)

	score[0] = 78
	score[1] = 76
	score[2] = 45

	score = append(score, 98, 67, 90)

	score = append(score[:2],score[3:]...)
	fmt.Println(score)
	sort.Ints(score)
	fmt.Println(score)

}