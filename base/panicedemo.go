package base

import "errors"

//发生 panic 程序会马上奔溃
func PanicDemo() {
	panic(errors.New("Panic"))
}
