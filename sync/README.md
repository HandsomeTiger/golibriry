# package sync
> sync包提供了基本的同步基元。

#### 常用方法 
* Locker
```go
    type Locker interface {
       Lock()
       Unlock()
    }
```
> Locker提供了锁的两种操作方法:Lock() 和Unlock()

* sync.Mutex 互斥锁
> sync.Mutex实现了Locker接口
> 互斥锁就是由任何一个线程锁定后，不允许对该锁的对象进行任何操作，只有在解锁之后，再重新争抢锁定
> **官方描述： Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁。意思是使用Lock方法的时候会检测该对象有没有加锁，如果有则阻塞，如过没有则加锁。那么如果另一个不需要Lock的方法是不是就不受影响，不会进行阻塞检测呢。（经测试，结果是这样的）**
      
注意：    
    * （1）在首次使用后不要复制互斥锁。
    * （2）对一个未锁定的互斥锁解锁将会产生运行时错误。

问题：
    * 有两个不同的方法对同一个对象进行操作，一个方法没锁，一个方法有锁，没锁的那个方法需要等解锁后才能执行吗？

* sync.Pool
> sync.Pool 临时对象池用于存储临时对象，将使用完毕的对象存入对象池中，在需要的时候取出来重复使用，目的是为了避免重复创建相同的对象造成GC负担过重。如果对象不再被其它变量引用，存放的临时对象可能会被GC回收掉。  
> 简单点描述就是可以存储对象的一个集合   
> **官方描述很清楚** Pool的合理用法是用于管理一组静静的被多个独立并发线程共享并可能重用的临时item。Pool提供了让多个线程分摊内存申请消耗的方法。
  * ```go
      type Pool struct {
            // 可选参数New指定一个函数在Get方法可能返回nil时来生成一个值    
            // 该参数不能在调用Get方法时被修改
            **New()函数的作用是当我们从Pool中Get()对象时，如果Pool为空，则先通过New创建一个对象，插入Pool中，然后返回对象。**
            New func() interface{}
            // 包含隐藏或非导出字段
        }
    ````
  * func (p *Pool) Get() interface{} 
  > Get方法去得p中的**任意**一个item，如果没有获取到item并且p的New() 为nil，那么返回nil，否则返回New()的值
  * func (p *Pool) Put(x interface{})
  
* sync.Once
  > Once是只执行一次动作的对象。Once只有一个方法Do，Do方法当且仅当第一次被调用时才执行函数f。
  * func (o *Once) Do(f func())

* sync.RWMutex
  > sync.RWMutex是针对读写操作的互斥锁，读写锁与互斥锁最大的不同就是可以分别对读、写进行锁定。一般用在大量读操作、少量写操作的情况。sync.RWMutex提供四种操作方法：
  ```go
    func (rw *RWMutex) Lock()
    func (rw *RWMutex) Unlock()
  
    func (rw *RWMutex) RLock()
    func (rw *RWMutex) RUnlock()
  ```
  RLock对读操作进行锁定，RUnlock对读锁定进行解锁，Lock对写操作进行锁定，Unlock对写锁定进行解锁。  
  sync.RWMutex锁定规则如下：  
  （1）同时只能有一个goroutine能够获得写锁定。    
  （2）同时可以有任意多个gorouinte获得读锁定。  
  （3）同时只能存在写锁定或读锁定（读和写互斥）。  
  （4）当有一个goroutine获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；当有一个goroutine获得读锁定，其它读锁定任然可以继续；当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁后才能够进行写锁定。  
  sync.RWMutex读写锁使用注意：  
  （1）在首次使用之后，不要复制读写锁。  
  （2）不要混用锁定和解锁，如：Lock和RUnlock、RLock和Unlock。对未读锁定的读写锁进行读解锁或对未写锁定的读写锁进行写解锁将会引起运行时错误。  

    如何理解读写锁呢？  
    * 同时只能有一个 goroutine 能够获得写锁定。
    * 同时可以有任意多个 gorouinte 获得读锁定。
    * 同时只能存在写锁定或读锁定（读和写互斥）。

* sync.WaitGroup
> sync.WaitGroup用于等待一组goroutine结束。  
  ```go
    func (wg *WaitGroup) Add(delta int)
    func (wg *WaitGroup) Done()
    func (wg *WaitGroup) Wait()
  ```
  
* sync.Cond 条件
> sync.Cond实现一个条件等待变量，即等待或宣布事件发生的goroutine的会合点。
```go
    func NewCond(l Locker) *Cond
    func (c *Cond) Broadcast()
    func (c *Cond) Signal()
    func (c *Cond) Wait()
```
    

#### 参考：  
* [Go语言开发（十三）、Go语言常用标准库三-sync](https://blog.51cto.com/9291927/2343533) 
* [Golang标准库-sync](https://studygolang.com/pkgdoc)
* [golang 关于锁 mutex,踩过的坑](https://blog.csdn.net/fwhezfwhez/article/details/82900498)
* [浅谈 Golang sync 包的相关使用方法](https://www.jianshu.com/p/2ec7479d8390)
* [golang之sync.Mutex互斥锁源码分析](https://www.jianshu.com/p/ffe646ada7b4)
* [源码剖析 Golang 中 sync.Mutex](https://www.lbbniu.com/7290.html)