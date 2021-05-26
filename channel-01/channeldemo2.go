package channel_01

import "fmt"

func RunChannel() {
	// 无缓存channel
	c := make(chan int)
	defer close(c)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
