package main

import (
	"fmt"
	"regexp"
)

func main() {
	example1()
	example2()
}

func example1() {
	str := "/html/body/div[3]/div/div[3]"
	pattern := `/(\w+?)/\w+`
	// Match 方法返回bool 返回是否匹配到
	reg, err := regexp.MatchString(pattern, str)
	if err != nil {
		return
	}
	fmt.Println(reg)

}

func example2() {
	str := "/html/body/div[3]/div/div[3]"
	pattern := `/(\w+?)/(\w+)`
	// Match 方法返回bool 返回是否匹配到
	reg := regexp.MustCompile(pattern)
	fmt.Println(reg.String())
	//reg.Longest()
	s := reg.FindStringSubmatch(str)
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[1])

	g := reg.FindAllStringSubmatch(str, 100)
	fmt.Println(g)
	fmt.Println(g[0])
	fmt.Println(g[0][0])
	fmt.Println(g[0][1])
	fmt.Println(g[1])
	fmt.Println(g[1][0])
	fmt.Println(g[1][1])
	//fmt.Println(g[2])
	//fmt.Println(g[2][0])
	//fmt.Println(g[2][1])
	pattern2 := `/.*`
	reg2 := regexp.MustCompile(pattern2)
	// Longest 最左边最长 First 最左边第一个
	reg2.Longest()

	s2 := reg2.FindAllString(str, 100)
	fmt.Println(s2)
	fmt.Println(s2[0])

}
