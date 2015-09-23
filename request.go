package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	User_id int    `json:"userId"`
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Comp    bool   `json:"completed"`
}

func main() {

	url := "http://jsonplaceholder.typicode.com/todos"
	res, err := http.Get(url)
	if err != nil {
		err.Error()
	}
	var user []User

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&user)

	//fmt.Println(user[0].Id)
	for _, value := range user {
		fmt.Println(value)
	}
}
