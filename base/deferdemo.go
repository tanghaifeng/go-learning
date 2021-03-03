package base

import (
	"fmt"
)

/**
defer3
defer2
defer1
defer 应该是个LIFO（先进后出）的栈
*/
func DeferDemo() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
}

/**
defer1
panic: panic
在panic发生后 定义的defer也将不再执行，所以最好是将defer写在函数开头的位置

defer 常用于关闭文件 管道关闭等
*/
func DeferDemo1() string {
	defer fmt.Println("defer1")
	panic("panic")
	defer fmt.Println("defer2")
	fmt.Println("11")
	return "111"
}
