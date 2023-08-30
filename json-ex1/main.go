package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// https://yourbasic.org/golang/json-example/

func fromJsonNoType() {
	jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var val interface{}
	json.Unmarshal(jsonData, &val)
	data := val.(map[string]interface{})

	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
}

type FruitBasket struct {
	Name    string
	Fruit   []string
	Id      int64  `json:"ref"`
	private string // An unexported field is not encoded.
	Created time.Time
}

func toJsonString() {

	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		Id:      999,
		private: "Second-rate",
		Created: time.Now(),
	}

	var jsonData []byte
	// jsonData, err := json.Marshal(basket)
	jsonData, err := json.MarshalIndent(basket, "", "    ")

	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
}

func fromJsonString() {
	jsonData := []byte(`
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket)
	fmt.Println(basket.Created)
}

func main() {
	// toJsonString()
	// fromJsonString()
	fromJsonNoType()
}
