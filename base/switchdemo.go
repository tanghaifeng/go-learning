package base

import "fmt"

/**
case 可以是表达式
*/
func Switch(a int) string {
	switch {
	case 0 < a && a < 60:
		fmt.Println("不及格")
		return "不及格"
	case a >= 60 && a <= 89:
		return "及格"
	case a >= 90:
		fmt.Println("优秀")
		return "优秀"
	}

	return "未知"
}

// 判断类型
func TypeSwitch(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("字符串")
	case int:
		fmt.Println("int整形")
	case bool:
		fmt.Println("布尔")
	}
}
