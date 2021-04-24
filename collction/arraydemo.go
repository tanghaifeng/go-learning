package collction

import (
	"fmt"
	"reflect"
)

// 数组
/**
创建数组必须制定数组长度和类型，数组的存储的也是连续的一段内存空间
声明后所有的元素都是对应类型的0值 [3]int   // [0 0 0]
[3]string [  ] 空字符串
*/
func Array1() {
	var arr [3]string // 数组声明
	//arr = [3]int{}
	fmt.Println(arr)

	var arrayL = [1]int{1} // 字面量
	fmt.Println(arrayL)

	var interfaceArray = [2]interface{}{111, "34324"}
	for i := 0; i < len(interfaceArray); i++ {
		fmt.Print(reflect.TypeOf(interfaceArray[i]), "\t")
		fmt.Println(interfaceArray[i])
	}

	// 指针数组
	array := [2]*int{new(int), new(int)}
	fmt.Println(array)
	*array[0] = 100 // 给指向的内存地址赋值
	fmt.Println(*array[0])

	var arrayObj [2]interface{} // interface{} 空接口 可以存任意类型
	arrayObj[0] = 1
	arrayObj[1] = "Tim"
	fmt.Println(arrayObj)

	// 二维数组
	array2 := [4][2]int{{1, 2}, {2, 3}, {2, 3}, {2, 3}}
	fmt.Println(array2)

	// 大数据最好是传输一个指针
	// go 传参都是值值类型，foo接受的其实是指针的地址值 0xc00000e008
	var bigArray [1e6]int
	foo(&bigArray)
}

func foo(arr *[1e6]int) {
	fmt.Println(&arr)
}
