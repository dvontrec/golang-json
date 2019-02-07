package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	// Properties have to be uppercase to export as json
	FName string
	Lname string
	Items []string
}

func main() {
	// basic routes
	http.HandleFunc("/", index)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.HandleFunc("/umshl", umshl)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func mshl(w http.ResponseWriter, req *http.Request) {
	// Sets the content type so the page can serve json
	w.Header().Set("Content-Type", "application/json")
	// Creates a new person with the given properties
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Swagger"},
	}
	// marshals the person as a slice of bytes to write to response
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	// prints the object to the response
	w.Write(json)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// creates two people
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Swagger"},
	}
	p2 := person{
		"Dvontre",
		"Coleman",
		[]string{"Heart"},
	}

	// puts created people into a slice of person objects
	people := []person{p1, p2}
	// encodes the people slice into an array of json objects and prints to response
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println(err)
	}
}

func umshl(w http.ResponseWriter, req *http.Request) {
	var p person
	rcvd := `{"FName":"James","Lname":"Bond","Items":["Suit","Gun","Swagger"]}`
	// converts the json from rcvd and stores it in p
	err := json.Unmarshal([]byte(rcvd), &p)
	if err != nil {
		log.Fatalln("Error unmarshalling", err)
	}
	// prints the person to the console
	fmt.Println(p)

}
