package main

import (
	"go-learning/base"
	_ "go-learning/base"
	"go-learning/gindemo"
)

type chan1 struct {
}

func main() {
	////const a int = 1
	////fmt.Printf("%T\n", a)
	//base.SetMember()
	//m := base.Member{}
	//m.SetMember()

	base.FmtDemo()
	gindemo.GinServer()

}
