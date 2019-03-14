# package strconv
> strconv包实现了基本数据类型和其字符串表示的相互转换。  

## 
* ParseBool 把字符串解析成bool类型
> 它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。  
* ParseInt 把字符串解析成int类型
`func ParseInt(s string, base int, bitSize int) (i int64, err error)`
> 返回字符串表示的整数值，接受正负号。  
> base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制； 表示s的进制   
> bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。  
* FormatBool 
> 把一个bool类型转换成string类型  
* FormatInt
> 把int类型转换成base进制的字符串
* AppendBool
> 等价于append(dst, FormatBool(b)...)
* func Quote(s string) string 返回字符串的双引号表示形式
> 返回字符串s在go语法下的双引号字面值表示，控制字符、不可打印字符会进行转义。（如\t，\n，\xFF，\u0100）


