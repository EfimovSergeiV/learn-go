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
	u1 := User{1, "John", "Wick", 35}
	json_data, err := json.Marshal(u1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))
}
