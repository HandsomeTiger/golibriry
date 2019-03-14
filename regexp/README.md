# package regexp
> regexp包实现了正则表达式搜索

regexp.Regexp{} 是一个编译好的正则表达式对象。  
通常需要使用正则表达式构建一个正则对象
```go
type Regexp struct {
   // read-only after Compile
   regexpRO

   // cache of machines for running regexp
   mu      sync.Mutex
   machine []*machine
}
```
常用接口：  
```go
// 将正则表达式编译成一个正则对象（使用PERL语法）该正则对象会采用“leftmost-first”模式。选择第一个匹配结果
Compile(expr string) (*Regexp, error)
// 将正则表达式编译成一个正则对象（正则语法限制在 POSIX ERE 范围内）。该正则对象会采用“leftmost-longest”模式。选择最长的匹配结果。
CompilePOSIX(expr string) (*Regexp, error)
// 编译失败时panic
MustCompile
//让正则表达式在之后的搜索中都采用“leftmost-longest”模式。
func (re *Regexp) Longest()
//返回编译时使用的正则表达式字符串
func (re *Regexp) String() string

//返回正则表达式中分组的数量
func (re *Regexp) NumSubexp() int
//返回正则表达式中分组的名字
//第 0 个元素表示整个正则表达式的名字，永远是空字符串。
func (re *Regexp) SubexpNames() []string
//返回正则表达式必须匹配到的字面前缀（不包含可变部分）。
//如果整个正则表达式都是字面值，则 complete 返回 true。
func (re *Regexp) LiteralPrefix() (prefix string, complete bool)


// MatchString reports whether the byte slice b
// contains any match of the regular expression pattern.
// More complicated queries need to use Compile and the full Regexp interface.
func Match(pattern string, b []byte) (matched bool, err error) {
	re, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.Match(b), nil
}
// Match 方法实际上就是吧正则字符串转换成正则表达式对象，然后调用re.Match方法

func MatchString(pattern string, s string) (matched bool, err error)
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
```
Find() 方法返回匹配到的第一个[]byte切片，如果没有匹配到返回nil  
FindIndex() 方法返回匹配到的最左侧的第一个切片的起始位置下标
FindSubmatch() 方法返回匹配到最左侧的结果和分组信息 其中[0]表示获取到的完整匹配 [1] 表示第一个分组...  
FindAll() 返回保管正则表达式re在b中的所有不重叠的匹配结果的[][]byte切片。如果没有匹配到，会返回nil。  

// 当正则表达式中包含能接受重复的限定符时，通常的行为是（在使整个表达式能得到匹配的前提下）匹配尽可能多的字符。以这个表达式为例：a.*b，它将会匹配最长的以a开始，以b结束的字符串。如果用它来搜索aabab的话，它会匹配整个字符串aabab。这被称为贪婪匹配。

> 在匹配文本时，该正则表达式会尽可能早的开始匹配，并且在匹配过程中选择回溯搜索到的第一个匹配结果。这种模式被称为“leftmost-first”，Perl、Python和其他实现都采用了这种模式，但本包的实现没有回溯的损耗。对POSIX的“leftmost-longest”模式，参见CompilePOSIX。

关于 leftmost-longest 和leftmost-first 参考标准库文档对于 Compile的说明：
> func Compile
  func Compile(expr string) (*Regexp, error)
  Compile解析并返回一个正则表达式。如果成功返回，该Regexp就可用于匹配文本。
  
> 在匹配文本时，该正则表达式会尽可能早的开始匹配，并且在匹配过程中选择回溯搜索到的第一个匹配结果。这种模式被称为“leftmost-first”，Perl、Python和其他实现都采用了这种模式，但本包的实现没有回溯的损耗。对POSIX的“leftmost-longest”模式，参见CompilePOSIX。
  
> func CompilePOSIX
  func CompilePOSIX(expr string) (*Regexp, error)
  类似Compile但会将语法约束到POSIX ERE（egrep）语法，并将匹配模式设置为leftmost-longest。
  
> 在匹配文本时，该正则表达式会尽可能早的开始匹配，并且在匹配过程中选择搜索到的最长的匹配结果。这种模式被称为“leftmost-longest”，POSIX采用了这种模式（早期正则的DFA自动机模式）。
  
> 然而，可能会有多个“leftmost-longest”匹配，每个都有不同的组匹配状态，本包在这里和POSIX不同。在所有可能的“leftmost-longest”匹配里，本包选择回溯搜索时第一个找到的，而POSIX会选择候选结果中第一个组匹配最长的（可能有多个），然后再从中选出第二个组匹配最长的，依次类推。POSIX规则计算困难，甚至没有良好定义。
  
> 参见http://swtch.com/~rsc/regexp/regexp2.html#posix获取细节。
## 文档
[正则表达式30分钟入门教程](https://www.jb51.net/tools/zhengze.html)  
