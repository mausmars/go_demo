package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Fruit struct {
	Name string `json":name"`
	PriceTag string `json:"priceTag"`
}

type People struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
	Created time.Time
	Fruit   []Fruit
}

func main(){
	jsonData := []byte(`
    {
        "id": 123,
        "name":"test_123",
        "age": 1,
		"Created": "2018-04-09T23:00:00Z",
		"Fruit" : [
		{
			"name": "Apple",
			"priceTag": "$1"
		},
		{
			"name": "Pear",
			"priceTag": "$1.5"
		}
]
    }`)
	var people People
	err := json.Unmarshal(jsonData, &people)
	if(err!=nil){
		fmt.Println("err ",err)
	}else{
		fmt.Println(people)
	}

	data,err:=json.Marshal(people)
	if(err!=nil){
		fmt.Println("err ",err)
	}else{
		fmt.Println(string(data[:]))
	}
}
