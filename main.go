package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Superhero struct represents the data of a superhero
type Superhero struct {
	Name      string `json:"name"`
	Biography struct {
		FullName string `json:"fullName"`
	} `json:"biography"`
	Powerstats struct {
		Intelligence int `json:"intelligence"`
		Strength     int `json:"strength"`
		Speed        int `json:"speed"`
		Durability   int `json:"durability"`
		Power        int `json:"power"`
		Combat       int `json:"combat"`
	} `json:"powerstats"`
	Images struct {
		Xs string `json:"xs"`
		Sm string `json:"sm"`
		Md string `json:"md"`
		Lg string `json:"lg"`
	} `json:"images"`
}

// Superheroes map contains the data of several superheroes
var Superheroes = map[string]Superhero{
	"Wolverine": {
		Name: "Wolverine",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "John Logan"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 63,
			Strength:     32,
			Speed:        50,
			Durability:   100,
			Power:        89,
			Combat:       100,
		},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg",
		},
	},
	"Spider-Man": {
		Name: "Spider-Man",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Peter Parker"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 90,
			Strength:     55,
			Speed:        67,
			Durability:   75,
			Power:        74,
			Combat:       85,
		},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg",
		},
	},
	"Iron Man": {
		Name: "Iron Man",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Tony Stark"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 100,
			Strength:     85,
			Speed:        58,
			Durability:   85,
			Power:        100,
			Combat:       64,
		},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg",
		},
	},
	"Black Widow": {
		Name: "Black Widow",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Natasha Romanoff"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 75,
			Strength:     13,
			Speed:        33,
			Durability:   30,
			Power:        36,
			Combat:       100,
		},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg",
		},
	},
	"Thor": {
		Name: "Thor",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Thor Odinson"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 69,
			Strength:     100,
			Speed:        83,
			Durability:   100,
			Power:        100,
			Combat:       100,
		},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg",
		},
	},
}

func main() {
	http.HandleFunc("/api/superhero", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		heroName := r.URL.Query().Get("hero")
		superhero, ok := Superheroes[heroName]
		if !ok {
			http.Error(w, "Superhero not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(superhero)
	})

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
