package main

import (
	_ "go-learning/base"
	interface_01 "go-learning/interface-01"
)

type chan1 struct {
}

func main() {

	t := &interface_01.Tim{"Tim"}
	b := &interface_01.Ben{"Ben"}

	interface_01.Read(b)

	interface_01.Read(t)

}
