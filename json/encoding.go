package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id       int
	Name     string
	LastName string
	Age      int
}

func main() {
	users := []User{
		{Id: 1, Name: "Roger", LastName: "Roe", Age: 25},
		{Id: 1, Name: "Roger", LastName: "Roe", Age: 25},
		{Id: 1, Name: "Roger", LastName: "Roe", Age: 25},
	}

	json_data, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))
}
