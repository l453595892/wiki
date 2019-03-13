## Golang高级技术（适合中高级gopher）
注：在google上有很多初级入门级golang介绍,例如：[Go语言圣经](https://wizardforcel.gitbooks.io/gopl-zh/index.html)

### 栈中分配的变量，出栈的时候不会被释放吗

```
func newInt *int {
        var i int
        return &i   //为何可以返回局部变量呢？
    }
    someInt := newInt()
```
- 释不释放跟golang的GC有关系，记得golang用的是跟OC一样的基于引用计数的GC方案。所以即使是在函数内部alloc的，如果return回去了编译器会将这个对象的引用计数+1导致不会被GC回收。

### golang逃逸分析
- go在一定程度消除了堆和栈的区别，因为go在编译的时候进行逃逸分析，来决定一个对象放栈上还是放堆上，不逃逸的对象放栈上，可能逃逸的放堆上。

#### 逃逸分析的用处（为了性能）
- 最大的好处应该是减少gc的压力，不逃逸的对象分配在栈上，当函数返回时就回收了资源，不需要gc标记清除。
- 因为逃逸分析完后可以确定哪些变量可以分配在栈上，栈的分配比堆快，性能好
- 同步消除，如果你定义的对象的方法上有同步锁，但在运行时，却只有一个线程在访问，此时逃逸分析后的机器码，会去掉同步锁运行。

```
go run -gcflags '-m -l' main.go
```

### Golang使用pprof监控性能及GC调优

### Golang 汇编
go tool compile -S main.go
