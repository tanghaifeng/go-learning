package sync_learn_01

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func add() {
	defer wg.Done()
	fmt.Println("测试add")
}

func AddTest() {
	wg.Add(1)
	go add()
	wg.Wait()
}
