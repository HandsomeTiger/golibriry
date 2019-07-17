# package net/url
> url包实现了对url的解析和加密的方法

## 
* func QueryEscape(s string) string // 对url进行转义（反斜线），类似于php的 addslashes方法
* func QueryUnescape(s string) (string, error)  // 对url进行反转义

## type URL
> type url URL类型代表一个解析后的URL（或者说，一个URL参照）。URL基本格式如下：
`scheme://[userinfo@]host/path[?query][#fragment]`
### URL 主要方法：
* func Parse(rawurl string) (url \*URL, err error) 传入一个原始url返回*URL结构体
* func ParseRequestURI(rawurl string) (url *URL, err error)
* func (u *URL) IsAbs() bool
* func (u *URL) Query() Values
* func (u *URL) RequestURI() string
* func (u *URL) String() string
* func (u *URL) Parse(ref string) (*URL, error)
* func (u *URL) ResolveReference(ref *URL) *URL

## type Userinfo
> Userinfo类型是一个URL的用户名和密码细节的一个不可修改的封装。一个真实存在的Userinfo值必须保证有用户名（但根据 RFC 2396可以是空字符串）以及一个可选的密码。  

## type Values
> Values将建映射到值的列表。它一般用于查询的参数和表单的属性。不同于http.Header这个字典类型，Values的键是大小写敏感的。

### Values 方法
* func ParseQuery(query string) (m Values, err error)
* func (v Values) Get(key string) string  // Get会获取key对应的值集的第一个值。如果没有对应key的值集会返回空字符串。获取值集请直接用map。
* func (v Values) Set(key, value string) // Set方法将key对应的值集设为只有value，它会替换掉已有的值集。
* func (v Values) Add(key, value string) // Add将value添加到key关联的值集里原有的值的后面。
* func (v Values) Del(key string) // Del删除key关联的值集。
* func (v Values) Encode() string // Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。

## 注意
golang里面的urlencode和urldecode是使用net/url 里面的url.Values.Encode 和 ParseQuery来实现的

## 参考资料
* [golang标准库文档+源码](https://studygolang.com/pkgdoc)
* [Golang语言标准库http/url的Values的详细介绍](https://studygolang.com/articles/01752)
* [Go语言学习之net/url包(url相关操作)(the way to go)](https://blog.csdn.net/wangshubo1989/article/details/75017632)
 
