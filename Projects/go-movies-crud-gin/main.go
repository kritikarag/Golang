package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FName string `json:"fname"`
	LName string `json:"lname"`
}

var movies []Movie

func getMovies(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, movies)
}

func getMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")

	for _, item := range movies {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func createMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	c.JSON(http.StatusOK, movie)
}

func updateMovies(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")

	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	c.JSON(http.StatusOK, movie)
}

func deleteMovies(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")

	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	c.JSON(http.StatusOK, movies)
}

func main() {
	r := gin.Default()

	movies = append(movies, Movie{ID: "1", ISBN: "764754", Title: "ABC", Director: &Director{FName: "Jonh", LName: "Walker"}})
	movies = append(movies, Movie{ID: "2", ISBN: "764751", Title: "EFG", Director: &Director{FName: "Taylor", LName: "Swift"}})
	movies = append(movies, Movie{ID: "3", ISBN: "764753", Title: "HIJ", Director: &Director{FName: "Charlie", LName: "Puth"}})

	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMovie)
	r.POST("/movies", createMovie)
	r.PUT("/movies/:id", updateMovies)
	r.DELETE("/movies/:id", deleteMovies)

	fmt.Println("Starting server at port 8083")

	if err := r.Run(":8083"); err != nil {
		log.Fatal(err)
	}
}
