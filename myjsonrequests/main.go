package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func main(){
	fmt.Println("Welcome to the web verb video - LCO")
	performGetRequest()
}

func performGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err:= http.Get(myurl)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ",response.StatusCode)
	fmt.Println("Content Length: ",response.ContentLength)

	//content,_:= ioutil.ReadAll(response.Body)

	var responseString strings.Builder
	content,_:= ioutil.ReadAll(response.Body)
	bytecount,_:= responseString.Write(content)

	fmt.Println("Byte Count is: ",bytecount)
	fmt.Println(responseString.String())

	fmt.Println(string(content))
}