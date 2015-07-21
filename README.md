# 学习杂记
学习项目
按照语言类型分类

# go语言学习
## flag

> flag为golang用于处理命名行参数的一个库。

```golang
flag.xxx(x1, x2, x3)
```

该函数用于指定命令字对应的参数类型，并返回对应数据的指针
- **x1**:命令字名称
- **x2**:命令字不存在时的默认值
- **x3**:命令字的说明

例如：
```golang
path := flag.String("d", "/home/zjh", "user directory")
```

函数返回解析到数据的指针，之后调用flag.Parse(),解析出数据，并填充到对应指针指向的数据段中，这之后，我们就可以访问到对应的解析数据。

例如：
```golang
fmt.Printf("path   : %s\n", *path)
```

## net包
> func (c *TCPConn) Read(b []byte) (int, error)

疑惑：这里传入一个切片，那么当缓冲区的数据比切片空间大，会改变切片的大小吗？
答案：不会，如果切片大小为1，那么读取出来的数据长度也为1，然后循环读取，因此这里在读取时应该根据自身的业务开辟一个比较合理的大小，提高读取效率。

## 获取code.google.com上的代码

hg clone https://code.google.com/p/goprotobuf/ 这样将克隆下来对应代码，并放在goprotobuf文件夹下


## log4go的坑
code.google.com/p/log4go，使用源码中举的例子既然没有打印输出，百事不得其解，最后发现是log4go的bug，这个bug是程序退出可能有日志没有flush到磁盘。例如编写如下例子：
```
func main() {

	//log4go.Debug("test")
	log := log4go.NewDefaultLogger(log4go.DEBUG)
	defer log.Close()

	log.AddFilter("log", log4go.DEBUG, log4go.NewFileLogWriter("example.log", true))
	log.Info("test log output")
}
```
你会发现不管是控制台，还是文件，都没有任何输出，即使你加上defer log.Close()，因为Close也是异步的，最后我尝试加上time.Sleep(time.Second * 10)，终于输出了。