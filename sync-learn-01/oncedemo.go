package sync_learn_01

import (
	"fmt"
	"sync"
)

// once

var count1 int

func Increment1() {
	count1++
}

var once sync.Once
var wg1 sync.WaitGroup

func Once() {
	wg1.Add(100)
	for i := 0; i < 100; i++ {
		defer wg1.Done()
		go func() {
			once.Do(Increment1)
		}()
	}
	fmt.Println(count1)
	wg1.Wait()
}
