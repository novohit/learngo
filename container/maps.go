package main

import "fmt"

func main() {
	// map[key类型]value类型{} 是一个hashmap 无序
	m := map[string]string{
		"name":  "zwx",
		"age":   "12",
		"phone": "1234556",
	}
	fmt.Println(m)
	fmt.Println("Getting values")
	fmt.Println(m["name"])
	fmt.Println(m["test"]) // 打印空串
	name, ok1 := m["name"]
	test, ok2 := m["test"]
	fmt.Println(name, ok1)
	fmt.Println(test, ok2)
	// 一般配合branch用
	if phone, ok3 := m["phone"]; ok3 {
		fmt.Println(phone)
	} else {
		fmt.Println("key does not exist")
	}

	m2 := make(map[string]int) // make函数定义map
	var m3 map[string]int      // var 定义map

	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	fmt.Println("Deleting values")
	delete(m, "name")
	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("name key does not exist")
	}
}
