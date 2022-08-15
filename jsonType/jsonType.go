package jsonType

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

func TestJson() {
	fmt.Println("============= json ===============")

	movie := Movie{"六号车厢", 2021, 100, []string{"ky", "kni"}}

	// struct -> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s \n", jsonStr)

	// json -> struct
	movie2 := Movie{}
	err = json.Unmarshal(jsonStr, &movie2)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf(" %v\n", movie2)
}
