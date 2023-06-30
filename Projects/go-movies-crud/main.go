package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FName string `json:"fname"`
	LName string `json:"lname"`
}

var movies[]Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)

	for _,item:=range movies{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)

	for index, item:=range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]... )
			break
		}
	}

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func deleteMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)

	for index, item:=range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]... )
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r:=mux.NewRouter()

	movies = append(movies, Movie{ID:"1",ISBN:"764754",Title:"ABC",Director: &Director{FName: "Jonh",LName:"Walker"}})
	movies = append(movies, Movie{ID:"2",ISBN:"764751",Title:"EFG",Director: &Director{FName: "Taylor",LName:"Swift"}})
	movies = append(movies, Movie{ID:"3",ISBN:"764753",Title:"HIJ",Director: &Director{FName: "Charlie",LName:"Puth"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies",updateMovies).Methods("UPDATE")
	r.HandleFunc("/movies",deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8082")

	log.Fatal(http.ListenAndServe(":8082", r))
}