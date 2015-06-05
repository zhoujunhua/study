# study
学习项目
按照语言类型分类
# golang
## flag

> flag为golang用于处理命名行参数的一个库
----------
```golang
flag.xxx(x1, x2, x3)
```
该函数用于指定命令字对应的参数类型，并返回对应数据的指针
- **x1**:命令字名称
- **x2**:命令字不存在时的默认值
- **x3**:命令字的说明
-例如：
```golang
path := flag.String("d", "/home/zjh", "user directory")
```
函数返回解析到数据的指针
之后调用flag.Parse(),解析出数据，并填充到对应指针指向的数据段中

这之后，我们就可以访问到对应的解析数据
例如：fmt.Printf("path   : %s\n", *path)

# java
# c/c++
# nodejs
# html/css

