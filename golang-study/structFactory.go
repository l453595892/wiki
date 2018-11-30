package main

import "fmt"

type K8sObjectFactory interface {
	GetObject() Object
}

type Object interface {
	Get()
	Update()
	Delete()
	Add()
}

type DeploymentFactory struct{}

func (*DeploymentFactory) GetObject() Object {
	return &Deployment{}
}

type Deployment struct{}

func (d *Deployment) Get() {
	fmt.Println("get deployment")
}

func (d *Deployment) Update() {
	fmt.Println("update deployment")
}

func (d *Deployment) Delete() {
	fmt.Println("delete deployment")
}

func (d *Deployment) Add() {
	fmt.Println("delete deployment")
}

// ----------
type ServiceFactory struct{}

func (*ServiceFactory) GetObject() Object {
	return &Service{}
}

type Service struct{}

func (s *Service) Get() {
	fmt.Println("get service")
}

func (s *Service) Update() {
	fmt.Println("update service")
}

func (s *Service) Delete() {
	fmt.Println("delete service")
}

func (s *Service) Add() {
	fmt.Println("delete service")
}

func main() {
	var factory K8sObjectFactory
	factory = &DeploymentFactory{}
	factory.GetObject().Add()

	factory = &ServiceFactory{}
	factory.GetObject().Add()
}
