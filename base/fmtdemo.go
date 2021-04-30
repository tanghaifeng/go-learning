package base

import "fmt"

type user struct {
	name string
}

func FmtDemo() {
	u := user{name: "tang"}
	fmt.Printf("%v\n", u)
	fmt.Printf("%#v\n", u)
	fmt.Println("%T\n", u)
}
