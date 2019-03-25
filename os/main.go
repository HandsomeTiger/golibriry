package main

import (
	"fmt"
	"os"
)

func main() {
	//example1()
	//fmt.Println("exit test!")
	//exampleMk()
	//exampleMkAll()
	//exampleRemove()
	//example2()

	err := insertLog("logger ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")
}

// os常用接口示例
func example1() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("get host error %v", host)
	}
	fmt.Printf("host name = %v \n\n", host)
	envs := os.Environ()
	for _, env := range envs {
		fmt.Printf("env = %+v \n", env)
	}
	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)
	pid := os.Getpid()
	fmt.Printf("pid = %v\n", pid)
	// get stat
	stat, err := os.Stat("./test.txt")
	if err != nil {
		fmt.Printf("got file stat error: %v\n", err)
	}
	fmt.Printf("stat: %v\n", stat)
	fmt.Printf("stat name: %v\n", stat.Name())

	wd, err := os.Getwd()
	fmt.Printf("wd %v\n", wd)
	// exit 会结束当前程序
	os.Exit(888)

	fmt.Println("exit test!")

}

func exampleMk() {
	err := os.Mkdir("testmk", 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func exampleMkAll() {
	err := os.MkdirAll("testmk/all", 0777)
	if err != nil {
		fmt.Errorf("error : %v", err.Error())
	}
}

func exampleRemove() {
	os.RemoveAll("testmk/all")

}

func example2() {
	// 预定义变量, 保存命令行参数
	fmt.Println(os.Args)
	// 获取host name
	fmt.Println(os.Hostname())
	fmt.Println(os.Getpid())
	// 获取全部环境变量
	env := os.Environ()
	for k, v := range env {
		fmt.Println(k, v)
	}
	// 终止程序
	// os.Exit(1)
	// 获取一条环境变量
	fmt.Println(os.Getenv("PATH"))
	// 获取当前目录
	dir, err := os.Getwd()
	fmt.Println(dir, err)
	// 创建目录
	err = os.Mkdir(dir+"/new_file", 0755)
	fmt.Println(err)
	// 创建目录
	err = os.MkdirAll(dir+"/new", 0755)
	fmt.Println(err)
	// 删除目录
	err = os.Remove(dir + "/new_file")
	err = os.Remove(dir + "/new")
	fmt.Println(err)
	// 创建临时目录
	tmp_dir := os.TempDir()
	fmt.Println(tmp_dir)
}
