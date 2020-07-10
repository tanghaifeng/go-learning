package collction

import "fmt"

// 数组
/**
创建数组必须制定数组长度和类型，数组的存储的也是连续的一段内存空间
声明后所有的元素都是对应类型的0值 [3]int   // [0 0 0]
[3]string [  ] 空字符串
*/
func Array1() {
	var arr [3]string
	//arr = [3]int{}
	fmt.Println(arr)

}
