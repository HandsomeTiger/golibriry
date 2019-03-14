# package log
> Go语言中log模块用于在程序中输出日志。  
  log模块提供了三类日志输出接口，Print、Fatal和Panic。Print是普通输出；  
  Fatal是在执行完Print后，执行 os.Exit(1)；  
  Panic是在执行完Print后调用panic()方法。  
  log模块对每一类接口其提供了3中调用方式，分别是"Xxxx、 Xxxxln、Xxxxf"。
  
## 基本用法
```go
 const (
     // 字位共同控制输出日志信息的细节。不能控制输出的顺序和格式。
     // 在所有项目后会有一个冒号：2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
     Ldate         = 1 << iota     // 日期：2009/01/23
     Ltime                         // 时间：01:23:23
     Lmicroseconds                 // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
     Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
     Lshortfile                    // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
     LstdFlags     = Ldate | Ltime // 标准logger的初始值
 )
```

```go
package main

import (
   "log"
)
func logPrintTest(){
   arr := []int {2,3}
   log.Print("Print array ",arr,"\n")
   log.Println("Println array",arr)
   log.Printf("Printf array with item [%d,%d]\n",arr[0],arr[1])
   // Fatal等价于{l.Print(v...); os.Exit(1)} 输出之后关闭程序
   // exit（0）：正常运行程序并退出程序；
   // exit（1）：非正常运行导致退出程序；
   log.Fatal("Print array ", arr, "\n")
   // Panic等价于{l.Print(v...); panic(...)} 输出之后抛出异常
   log.Panic("Print array ", arr, "\n")
  
}

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
	// Logger.Output是执行具体的将日志刷入到对应位置的方法
	logger.Output(2, "output test")
}
```

logger 定制:  
```go
func New(out io.Writer, prefix string, flag int) *Logger {
   return &Logger{out: out, prefix: prefix, flag: flag}
}
```
> 函数接收三个参数分别是日志输出的IO对象，
日志前缀和日志包含的通用信息标识位，
通过对参数进行设置可以对Logger进行定制。其中IO对象通常是标准输出os.Stdout，os.Stderr，或者绑定到文件的IO。
日志前缀（prefix）和信息标识位（flags）可以对日志的格式进行设置。

一条日志由三个部分组成，其结构如下：  
`{prefix} {flag1} {flag2} ... {flagn} {日志内容}`

log 日志分级：  
debug，error ... [代码实现参考](https://blog.51cto.com/9291927/2294270)  
[github log项目](github.com/sirupsen/logrus)


## 文档：
[Golang 源码剖析：log 标准库](https://www.cnblogs.com/lalalagq/p/9968997.html)  
[Golang 标准库log的实现](https://www.cnblogs.com/lvdongjie/p/6511029.html)
[log-51cto](https://blog.51cto.com/9291927/2294270)