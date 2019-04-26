package main

import "fmt"

const (
	Add      = "+"
	Multiple = "*"
)

type Operator interface {
	Operate(int, int) int
}

type AddOperate struct {
}

func (this *AddOperate) Operate(rhs int, lhs int) int {
	return rhs + lhs
}

type MultipleOperate struct {
}

func (this *MultipleOperate) Operate(rhs int, lhs int) int {
	return rhs * lhs
}

type OperateFactory struct {
}

func NewOperateFactory() *OperateFactory {
	return &OperateFactory{}
}

func (this *OperateFactory) NewOperate(operatename string) Operator {
	switch operatename {
	case "+":
		return &AddOperate{}
	case "*":
		return &MultipleOperate{}
	default:
		panic("无效运算符号")
		return nil
	}
}

// 简单工厂
func main() {
	Operator := NewOperateFactory().NewOperate(Add)
	fmt.Printf("add result is %d\n", Operator.Operate(1, 2))
}
