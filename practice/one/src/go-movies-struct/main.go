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

type movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lirstname"`
}

var movies []movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovieById(id string) {
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteMovieById(params["id"])
	json.NewEncoder(w).Encode(movies)
}

func getMoviesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var moviee movie

	_ = json.NewDecoder(r.Body).Decode(&moviee)
	moviee.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, moviee)
	json.NewEncoder(w).Encode(moviee)
}

func updateMoveis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteMovieById(params["id"])

	var moviee movie

	_ = json.NewDecoder(r.Body).Decode(&moviee)
	moviee.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, moviee)
	json.NewEncoder(w).Encode(movies)

}

func main() {
	fmt.Println("hello")

	r := mux.NewRouter()

	movies = append(movies, movie{ID: "1", Isbn: "43827", Title: "Movie One", Director: &Director{Firstname: "Mohan", Lastname: "Raj"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMoviesById).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMoveis).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("start server 8000")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
