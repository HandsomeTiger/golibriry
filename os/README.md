# package os
os包提供了**操作系统**函数的不依赖平台的接口。  
##### os包的接口规定为在所有操作系统中都是一致的。非公用的属性可以从操作系统特定的syscall包获取。
### os常用的接口
```
func Hostname() (name string, err error) // Hostname返回内核提供的主机名
func Environ() []string // Environ返回表示环境变量的格式为”key=value”的字符串的切片拷贝
func Getenv(key string) string //  Getenv检索并返回名为key的环境变量的值
func Getpid() int // Getpid返回调用者所在进程的进程ID
func Exit(code int) // Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行
func Stat(name string) (fi FileInfo, err error) // 获取文件信息
func Getwd() (dir string, err error) // Getwd返回一个对应当前工作目录的根路径
func Mkdir(name string, perm FileMode) error // 使用指定的权限和名称创建一个目录
func MkdirAll(path string, perm FileMode) error // 使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误
func Remove(name string) error // 删除name指定的文件或目录
func TempDir() string // 返回一个用于保管临时文件的默认目录
var Args []string // os.Args返回一个字符串数组，其中第一个参数就是执行文件本身
```
### File 结构体：
file 结构体主要处理文件的创建，打开，读取，写入，关闭操作
```
func Create(name string) (file *File, err error) // Create采用模式0666创建一个名为name的文件，如果文件已存在会截断（为空文件）
func Open(name string) (file *File, err error) // Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据
func (f *File) Stat() (fi FileInfo, err error) // Stat返回描述文件f的FileInfo类型值
func (f *File) Readdir(n int) (fi []FileInfo, err error) //  Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo
func (f *File) Read(b []byte) (n int, err error) // Read方法从f中读取最多len(b)字节数据并写入b
func (f *File) WriteString(s string) (ret int, err error) // 向文件中写入字符串
func (f *File) Sync() (err error) // Sync递交文件的当前内容进行稳定的存储
func (f *File) Close() error // Close关闭文件f，使文件不能用于读写
```
* f.Read() 返回长度，如果为0是时返回err，通知指针位置指向该位置
* f.Write() 从指针位置写入内容到文件
* f.Seek(0,0) 移动指针位置 

### FileInfo 结构体 返回文件的属性
> os.stat, Lstat, os.File.stat 返回的都是FileInfo结构
```
type FileInfo interface {
   Name() string       // base name of the file
   Size() int64        // length in bytes for regular files; system-dependent for others
   Mode() FileMode     // file mode bits
   ModTime() time.Time // modification time
   IsDir() bool        // abbreviation for Mode().IsDir()
   Sys() interface{}   // underlying data source (can return nil)
}```

### 文档资料
[Golang读写文件的几种方式](https://www.jianshu.com/p/7790ca1bc8f6)