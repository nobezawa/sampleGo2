package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"-"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
	memo    string
}

func main() {

	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "sample@co.jp",
		Age:     5,
		Address: "",
		memo:    "Gopher",
	}
	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}

	//	var person Person
	//	b := []byte(`{"id": 1, "name": "Gopher", "age": 5}`)
	//	err := json.Unmarshal(b, &person)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(person)
	//
	//	file, err := os.Create("./file.txt")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	defer file.Close()
	//
	//	message := []byte("hello world\n")
	//
	//	_, err = file.Write(message)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
}
