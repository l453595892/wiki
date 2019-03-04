## 指针,地址,形参,实参
```
package main

import (
    "fmt"
)

func main(){
    vat i int
    var p *int
    i = 1
    p = &i
    fmt.Printf("i=%d;p=%d;*p=%d\n",i,p,*p)
    *p = 2
    fmt.Printf("i=%d;p=%d;*p=%d\n",i,p,*p)
    i = 3
    fmt.Printf("i=%d;p=%d;*p=%d\n",i,p,*p)
}

```
<br> 这段代码的结果是:
- i=1;p=0x4212e100;*p=1
- i=2;p=0x4212e100;*p=2
- i=3;p=0x4212e100;*p=3

```
package main

import (
    "fmt"
)

type abc struct{
    v int
}

func (a abc) aaaa(){
    a.v=1
    fmt.Printf("1:%d\n",a.v)
}

func (a *abc) bbbb(){
    fmt.Printf("2:%d\n",a.v)
    a.v=2
    fmt.Printf("3:%d\n",a.v)
}

func (a *abc) cccc(){
    fmt.Printf("4:%d\n",a.v)
}

func main(){
    aobj := abc{}
    aobj.aaaa()
    aobj.bbbb()
    aobj.cccc()
}
```
<br> 这段代码的结果是:
- 1:1
- 2:0
- 3:2
- 4:2

<br> aaaa的 [接收实体] (也就是abc)是一个实参,在go语言中,实参其实就是将参数的值复制到函数里来(参数与函数调用前在内存里的地址是不一样的).bbbb和cccc的[接收实体]是一个形参,也就是说,函数调用前后参数所在内存地址是一样的!所以bbbb中,第一行的v还没赋值所以为0,第二行的v赋值2以后在cccc中打印v的值也为2
<br> 对于[goroutin],[切片],[映射]这三种类型来说,只有形参,而且不需要加[*]号
<br> 对于参数类型是[interface]的函数参数,只有实参,而且不会将[interface]结构所包含的地址复制

<br> **golang中只存在值拷贝**
```
package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("value: %#v\n", b) // b的值是a的地址
	fmt.Printf("addr: %#v\n", &b) // b的地址
	pFoo(b)
	fmt.Printf("value: %#v\n", a) // a的值发生改变
}

func pFoo(p *int) {
	fmt.Printf("value: %#v\n", p) // 拷贝了b的值(也就是a的地址)
	fmt.Printf("addr: %#v\n", &p) // p的地址
	*p = 2
}
```