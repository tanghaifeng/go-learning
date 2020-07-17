package interface_01

import "fmt"

type Animal interface {
	Run()
	Eat()
}

// dog
type Dog struct {
}

// go 的接口是鸭子类型
func (d Dog) Run() {
	fmt.Println("dog run")
}

func (d Dog) Eat() {
	fmt.Println("dog eat food")
}

// Cat
type Cat struct {
}

func (d Cat) Run() {
	fmt.Println("car run")
}

func (d Cat) Eat() {
	fmt.Println("car eat food")
}
