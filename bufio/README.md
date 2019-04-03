# package bufio
> bufio模块通过对io模块的封装，提供了数据缓冲功能，能够一定程度减少大块数据读写带来的开销。
  在bufio各个组件内部都维护了一个缓冲区，数据读写操作都直接通过缓存区进行。当发起一次读写操作时，会首先尝试从缓冲区获取数据；只有当缓冲区没有数据时，才会从数据源获取数据更新缓冲。

* Reader
```go
type Reader struct {
   buf          []byte
   rd           io.Reader // reader provided by the client
   r, w         int       // buf read and write positions
   err          error
   lastByte     int
   lastRuneSize int
}
```

可以通过NewReader函数创建bufio.Reader对象，函数接收一个io.Reader作为参数。因此，bufio.Reader不能直接使用，需要绑定到某个io.Reader上。

```go
func NewReader(rd io.Reader) *Reader
func NewReaderSize(rd io.Reader, size int) *Reader // 可以配置缓冲区的大小
```
相较于io.Reader，bufio.Reader提供了很多实用的方法，能够更有效的对数据进行读取。bufio.Reader能够对Reader进行细粒度的操作：  
A、Read，读取n个byte数据  
B、Discard，丢弃接下来n个byte数据  
C、Peek，获取当前缓冲区内接下来的n个byte，但不移动指针  
D、Reset，清空整个缓冲区  
```go
func (b *Reader) Read(p []byte) (n int, err error)
func (b *Reader) Discard(n int) (discarded int, err error)
func (b *Reader) Peek(n int) ([]byte, error)
func (b *Reader) Reset(r io.Reader)
```
bufio.Reader还提供了多个更高抽象层次的方法对数据进行简单的结构化读取，如下：  
A、ReadByte，读取一个byte  
B、ReadRune，读取一个utf-8字符  
C、ReadLine，读取一行数据，由’\n’分隔  
D、ReadBytes，读取一个byte列表  
E、ReadString，读取一个字符串  
```go
func (b *Reader) ReadByte() (byte, error)
func (b *Reader) ReadRune() (r rune, size int, err error)
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
func (b *Reader) ReadBytes(delim byte) ([]byte, error)
func (b *Reader) ReadString(delim byte) (string, error)
```

* Writer
```go
type Writer struct {
   err error
   buf []byte
   n   int
   wr  io.Writer
}
```
```go
func NewWriter(w io.Writer) *Writer
func NewWriterSize(w io.Writer, size int) *Writer
```
创建Writer对象的接口: 
```go
func (b *Writer) Write(p []byte) (nn int, err error) // 写入n byte数据
func (b *Writer) Reset(w io.Writer) // 重置当前缓冲区
func (b *Writer) Flush() error // 清空当前缓冲区，将数据写入输出
func (b *Writer) WriteByte(c byte) error  // 写入一个字节
func (b *Writer) WriteRune(r rune) (size int, err error） // 写入一个字符
func (b *Writer) WriteString(s string) (int, error) // 写入一个字符串
func (b *Writer) Available() int // 缓存中有多少字节空间可用
func (b *Writer) Buffered() int // 当前缓存已经写入了多少字节
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error) // 实现io.ReaderFrom
```
* ReadWriter
```go
type ReadWriter struct {
   *Reader
   *Writer
}
```
ReadWriter实现了io.ReadWriter。  
func NewReadWriter(r *Reader, w *Writer) *ReadWriter  
ReadWriter对象创建
  
* Scanner
```go
type Scanner struct {
   r            io.Reader // The reader provided by the client.
   split        SplitFunc // The function to split the tokens.
   maxTokenSize int       // Maximum size of a token; modified by tests.
   token        []byte    // Last token returned by split.
   buf          []byte    // Buffer used as argument to split.
   start        int       // First non-processed byte in buf.
   end          int       // End of data in buf.
   err          error     // Sticky error.
   empties      int       // Count of successive empty tokens.
   scanCalled   bool      // Scan has been called; buffer is in use.
   done         bool      // Scan has finished.
}
```
工程开发中推荐使用Scanner对数据进行读取，而非直接使用Reader类。Scanner可以通过splitFunc将输入数据拆分为多个token，然后依次进行读取。
func NewScanner(r io.Reader) *Scanner
创建scanner对象
func (s *Scanner) Split(split SplitFunc)
设置scanner的分割函数。
在使用scanner前还需要设置splitFunc（默认为ScanLines），splitFunc用于将输入数据拆分为多个token。bufio模块提供了几个默认splitFunc，能够满足大部分场景的需求，包括：
A、ScanBytes，按照byte进行拆分
B、ScanLines，按照行(“\n”)进行拆分
C、ScanRunes，按照utf-8字符进行拆分
D、ScanWords，按照单词(” “)进行拆分
通过Scanner的Split方法，可以为Scanner指定splitFunc。使用方法如下：

scanner := bufio.NewScanner(os.StdIn)
scanner.split(bufio.ScanWords）
设置分割方式
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
函数接收两个参数，第一个参数是输入数据，第二个参数是一个标识位，用于标识当前数据是否为结束。函数返回三个参数，第一个是本次split操作的指针偏移；第二个是当前读取到的token；第三个是返回的错误信息。

```go
func (s *Scanner) Scan() bool
func (s *Scanner) Text() string
func (s *Scanner) Bytes() []byte
```
在完成Scanner初始化后，通过Scan方法可以在输入中向前读取一个token，读取成功返回True；使用Text和Bytes方法获取token，Text返回一个字符串，Bytes返回字节数组。
Scanner使用示例：

```go
package main

import (
   "strings"
   "fmt"
   "bufio"
)

func main() {
   scanner := bufio.NewScanner(strings.NewReader("hello world !"))
   scanner.Split(bufio.ScanWords)
   for scanner.Scan() {
      fmt.Println(scanner.Text())
   }
}
```
// output
// hello
// world
// !




  
  
  
### 参考资料  
* [golang bufio、ioutil读文件的速度比较（性能测试）和影响因素分析](https://studygolang.com/articles/11452)
* [51cto bufio](https://blog.51cto.com/9291927/2294279)