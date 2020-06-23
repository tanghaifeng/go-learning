package tests

import (
	sync_learn_01 "go-learning/sync-learn-01"
	"testing"
)

func TestOnce(t *testing.T) {
	// 打印的是1 而不是100
	sync_learn_01.Once()
}
