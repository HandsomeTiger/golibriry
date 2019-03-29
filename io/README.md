# package io
* io包是对底层I/O实现了一些封装的接口
* EOF = "end of file"
> io包提供了对I/O原语的基本接口。本包的基本任务是**包装这些原语已有的实现**（如os包里的原语），使之成为共享的公共接口，这些公共接口抽象出了泛用的函数并附加了一些相关的原语的操作。
  
> 因为这些接口和原语是对底层实现完全不同的低水平操作的包装，除非得到其它方面的通知，客户端不应假设它们是并发执行安全的。

### 核心
* type Reader      
    * Reader.Read(p []byte)(n int,err error)
    > Read方法读取len(p)字节数据写入p。它返回写入的字节数和遇到的任何错误。
    > 当Read到文件结尾时,会返回读到的内容的字节数和一个err==EOF 或err==nil
    
* type Writer
    * Write(p []byte) (n int, err error)
   
* type ReadWriter
    *  type ReadWriter interface {
           Reader
           Writer
       }


       
### 參考資料
* [golang中io包用法（二）](https://blog.csdn.net/chenbaoke/article/details/42458915)
> 实现了io.Reader接口或io.Writer接口的类型 
  
  初学者看到函数参数是一个接口类型，很多时候有些束手无策，不知道该怎么传递参数。还有人问：标准库中有哪些类型实现了io.Reader或io.Writer接口？
  
  通过本节上面的例子，我们可以知道，os.File同时实现了这两个接口。我们还看到 os.Stdin/Stdout这样的代码，它们似乎分别实现了 io.Reader/io.Writer接口。没错，实际上在os包中有这样的代码：
  
  var (
  	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
  	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
  	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
  )
  
  也就是说，Stdin/Stdout/Stderr 只是三个特殊的文件（即都是os.File的实例），自然也实现了io.Reader和io.Writer。
  
  目前，Go文档中还没法直接列出实现了某个接口的所有类型。不过，我们可以通过查看标准库文档，列出实现了io.Reader或io.Writer接口的类型（导出的类型）：
  
  - os.File 同时实现了io.Reader和io.Writer
  - strings.Reader 实现了io.Reader
  - bufio.Reader/Writer 分别实现了io.Reader和io.Writer
  - bytes.Buffer 同时实现了io.Reader和io.Writer
  - bytes.Reader 实现了io.Reader
  - compress/gzip.Reader/Writer 分别实现了io.Reader和io.Writer
  - crypto/cipher.StreamReader/StreamWriter 分别实现了io.Reader和io.Writer
  - crypto/tls.Conn 同时实现了io.Reader和io.Writer
  - encoding/csv.Reader/Writer 分别实现了io.Reader和io.Writer
  - mime/multipart.Part 实现了io.Reader</span>
  
  除此之外，io包本身也有这两个接口的实现类型。如：
  
  实现了Reader的类型：LimitedReader、PipeReader、SectionReader
  实现了Writer的类型：PipeWriter
  
  以上类型中，常用的类型有：os.File、strings.Reader、bufio.Reader/Writer、bytes.Buffer、bytes.Reader
  
  **小贴士**
  
  从接口名称很容易猜到，一般地，Go中接口的命名约定：接口名以er结尾。注意，这里并非强行要求，你完全可以不以 er 结尾。标准库中有些接口也不是以 er 结尾的。
  
  > Go 语言中，为了方便开发者使用，将 IO 操作封装在了如下几个包中： 
    - io 为 IO 原语（I/O primitives）提供基本的接口 
    - io/ioutil 封装一些实用的 I/O 函数 
    - fmt 实现格式化 I/O，类似 C 语言中的 printf 和 scanf 
    - bufio 实现带缓冲I/O
    --------------------- 
    作者：张威伦 
    来源：CSDN 
    原文：https://blog.csdn.net/qq_30259339/article/details/54745112 
    版权声明：本文为博主原创文章，转载请附上博文链接！