package main

import (
	"fmt"
)

func main1() {
	var age int = 18
	//	&变量 可以获取这个变量的内存地址
	fmt.Println(&age)

	//	定义一个指针变量
	//	ptr是指针变量的名字，ptr对应的类型是*int，它是一个指针类型（可以理解为指向int类型的指针）
	//	&age是一个地址，是ptr变量的具体的值
	//	总结：指针就是一个内存地址
	//	&获取内存地址，*获取变量的值
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr本身这个存储空间的地址为：", &ptr)
	//	如果想获取ptr这个指针指向的数值
	fmt.Println("ptr指向的数值的值：", *ptr)

	//	可以通过指针改变指向值
	//	指针变量接收的一定是地址值 (不可以这样，只能赋值地址 ptr *int = num)
	//	指针变量的地址不可以不匹配(不可以这样，类型不匹配 ptr *float = &intnumber)
	//	基本数据类型（又叫做值类型），都有对应的指针类型，形式为*数据类型: *byte,*float32, *int64
	var num int = 10
	fmt.Println(num)
	var ptr2 *int = &num
	*ptr2 = 20
	fmt.Println(num)
}
