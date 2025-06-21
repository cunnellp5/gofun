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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

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

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(88888888))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(
		movies,
		Movie{ID: "1", Isbn: "42069", Title: "Boots", Director: &Director{Firstname: "jay", Lastname: "silentbob"}},
		Movie{ID: "2", Isbn: "42070", Title: "and cats", Director: &Director{Firstname: "spboobs", Lastname: "frankklin"}},
		Movie{ID: "3", Isbn: "42071", Title: "Dart worm", Director: &Director{Firstname: "missy", Lastname: "elliot"}},
		Movie{ID: "4", Isbn: "42072", Title: "Dr chaos", Director: &Director{Firstname: "steezy", Lastname: "g"}},
		Movie{ID: "5", Isbn: "42073", Title: "professor galaxy", Director: &Director{Firstname: "home", Lastname: "sampson"}},
		Movie{ID: "6", Isbn: "42074", Title: "birdo", Director: &Director{Firstname: "rillo", Lastname: "debrief"}},
		Movie{ID: "7", Isbn: "42075", Title: "trash cookie", Director: &Director{Firstname: "eddgs", Lastname: "salasd"}},
		Movie{ID: "8", Isbn: "42076", Title: "lets", Director: &Director{Firstname: "mary", Lastname: "jane"}},
		Movie{ID: "9", Isbn: "42077", Title: "legs", Director: &Director{Firstname: "grays", Lastname: "anatomy"}},
		Movie{ID: "10", Isbn: "42078", Title: "wine", Director: &Director{Firstname: "jay", Lastname: "silentbob"}},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server on port 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
