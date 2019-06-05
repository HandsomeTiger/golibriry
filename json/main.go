package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
	sex  string
}

func main() {
	structEncode()
	tagEncode()
	writerEncode()
}

func structEncode() {
	bauer := Person{"Bauer", 25, "Man"}
	bytes, err := json.Marshal(bauer)
	if err != nil {
		fmt.Println("Marshal failed.")
	}
	fmt.Println(string(bytes)) // {"Name":"Bauer","Age":25}
}

type Person2 struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age"`
	Sex  string `json:"-"`
}

func tagEncode() {
	bauer := Person2{"", 25, "Man"}
	bytes, err := json.Marshal(bauer)
	if err != nil {
		fmt.Println("Marshal failed.")
	}
	fmt.Println(string(bytes)) // {"name":"Bauer","age":25}
}

type Person3 struct {
	Name string
	Age  int
}

func writerEncode() {
	w := os.Stdout
	persons := []Person3{
		{"Bauer", 30},
		{"Bob", 20},
		{"Lee", 24},
	}

	encoder := json.NewEncoder(w)
	for _, person := range persons {
		encoder.Encode(person)
	}
}
