package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID 			string 		`json:"id"`
	Isbn 		string 		`json:"isbn"`
	Title 		string 		`json:"title"`
	Director 	*Director 	`json:"director"`
}

type Director struct {
	Firstname 	string 		`json:"firstname"`
	Lastname 	string 		`json:"lastname"`
}

// central movies slice
var movies []Movie

// get all movies function | GET
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// delete the movie having particular ID | DELETE
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// get a particular movie using ID | GET
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// create a particular movie | POST
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// update a particular movie | PUT
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// content type = json
	w.Header().Set("Content-Type", "application/json")

	// getting params
	params := mux.Vars(r)

	// iterate/range over the movies
	// delete that same movie with that particular ID 
	// add a new movie that is sent through the body
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

// Main function with all the major functionality callings
func main() {
	route := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "0036924", Title: "First Movie", Director : &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "1036924", Title: "Second Movie", Director : &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "3", Isbn: "3036924", Title: "Third Movie", Director : &Director{Firstname: "Alex", Lastname: "Wattney"}})

	route.HandleFunc("/movies", getMovies).Methods("GET")
	route.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	route.HandleFunc("/movies", createMovie).Methods("POST")
	route.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	route.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// TEST URL = http://localhost:6000/movies
	fmt.Printf("Starting server at port 6000\n")
	log.Fatal(http.ListenAndServe(":6000", route))
}
