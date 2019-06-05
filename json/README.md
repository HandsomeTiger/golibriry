# package encoding/json
> json包实现了json对象的编解码

#### 基本用法

##### Marshal
`func Marshal(v interface{}) ([]byte, error)`
在编码过程中，json包会将Go的类型转换为JSON类型，转换规则如下： 
* bool 转换为JSON boolean 
* 浮点数, 整数, Number 转换为：JSON number 
* string转换为：JSON string 
* 数组、切片 转换为：JSON数组 
* []byte 转换为：base64 string 
* struct、map转换为：JSON object
> Marshal函数返回v的json编码。Marshal函数会递归的处理值。如果一个值实现了Marshaler接口且非nil指针，会调用其MarshalJSON方法来生成json编码。否则，Marshal函数使用默认编码格式。  

##### 注意：
* 结构体转换的时候，只有公共的会被转换，私有的不会被转换，因为json包会遍历结构体的每个值进行编码

##### 标签
* 指定编码后字段的名称
* omitempty 如果为空则忽略
* - 不解析
* string 选项标记一个字段在编码json时应编码为字符串。它只适用于字符串、浮点数、整数类型的字段。

##### 匿名字段
> json包在解析匿名字段时，会将匿名字段的字段当成该结构体的字段处理。

#####　转换接口
```go
type Marshaler interface {
   MarshalJSON() ([]byte, error)
}
```
##### 编码到输出流
func NewEncoder(w io.Writer) *Encoder
NewEncoder创建一个将数据写入w的*Encoder。
func (enc *Encoder) Encode(v interface{}) error
Encode将v的json编码写入输出流，并会写入一个换行符。

#### 解码
##### UnMarshal
`func Unmarshal(data []byte, v interface{}) error`
Unmarshal函数解析json编码的数据data并将结果存入v指向的值，v通常传入指针，否则解析虽不报错，但数据无法赋值到接受体中。
解码将JSON转换为Go数据类型。在解码过程中，json包会将JSON类型转换为Go类型，转换规则如下： 
* JSON boolean 转换为 bool 
* JSON number 转换为 float64 
* JSON string 转换为 string 
* JSON数组 转换为 []interface{} 
* JSON object 转换为 map 
* null 转换为 nil  

##### 其他用法跟编码相似 
...

#### 参考
* [Go语言开发（十六）、Go语言常用标准库六](https://blog.51cto.com/9291927/2344741)
* [Golang标准库文档-encodding/json](https://studygolang.com/pkgdoc)

