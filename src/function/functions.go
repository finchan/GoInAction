package function

import "fmt"

/*
Go不支持方法重载（方法名相同，参数不同）
Go支持可变参数 ...（参数个数可变）
基本数据类型和数组默认都是值传递，进行值拷贝，在函数内修改，不会改变值
以值传递方式的数据类型，如果希望在函数内的变量能修改函数外的变量，可以传入变量的地址&,
				函数内以指针的方式操作变量。
在Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用。
函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用（把函数本身当作一种数据类型）。
为了简化数据类型定义，Go支持自定义数据类型，基本语法：
		type 自定义数据类型名 数据类型
		相当于起了一个别名，虽然是别名但在Go中编译识别还是认为这两种数据不是一种数据类型，不能赋值，但可以通过类型转换来实现num2 = int(num1)
		例如type myInt int ---> 这是myInt就等价int来使用了
		type mySum func(int, int) int  ---> mySum就等价一个函数类型func(int, int) int
支持对函数返回值命名
*/

func test(args ...int) {
	//	函数内部处理可变参数，将可变参数当作切片处理
	//	遍历可变数组
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func test2(num int) {
	num = 30
}

// 指针形式的值传递
func test3(num *int) {
	*num = 333
}

// 函数做参数
func test4(num1 int, testFunc func(int)) {
	fmt.Println("test4")
}

type myFunc func(int)

func test5(num1 int, testFunc myFunc) {
	fmt.Println("test5")
}

func test6(num1 int, num2 int) (sub int, sum int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main1() {
	fmt.Println("---")
	test()
	fmt.Println("---")
	test(3)
	fmt.Println("---")
	test(3, 4, 5)
	fmt.Println("---")
	var num int = 10
	test2(num)
	fmt.Println(num)
	var num2 int = 100
	test3(&num2)
	fmt.Println(num2)
	a := test3
	fmt.Printf("a的类型是： %T; test3的类型是%T\n", a, test3)
	num2 = 100
	a(&num2)
	fmt.Println(num2)
	test4(10, test2)
	test5(10, test2)
	m, n := test6(2, 3)
	fmt.Println(m, n)

}
