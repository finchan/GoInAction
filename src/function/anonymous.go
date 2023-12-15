package function

/*
1. Go支持匿名函数，如果某个函数只使用一次，可以考虑使用匿名函数
2. 匿名函数的使用方式：
		在定义匿名函数是就直接调用，这种方式只能调用一次(用的多）
		将匿名函数给一个变量，再通过该变量来调用匿名函数
3， 如何让一个匿名函数在整个过程中有效呢？将匿名函数给一个全局变量就可以了
*/

import "fmt"

var fun01 = func(num1 int, num2 int) int {
	return num1 * num2
}

func Anonymous() {
	//	定义匿名函数
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println(result)

	//	将匿名函数赋值给一个函数变量
	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Println(sub(10, 20))
	fmt.Println(fun01(10, 20))
}
