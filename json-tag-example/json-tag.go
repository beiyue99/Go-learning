package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "xiaozhu"}}
	//结构体->json
	//	jsonStr={"title":"喜剧之王","year":2000,"price":10,"actors":["xingye","xiaozhu"]}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json Marshal error:", err)
		return
	}
	fmt.Printf("jsonStr=%s\n", jsonStr)

	//json->结构体
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json.Unmarshal failed,err=", err)
		return
	}
	fmt.Println(my_movie)
}
