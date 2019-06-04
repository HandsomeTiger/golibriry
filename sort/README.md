# package sort
> sort包主要提供了内置的一些排序的方法和自定义排序的接口

#### 核心 interface
```go
type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}
```

实现 该Interface接口就可以用sort包里的方法实现排序

* func Sort(data Interface)
* func Stable(data Interface)
* func IsSorted(data Interface) bool
* func Reverse(data Interface) Interface
* func Search(n int, f func(int) bool) int

#### 参考
* [golang标准版文档-sort](https://studygolang.com/pkgdoc)
