package channel_01

import (
	"fmt"
	"sync"
)

// 简单的channel创建
var wg sync.WaitGroup

func Channel() {
	wg.Add(2)
	defer wg.Done()
	c := make(chan int, 1) // 有缓存的channel chan 可以是任意类型，interface{}
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	wg.Wait()
}
