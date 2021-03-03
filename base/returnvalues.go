package base

import (
	"errors"
)

func Write(b []byte) (n int, err error) {
	return 0, errors.New("这是个错误")
}
