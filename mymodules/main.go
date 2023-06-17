package main

import(
	"fmt"
	"net/http" 
	"github.com/gorilla/mux"
	"log"
)

func main(){
	fmt.Println("Welcome to mod introduction")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000",r))
}


func greeter()  {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>Welcome to golang series</h1>"))
}