package main

import (
	_ "go-learning/base"
	channel_01 "go-learning/channel-01"
)

type chan1 struct {
}

func main() {
	//const a int = 1
	//fmt.Printf("%T\n", a)
	//base.SetMember()
	//m := base.Member{}
	//m.SetMember()
	channel_01.RunChannel1()

}
