package collction

import (
	"fmt"
)

func SliceDemo() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	var arr1 = []int{1, 2, 3, 4, 5, 5}
	arr = append(append(arr, arr1...), 99)
	fmt.Println(arr)

}
