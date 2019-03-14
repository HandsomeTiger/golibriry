package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	example1()
}

func example1() {
	// ParseBool 把一个字符串解析成bool类型
	bstr := "T"
	b, err := strconv.ParseBool(bstr)
	if err != nil {
		fmt.Println(err.Error())
	}
	logDebuf(b)
	istr := "1110111011101011"
	// ParseInt 把base进制的str转换成int64类型
	i, err := strconv.ParseInt(istr, 2, 64)
	logDebuf(i)
	// FormatInt 把int类型的i转换成base进制的字符串类型
	i2 := strconv.FormatInt(i, 8)
	logDebuf(i2)
	i3 := strconv.FormatInt(i, 16)
	logDebuf(i3)

	byteStr := make([]byte, 100)
	logDebuf(byteStr)
	// 把farmatint 追加到[]byte
	i4 := strconv.AppendInt(byteStr, i, 8)
	logDebuf(string(i4))
	logDebuf(string(byteStr))
	//	 Quote
	strQ := `jksdhgfjka/kjlsdfh\s`
	i5 := strconv.Quote(strQ)
	logDebuf(i5)
	var a byte
	a = 'N' + 'h'
	logDebuf(a)
}

func logDebuf(v interface{}) {
	l := log.New(os.Stdout, "strconv_logDebug: ", log.LstdFlags)
	out := fmt.Sprintf("%+v", v)
	l.Output(0, out)
}
