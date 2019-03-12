package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	example4()
}

//
func example4() {
	// Flags返回标准logger的输出选项。
	log.Flags()
	log.SetFlags(log.Ldate)
	// Prefix返回标准logger的输出前缀。
	log.Prefix()
	//	 SetOutput设置标准logger的输出目的地，默认是标准错误输出。
	log.SetOutput(os.Stderr)
	// 新创建一个logger
	logger := log.New(os.Stderr, "example4：", log.LstdFlags)
	// 输出
	logger.Output(2, "output test")
}

// example1 基本用法
func example1() {
	arr := []int{2, 3}
	log.Print("Print array ", arr, "\n")
	log.Println("Println array", arr)
	log.Printf("Printf array with item [%d,%d]\n", arr[0], arr[1])
	log.Fatal("Print array ", arr, "\n")
	log.Fatalln("Println array", arr)
	log.Fatalf("Printf array with item [%d,%d]\n", arr[0], arr[1])

}

func example3() {
	level := "Panic"
	defer func() {
		fmt.Println("defer Panic 1")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Panic(level, " level")
	defer func() {
		fmt.Println("defer Panic 2")
	}()
}

func example2() {
	//创建输出日志文件
	logFile, err := os.Create("./" + time.Now().Format("20060102") + ".txt")
	if err != nil {
		fmt.Println(err)
	}

	//创建一个Logger
	//参数1：日志写入目的地
	//参数2：每条日志的前缀
	//参数3：日志属性
	loger := log.New(logFile, "test_", log.Ldate|log.Ltime|log.Lshortfile)

	//Flags返回Logger的输出选项
	fmt.Println(loger.Flags())

	//SetFlags设置输出选项
	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//返回输出前缀
	fmt.Println(loger.Prefix())

	//设置输出前缀
	loger.SetPrefix("test_")

	//输出一条日志
	loger.Output(2, "打印一条日志信息")

	//格式化输出日志
	loger.Printf("第%d行 内容:%s", 11, "我是错误")

	//等价于print();os.Exit(1);
	loger.Fatal("我是错误")

	//等价于print();panic();
	loger.Panic("我是错误")

	//log的导出函数
	//导出函数基于std,std是标准错误输出
	//var std = New(os.Stderr, "", LstdFlags)

	//获取输出项
	fmt.Println(log.Flags())
	//获取前缀
	fmt.Printf(log.Prefix())
	//输出内容
	log.Output(2, "输出内容")
	//格式化输出
	log.Printf("第%d行 内容:%s", 22, "我是错误")
	log.Fatal("我是错误")
	log.Panic("我是错误")
}
