package sync_learn_01

import (
	"sync"
	"time"
)

// 读写锁

// 生产者
func producer(wg *sync.WaitGroup, lock sync.Locker) {
	defer wg.Done()
	for i := 5; i > 0; i-- {
		lock.Lock()
		lock.Unlock()
		time.Sleep(time.Second)
	}
}

func observer(wg *sync.WaitGroup, lock sync.Locker) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
}

func Test(count int, mutex, rwMutex sync.Locker) time.Duration {
	var wg sync.WaitGroup
	wg.Add(count + 1)
	beginTestTime := time.Now()
	go producer(&wg, mutex)
	for i := count; i > 0; i-- {
		go observer(&wg, rwMutex)
	}
	wg.Wait()
	return time.Since(beginTestTime)
}
