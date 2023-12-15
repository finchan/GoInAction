package main

import (
	cal "calcutils"
	"fmt"
	"function"
	"utils" //	包路径，自src以下的通过/a/b/c到最终的文件夹
)

/*
1. 建议包名和文件夹名称相同
2. main包（package main）是程序入口，一般main函数(func main())放在这个包下，否则无法编译执行
3. 函数调用前面要有包的定位
4. 如果希望被其他包调用，函数名或者变量首字母要大些的
5. 一个目录下不能有重复的函数
6. 包名和文件夹名可以不一样
7. 一个目录下的同级文件对应的包名必须相同
8. 可以给包取别名，取别名后，原来的包名就不能使用了
*/
//	下载main前后都行
var num int = test()

func test() int {
	fmt.Println("test函数被执行")
	return 10
}

func init() {
	fmt.Println("init 函数被执行")
}

func main() {
	fmt.Println("Main Methods - ")
	fmt.Printf(util.StuName)
	util.GetConn()
	cal.Test()
	fmt.Println("-------------------------------")
	fmt.Println(util.Age)
	function.Anonymous()
	function.Closure()
	fmt.Println(function.DeferTest(10, 20))
}
