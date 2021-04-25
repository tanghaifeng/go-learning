package base

import "fmt"

// 自定义类型

type Member struct {
	name string
	age  int
}

func SetMember() {
	m := &Member{}
	m.age = 30
	m.name = "Tim"
	fmt.Println(m)

	var m1 Member
	m1.name = "Tim1"
	m1.age = 30
	fmt.Println(m1)
}

// 方法
func (m *Member) SetMember() {
	m.age = 18
	m.name = "Tim"
	fmt.Println(m)
}
