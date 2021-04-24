package collction

import (
	"fmt"
)

/**
slice是基于数组，数据结构有三个字段 指向数组的指针，存放的长度，容量，可以动态扩容

*/
func SliceDemo() {
	// slice make 声明
	slice1 := make([]int, 2) // 利用内置函数make来声明函数，make函数声明了切片并且进行了初始化赋值 赋类型的0值
	fmt.Println(slice1)
	slice2 := make([]int, 2, 5)
	fmt.Println(slice2)

	// 字面量声明
	slice3 := []int{1, 23, 4, 6}
	fmt.Println(slice3)

	var arr = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	// 长度 arr[i:j] j-i 容量 k-i
	arr2 := arr[1:2] // 下表为1的 第二个且不包括第三个的数 [1:2)
	fmt.Println(arr2, len(arr2), cap(arr2))

	// 迭代
	for k, v := range arr {
		fmt.Println(k, v)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
