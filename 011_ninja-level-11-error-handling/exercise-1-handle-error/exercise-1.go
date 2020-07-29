package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		//log.Println(err) // Remember to consider the programs purpose. The entire purpose of this program is to marshal p1, if that fails program should exit. Can use log.Fatal
		log.Fatalln("JSON failed to marshal: ", err)
	}
	fmt.Println(string(bs))

}
