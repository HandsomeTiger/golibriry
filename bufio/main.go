package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	//bufio.Reader.
	readString()
}

func readString() {
	sr := strings.NewReader("string hello world!")
	reader := bufio.NewReader(sr)
	bytes, _ := reader.Peek(5)
	fmt.Printf("%s\n", bytes)
	sbyte := make([]byte, 6)
	n, _ := reader.Read(sbyte)
	fmt.Println(n)
	fmt.Println(string(sbyte))
	reader.Discard(1)

	for {
		str, err := reader.ReadString(byte(' '))
		fmt.Println(str)
		if err != nil {
			return
		}
	}
}
