package main

import "fmt"

func main() {
	array := [5]int{10, 20, 30, 40, 50}
	array1 := [...]int{10, 20, 30, 40, 50}
	array2 := [...]int{1: 10, 20, 30, 40, 50}

	fmt.Println(array)
	fmt.Println(array1)
	fmt.Println(array2)

	var arrayStr [3]*string
	arrayStr2 := [3]*string{new(string), new(string), new(string)}

	*arrayStr2[0] = "Red"
	*arrayStr2[1] = "Blue"
	*arrayStr2[2] = "Green"
	arrayStr = arrayStr2
	fmt.Println(*arrayStr2[0])
	*arrayStr2[0] = "Red--------modify"
	fmt.Println(*arrayStr[0])

	// 声明一个8MB的数组
	var arrayInt [1e6]int
	// 将数组传给函数
	foo(arrayInt)
	slice := make([]string, 3, 5)
	slice1 := []int{10, 20, 30}
	fmt.Println("切片容量：", cap(slice))
	fmt.Println(slice1)
	// 映射
	dict := make(map[string]int)
	dict1 := map[string]int{"red": 1, "blue": 2}
	fmt.Println(dict)
	fmt.Println(dict1)
}
func foo(array [1e6]int) {

}
