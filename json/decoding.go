package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var u1 User

	data := []byte(`{
		"Id": 1,
		"Name": "John",
		"Age": 20
	}`)

	err := json.Unmarshal(data, &u1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("This %s, %d Age", u1.Name, u1.Age)
}
