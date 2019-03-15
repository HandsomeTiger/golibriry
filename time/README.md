# package time

### struct
* Weekday
* Month
* Location 
* Time

### constants 主要用来time.Parse和time.Format
```go
const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // 使用数字表示时区的RFC822
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // 使用数字表示时区的RFC1123
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // 方便的时间戳
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
```

### time包常用方法
* func (d Weekday) String() string 
> String返回该日（周几）的英文名（"Sunday"、"Monday"，……）
* func LoadLocation(name string) (*Location, error)
> LoadLocation返回使用给定的名字创建的Location。
* func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
> Date 格式化一个时间类型
* ParseInLocation 
> ParseInLocation类似Parse但有两个重要的不同之处。第一，当缺少时区信息时，Parse将时间解释为UTC时间，而ParseInLocation将返回值的Location设置为loc；第二，当时间字符串提供了时区偏移量信息时，Parse会尝试去匹配本地时区，而ParseInLocation会去匹配loc。
* Unix 返回unix时间戳
* UnixNano 返回纳秒时间戳
* Zone 返回时区和相对于UTC的时间偏移量
* func (t Time) In(loc *Location) Time 
> In返回采用loc指定的地点和时区，但指向同一时间点的Time。如果loc为nil会panic。
* IsZone 

### time类型比较
* IsZone 判断time是不是零值
* Equal 等于
* Before 在time之前
* After 在time之后

### time.time结构体方法
* Clock
* Year
* Month
* Day
* Hour
* Minute
* Second
* Nanosecond
* Add(d)
* AddDate()
* Sub(d)
* Format

###  type Duration  (d)
> Duration类型代表两个时间点之间经过的时间，以*纳秒*为单位。可表示的最长时间段大约290年。

Duration 是时间段的基本单位
```go
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

func ParseDuration(s string) (Duration, error)
// Since返回从t到现在经过的时间，等价于time.Now().Sub(t)。
func Since(t Time) Duration
func (d Duration) Hours() float64
func (d Duration) Minutes() float64
func (d Duration) Seconds() float64
func (d Duration) Nanoseconds() int64
func (d Duration) String() string
```

### type Timer
NewTimer

### type Ticker

