# package bufio
> bufio模块通过对io模块的封装，提供了数据缓冲功能，能够一定程度减少大块数据读写带来的开销。
  在bufio各个组件内部都维护了一个缓冲区，数据读写操作都直接通过缓存区进行。当发起一次读写操作时，会首先尝试从缓冲区获取数据；只有当缓冲区没有数据时，才会从数据源获取数据更新缓冲。
  
  
  
  
### 参考资料  
* [golang bufio、ioutil读文件的速度比较（性能测试）和影响因素分析](https://studygolang.com/articles/11452)
* [51cto bufio](https://blog.51cto.com/9291927/2294279)