package main

import (
	"fmt"
)

type Operation struct {
	a float64
	b float64
}

func (op *Operation) SetA(a float64) {
	op.a = a
}

func (op *Operation) SetB(b float64) {
	op.b = b
}

type OperationI interface {
	GetResult() float64
	SetA(float64)
	SetB(float64)
}

type AddOperation struct {
	Operation
}

func (this *AddOperation) GetResult() float64 {
	return this.a + this.b
}

type SubOperation struct {
	Operation
}

func (this *SubOperation) GetResult() float64 {
	return this.a - this.b
}

type MulOperation struct {
	Operation
}

func (this *MulOperation) GetResult() float64 {
	return this.a * this.b
}

type DivOperation struct {
	Operation
}

func (this *DivOperation) GetResult() float64 {
	return this.a / this.b
}

type IFactory interface {
	CreateOperation() OperationI
}

type AddFactory struct {
}

func (this *AddFactory) CreateOperation() OperationI {
	return &(AddOperation{})
}

type SubFactory struct {
}

func (this *SubFactory) CreateOperation() OperationI {
	return &(SubOperation{})
}

type MulFactory struct {
}

func (this *MulFactory) CreateOperation() OperationI {
	return &(MulOperation{})
}

type DivFactory struct {
}

func (this *DivFactory) CreateOperation() OperationI {
	return &(DivOperation{})
}

func main() {
	fac := &(AddFactory{})
	oper := fac.CreateOperation()
	oper.SetA(1)
	oper.SetB(2)
	fmt.Println(oper.GetResult())
}
