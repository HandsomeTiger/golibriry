package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var locker mutex

func mutexDemo() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			locker.Lock()
			defer locker.Unlock()
			fmt.Printf("Goroutine %d : value: %d\n", i, locker.value)
			locker.value++
			time.Sleep(1 * time.Second)
		}(i)
	}
}
func main() {
	//go mutexTest()
	//go mutexDemo()
	//time.Sleep(time.Second * 6)
	PoolDemo(os.Stdout, "key", "value")
}

type mutex struct {
	value int
	sync.Mutex
}

func mutexTest() {
	for i := 1; i < 100; i++ {
		go func() {
			//locker.Lock()
			//defer locker.Unlock()
			fmt.Println(locker.value)
			locker.value++
		}()
	}
}

var pool sync.Pool = sync.Pool{
	// //New()函数的作用是当我们从Pool中Get()对象时，如果Pool为空，则先通过New创建一个对象，插入Pool中，然后返回对象。
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func PoolDemo(w io.Writer, key, val string) {
	b := pool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(time.Now().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	// 将临时对象放回到Pool中
	pool.Put(b)
}
