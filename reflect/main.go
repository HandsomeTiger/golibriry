package main

import (
	"fmt"
	"reflect"
)

type r struct {
	ID int
}

func main() {
	a := r{ID: 1}
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Println(t)
	fmt.Println(v)
	fmt.Println(t.Field(0).Type)
	fmt.Println(t.Kind())

}
