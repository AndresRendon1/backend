package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

var superheroes = map[string]Superhero{
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
		}{Intelligence: 90, Strength: 55, Speed: 67, Durability: 75, Power: 74, Combat: 85},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg"},
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
		}{Intelligence: 100, Strength: 85, Speed: 58, Durability: 85, Power: 100, Combat: 64},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg"},
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
		}{Intelligence: 75, Strength: 13, Speed: 33, Durability: 30, Power: 36, Combat: 100},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/170-black-widow.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/170-black-widow.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/170-black-widow.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/170-black-widow.jpg"},
	},
	"Captain America": {
		Name: "Captain America",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Steven Rogers"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{Intelligence: 70, Strength: 80, Speed: 55, Durability: 95, Power: 65, Combat: 100},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/396-captain-america.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/396-captain-america.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/396-captain-america.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/396-captain-america.jpg"},
	},
	"Wolverine": {
		Name: "Wolverine",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "James Howlett"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{Intelligence: 65, Strength: 95, Speed: 50, Durability: 95, Power: 60, Combat: 100},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/253-wolverine.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/253-wolverine.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/253-wolverine.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/253-wolverine.jpg"},
	},
	"Thor": {
		Name: "Thor",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Thor odinson"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{Intelligence: 69, Strength: 100, Speed: 83, Durability: 100, Power: 100, Combat: 100},
		Images: struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		}{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg"},
	},
}

func main() {
	http.HandleFunc("/api/superhero", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		hero := r.URL.Query().Get("hero")
		superhero, ok := superheroes[hero]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		response, err := json.Marshal(superhero)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
