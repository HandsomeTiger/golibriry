package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//example1()

	//example2()
	//example3()
	TestTimer()
}
func example1() {
	now, _ := time.Parse(time.ANSIC, "Mon Jan 2 15:04:05 MST 2006")
	log.Println(now)
	n := time.Now().Format(time.UnixDate)
	log.Println(n)

	var mon time.Month
	mon = 11
	log.Println(mon)

	// 加载Location
	loc, _ := time.LoadLocation("Local")
	log.Println(loc)

}

func example2() {
	// 返回当前时间的time类型
	now := time.Now()
	fmt.Println(now)
	pstr := "2018-01-01 12:22:22"
	// 格局layout解析一个时间字符串为time类型
	ptime, _ := time.Parse("2006-01-02 15:04:05", pstr)
	fmt.Println(ptime)
	//
	pltime, _ := time.ParseInLocation("2006-01-02 15:04:05", pstr, time.Local)
	fmt.Println(pltime)
	lo := pltime.Location()
	fmt.Println(lo)

	unix := now.Unix()
	fmt.Println(unix)

	a, n := now.Zone()
	fmt.Printf("%v %v", a, n)

}

// Duration example
func example3() {
	// Duration
	now := time.Now()
	d := time.Since(now)
	fmt.Println(d)

	pd, _ := time.ParseDuration("10h300ms")
	fmt.Println(pd)

}

func TestTimer() {
	t := time.NewTimer(10 * time.Second)
	fmt.Println(time.Now().Unix())
	for {
		a, ok := <-t.C
		if !ok {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix())
		} else {
			fmt.Println(a.Unix())
			fmt.Println(a)
			break
		}
	}

}
