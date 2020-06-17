package sync_learn_01

import (
	"fmt"
	"sync"
)

/**
提供了基本的同步原语

模型：
	加锁
	TODO
	解锁
*/

var count int
var lock sync.Mutex

func Increment() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Println("increment: ", count)
}

func Decrement() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Println("decrement: ", count)
}
