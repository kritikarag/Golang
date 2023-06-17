package main

import (
	"fmt" 
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
func main(){
	fmt.Println("Welcome to handling URLs in Golang")
	fmt.Println(myurl)
	result, err := url.Parse(myurl)

	if err!=nil{
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	
	fmt.Println(qparams["coursename"])	

	for _,j:=range(qparams){
		fmt.Println("Param is: ",j)
	}

	partsofurl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=kritika",
	}

	anotherURL := partsofurl.String() 
	fmt.Println(anotherURL)
}