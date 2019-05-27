package main

import (
	"bytes"
	"fmt"
)

func main() {
	//exampleBase()
	exampleCompare()
}

func exampleBase() {
	str := "hello"
	byte1 := []byte(str)
	str1 := "hl"
	byte2 := []byte(str1)
	// Contains 判断a中是否包含有b
	bo := bytes.Contains(byte1, byte2)
	fmt.Printf("%+v\n", bo)
	// Count计算s中有多少个不重叠的sep子切片。
	c := bytes.Count(byte1, []byte{'h', 'o'})
	fmt.Println(c)

	r := bytes.Repeat(byte2, 5)
	fmt.Printf("%s\n", r)

}

func exampleCompare() {

	a := []byte("hello")
	c := "呵ba哈"
	d := []byte(c)
	d = append(d, 97)
	println(d)
	println(string(c))
	b := []byte("haaaaas")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(bytes.Compare(a, b))

}

type test1 struct {
	Phone string
	Age   int
	Name  string
}
