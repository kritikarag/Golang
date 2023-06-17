package main

import (
	"fmt"
)

func main(){
	fmt.Println("Maps in Golang")
	lang := make(map[int]string)

	lang[1] = "KR"
	lang[2] = "MG"
	lang[3] = "TP"

	fmt.Println(lang)
	fmt.Println(len(lang))

	delete(lang,2)

	fmt.Println(lang)
	fmt.Println(len(lang))

	for key,value:=range lang{
		fmt.Printf("%v : %v ",key,value)
	}
}