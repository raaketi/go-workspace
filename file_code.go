package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	Name       string
	Occupation string
	Born       string
	Fav_food   []string
}

type fav_foods interface {
	// Methods
	Get_fav_food() string
}

func (U User) Get_fav_food() string {
	// fmt.Println(U)
	return U.Fav_food[0] + "," + U.Fav_food[1]
}

func fetch_fav_food(My_fav fav_foods) {
	fmt.Println(My_fav.Get_fav_food())

}

func main() {
	filename, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(filename)
	defer filename.Close()
	if err != nil {
		log.Fatal(err)
	}

	var p []User
	jsonErr := json.Unmarshal([]byte(data), &p)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for _, b := range p {
		// fmt.Println(b)
		// fmt.Println(b.type())
		fetch_fav_food(&b)
	}
}
