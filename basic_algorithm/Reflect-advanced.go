package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func main() {
	user := User{
		Name: "123",
	}
	auto(reflect.TypeOf(&user), reflect.ValueOf(&user), "ReflectCallFuncHasArgs", "123", 10)
}

func auto(val reflect.Type, par reflect.Value, f string, params ...interface{}) {
	for i := 0; i < val.NumMethod(); i++ {
		method := val.Method(i)
		paramObject := par.Method(i)
		paramcount := paramObject.Type().NumIn()
		var methodValue reflect.Value
		var args []reflect.Value
		if method.Name == f {
			methodValue = par.MethodByName(f)
		}
		for n := 0; n < paramcount; n++ {
			for _, v := range params {
				if paramObject.Type().In(n).Name() == reflect.TypeOf(v).Name() {
					args = append(args, reflect.ValueOf(v))
				}
			}
		}
		methodValue.Call(args)
	}
}
