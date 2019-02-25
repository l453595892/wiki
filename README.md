# wiki
一个关于我所了解的技术的wiki指南，将我所认为的比较好的网络内容都记录在这个wiki里面

## 请点击Wiki来了解你想知道的技术，当然你也可以和我联系一起学习一起做k8s相关方向

## GOLANG

### 栈中分配的变量，出栈的时候不会被释放吗
```
func newInt *int {
        var i int
        return &i   //为何可以返回局部变量呢？
    }
    someInt := newInt()
```

- 释不释放跟golang的GC有关系，记得golang用的是跟OC一样的基于引用计数的GC方案。所以即使是在函数内部alloc的，如果return回去了编译器会将这个对象的引用计数+1导致不会被GC回收。
