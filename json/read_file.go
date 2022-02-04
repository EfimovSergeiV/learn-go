package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Prices struct {
	Shop     string  `json:"shop_UID"`
	Prod     string  `json:"prod_UID"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	Quantity float64 `json:"quantity"`
}

func main() {
	filename, err := os.Open("upload-data.json")
	if err != nil {
		log.Fatal(err)
	}

	defer filename.Close()

	data, err := ioutil.ReadAll(filename)
	if err != nil {
		log.Fatal(err)
	}

	var result []Prices

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for price := range result {
		fmt.Println(result[price])
	}
}
