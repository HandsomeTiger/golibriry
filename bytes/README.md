# package bytes
> bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似。  
> 字节切片处理的函数比较多，分为`基本处理函数`、`比较函数`、`后缀检查函数`、`索引函数`、`分割函数`、`大小写处理函数`和`子切片处理函数`等。  

### 基本函数
* func Contains(b,subslice []bytes) bool
> 检查字节切片b是否包含子切片subslice,如果包含返回true,否则返回false。
* func Count(s,sep []byte) int
> Count计算s中有多少个不重叠的sep子切片。
* func Repeat(b[]byte,count int) []byte
> 类似于php中的str_repeat
* func Replace(s,old,new []byte,n int) []byte
> 把s中的old替换为new，替换n次（如果n<0则替换所有），返回替换后的值
* func Runes(s []byte) []rune
> Runes函数返回和s等价的[]rune切片。（将utf-8编码的unicode码值分别写入单个rune）
* func Join(s [][]byte,sep[]byte) []byte
> 类似于php的implode

### 比较函数
* func Compare(a,b[]byte) int
> 根据`字节的值`比较字节切片a和b的大小,如果a=b,返回0,如果a>b返回1,如果a小于b返回-1。`类似于php中的strcmp(),按ascii码从左向右比较`
* func Equal(a,b[]byte) bool
> 比较两个字节切片是否相等
* func EqualFold(s,t[]byte) bool
> 判断两个utf-8编码切片（将unicode大写、小写、标题三种格式字符视为相同）是否相同。(`把s和t转换成UTF-8字符串进行比较,并且忽略大小写,如果s=t,返回true,否则,返回false。`)

### 前后缀检查
* func HasPrefix(s,prefix[]byte) bool
> 检查字节切片s的前缀是否为prefix,如果是返回true,如果不是返回false
* func HashSuffix(s,suffix[]byte) bool
> 检查字节切片s的后缀是否为suffix,如果是返回true,否则返回false。

### 位置索引(`重点理解rune，byte，utf-8之间的关系和区别`)
> 字节切片位置索引函数共有8个，Index()、IndexAny()、IndexByte()、IndexFunc()、IndexRune()、LastIndex()、LastIndexAny()和LastIndexFunc()。
* func Index(s,sep []byte) int
> 子切片sep在s中第一次出现的位置，不存在则返回-1。
* func IndexAny(s []byte,chars string) int
> 把s解析为UTF-8编码的字节序列,返回chars中任何一个字符在s中第一次出现的索引位置;如果s中不包含chars中任何一个字符,则返回-1
* func IndexByte(s[]byte,c byte) int
> 检查字节c在s中第一次出现的位置索引;如果s中不包含c则返回-1. `类似于php中的strpos`
* func IndexFunc(s[]byte,f func(r rune)bool) int
> s中第一个满足函数f的位置i（该处的utf-8码值r满足f(r)==true），不存在则返回-1
* func LastIndex(s,sep[]byte) int
* func LastIndexAny(s[]byte,chars string) int
* func LastIndexFunc(s[]byte,f func(r rune)bool) int

### 分割函数
> 字节切片分割函数共有6个,Fields(),FieldsFunc(),Split(),SplitN(), 
* func Fields(s[]byte) [][]byte
> 按空白分割
* func FieldsFunc(s []byte,f func(r rune)bool) [][]byte
* func Split(s,sep[]byte)[][]byte
> 把s用sep分割成多个字节切片并返回,如果sep为空,Split则把s切分成每个字节切片对应一个UTF-8字符,Split()等效于参数为n的splitN()函数。
* func SplitN(s,sep []byte,n int)[][]byte
* func SplitAfter(s,sep[]byte)[][]byte
* func SplitAfterN(s,sep[]byte,n int)[][]byte

### 大小写处理函数
> 共有7个函数，Title()，ToTitle()，ToTitleSpecial()，ToLower()，ToLowerSpecial()，ToUpper() 和ToUpperSpecial()。

### 子字节切片处理函数
> 子字节切片处理函数共有9个，Trim()，TrimFunc()，TrimLeft()，TrimLeftFunc()，TrimRight()，TrimRightFunc()，TrimSpace()，TrimPrefix()和TrimSuffix()。





