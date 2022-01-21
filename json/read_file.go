package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// {"shop_UID": "3445c227-7efc-11e5-be77-24fd52940c70", "prod_UID": "156ee9ce-a431-11eb-975c-080027048daf", "price": 0.0, "currency": "RUB", "quantity": 0.0}, {"shop_UID": "1d162cd3-886a-11e5-96e1-14dae9ee1802", "prod_UID": "156ee9ce-a431-11eb-975c-080027048daf", "price": 0.0, "currency": "RUB", "quantity": 0.0},

type Prices struct {
	ShopUID  string
	ProdUID  string
	Price    float64
	Currency string
	Quantity float64
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
