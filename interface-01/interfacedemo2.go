package interface_01

import "fmt"

type reader interface {
	read()
}

type Tim struct {
	Name string
}

func (t Tim) read() {
	fmt.Println(t.Name)
}

type Ben struct {
	Name string
}

func (b Ben) read() {
	fmt.Println(b.Name)
}

func Read(r reader) {
	r.read()
}
