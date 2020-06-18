package tests

import (
	"fmt"
	sync_learn_01 "go-learning/sync-learn-01"
	"sync"
	"testing"
)

func TestCacl(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			sync_learn_01.Increment()
		}()

	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			sync_learn_01.Decrement()
		}()

	}
	wg.Wait()
	fmt.Println("complete!")
}

//
//func BenchTestCacl(b *testing.B) {
//	var wg sync.WaitGroup
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func() {
//			wg.Done()
//			sync_learn_01.Increment()
//		}()
//
//	}
//
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func() {
//			wg.Done()
//			sync_learn_01.Decrement()
//		}()
//
//	}
//	wg.Wait()
//	fmt.Println("complete!")
//}
