package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//test1()
	//test2()
	test3()
}

func test() {
	var ip = flag.Int("flagname", 1234, "help message for flagname")
	var ip2 = flag.Int("flagnames", 111, "help message for flagname")
	flag.Parse()
	flag.Usage()
	fmt.Println(*ip)
	fmt.Println(*ip2)
	fmt.Println(flag.Args())
}

func test1() {
	fs := flag.NewFlagSet("flag", flag.ContinueOnError)
	ip := fs.Int("ints", 122, "help message")
	flag.Parse()
	fmt.Println(*ip)
}

func test2() {
	var mothed int
	flag.IntVar(&mothed, "m", 123, "message")
	flag.Parse()
	fmt.Println(mothed)
}

func test3() {
	args := os.Args
	fmt.Println(args[0])

}
