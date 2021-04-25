package collction

import "fmt"

func MapDemo() {
	// make函数创建map
	dict := make(map[string]interface{})
	dict["name"] = "Tim"
	dict["age"] = 30
	fmt.Println(dict)

	// 字面量

	dict1 := map[string]interface{}{"age": 30, "name": "Tim"}
	fmt.Println(dict1)

	fmt.Println(dict["name"])

	value, ok := dict["name"]
	if ok {
		fmt.Println(value)
	}

	member := map[string]interface{}{
		"name": "Tim",
		"age":  30,
	}
	fmt.Println(member)

	delete(member, "age")
	fmt.Println(member)
}
