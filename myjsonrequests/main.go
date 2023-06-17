package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"net/url"
	 
)

func main(){
	fmt.Println("Welcome to the web verb video - LCO")
	//performGetRequest()
	//performPostRequest()
	PerformPostFormRequest()
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

func performPostRequest(){
	const myurl = "http://localhost:8000/post"

	requestBody:= strings.NewReader(`
		{
			"name":"Kritika Rag",
			"age":21,
			"city":"Delhi"
		}
	`)
	response ,err := http.Post(myurl,"application/json",requestBody)
		
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname","Kritika")
	data.Add("lastname","Rag")
	data.Add("email","kritika@gmail.com")

	response, err:= http.PostForm(myurl,data)
	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()
	content,_:= ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}