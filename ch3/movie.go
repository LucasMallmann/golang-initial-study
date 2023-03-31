package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
	Age    int
}

var movies = []Movie{
	{
		Title:  "The Shawshank Redemption",
		Year:   1994,
		Color:  true,
		Actors: []string{"Tim Robbins", "Morgan Freeman"},
		Age:    10,
	},
	{
		Title:  "The Godfather",
		Year:   1972,
		Color:  false,
		Actors: []string{"Marlon Brando", "Al Pacino", "James Caan"},
		Age:    12,
	},
	{
		Title:  "The Dark Knight",
		Year:   2008,
		Color:  true,
		Actors: []string{"Christian Bale", "Heath Ledger", "Aaron Eckhart"},
		Age:    13,
	},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var years []struct {
		Age int
	}
	if err := json.Unmarshal(data, &years); err != nil {
		log.Fatalf("Json Unmarshaling: %s", err)
	}
	fmt.Println(years)
}
