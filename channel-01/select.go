package channel_01

import (
	"fmt"
	"time"
)

// select 没有测试顺序，没有满足条件执行也不会是失败
func SelectDemo() {
	c := make(chan interface{})
	fmt.Println("阻塞ing")
	select {
	case <-c:
		fmt.Println("1")
	case <-time.After(time.Second * 5): // timeout
		close(c)
		fmt.Println("close chan")
	}
}
