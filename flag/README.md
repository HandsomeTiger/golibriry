package flag
> flag包实现了命令行参数的解析。

#### 基本用法
* flag.String(), Bool(), Int() 注册
* flag.Parse() 解析
    - **Flag解析在第一个非flag参数（单个"-"不是flag参数）之前停止，或者在终止符"--"之后停止。**
* flag.Args() 获取flag后面的参数
* 命令行
    - -flag 取默认值 
    - -flag=x
    - -flag x  // 只有非bool类型的flag可以
    - 可以使用1个或2个'-'号，效果是一样的。
* CommandLine 该包提供了一个默认变量：CommandLine，是FlatSet的一个实例
* FlagSet 
```go
type FlagSet struct {
    // Usage函数在解析flag出现错误时会被调用
    // 该字段为一个函数（而非采用方法），以便修改为自定义的错误处理函数
    Usage func() `这个方法用于覆盖默认的函数实现，可以自定义输出所有定义了的命令行参数和帮助信息`
    // 内含隐藏或非导出字段
}
```
* func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
> NewFlagSet创建一个新的、名为name，采用errorHandling为错误处理策略的FlagSet。
* func (f *FlagSet) SetOutput(output io.Writer)
> 设置使用信息和错误信息的输出流，如果output为nil，将使用os.Stderr。

#### 参考
* [golang flag包使用笔记](https://www.imooc.com/article/45805)
* [golang标准库-flag](https://studygolang.com/pkgdoc)
* [Golang学习笔记--flag包源码分析](https://blog.csdn.net/cza55007/article/details/43817239)
