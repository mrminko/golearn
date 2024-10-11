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
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Encoding failed, %s", err)
	}
	//fmt.Println(data)
	fmt.Printf("%s\n", data)
	data, err = json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("Encoding failed, %s", err)
	}
	fmt.Printf("%s\n", data)

	var Titles []struct{ Title string }
	if err := json.Unmarshal(data, &Titles); err != nil {
		log.Fatalf("Error when decoding, %s", err)
	}
	fmt.Println(Titles)
}
