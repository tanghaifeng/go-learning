package base

import "fmt"

func Recover() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("捕获异常")
		}
	}()
	panic("panic")
}
