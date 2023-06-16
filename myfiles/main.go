package main

import (
	"fmt" 
	"os"
	"io"
	"io/ioutil"
)

func main(){
	fmt.Println("Welcome to files in Golang!")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err!=nil{
		panic(err)
	}
	length, err := io.WriteString(file, content)

	if err!=nil{
		panic(err)
	}

	fmt.Println("Length is: ",length)

	defer file.Close()
	readfile("./mylcogofile.txt")
}

func readfile(filename string){
	databyte, err := ioutil.ReadFile(filename)

	if err!=nil{
		panic(err)
	}
	fmt.Println("Text data inside the file is\n", string(databyte))

}