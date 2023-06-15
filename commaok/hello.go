package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    fmt.Println("Hello, world!");
	var name string= "Kritika"
	fmt.Println(name);
	lastname:= "Rag"
	fmt.Print(lastname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")

		//comma ok syntax
	input, _:= reader.ReadString('\n')
	fmt.Println("Here is the rating", input)

	num , err := strconv.ParseInt(strings.TrimSpace(input),0,64)
	if err!= nil{
		fmt.Print(err)
	}else{
		fmt.Print(num+1)
	}
}