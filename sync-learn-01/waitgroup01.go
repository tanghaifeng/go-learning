package sync_learn_01

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func add() {
	// defer 确保在goroutine退出之前执行Done,向WaitGroup表明 已经退出了
	defer wg.Done()
	fmt.Println("test add")
}

func AddTest() {
	// Add 表示一个goroutine 开始了
	wg.Add(1)
	go add()
	// 阻塞父协程
	// Wait 中有个for循环在不断的阻塞查看goroutine状态
	wg.Wait()
}
