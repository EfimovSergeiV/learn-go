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
	var user_data []User

	data := []byte(`[
		{"Id": 1, "Name": "John", "Age": 20},
		{"Id": 2, "Name": "John1", "Age": 30},
		{"Id": 3, "Name": "John2", "Age": 40}
	]`)

	err := json.Unmarshal(data, &user_data)

	if err != nil {
		log.Fatal(err)
	}

	for user := range user_data {
		fmt.Println(user_data[user])
	}

	json_data, json_err := json.MarshalIndent(user_data, "", "    ")

	if json_err != nil {
		log.Fatal(json_err)
	}

	fmt.Println(string(json_data))
}

// type Parse struct {
//     NameId string json:name_id"
// }
