package main

import (
	"fmt"
	"utils" //	自定义包所对应的文件夹名，使用时使用packageName.属性名进行引用,最好的情况是文件夹名和包名相同
)

// 全局变量
var n7, n8 = 8, 9

// 也可以一次性声明
var (
	netty = "netty"
	n9    = 9
)

/*
1. 标识符：数字、字母、下划线
2. 不能以数字开头，严格区分大小写，不能包含空格，不能用Go的保留字
3. 下划线本身是一个特殊的标识符，成为空标识符。代表任何其他标识符，
	但它对应的值被忽略，仅能被用作占位符使用，不能单独作为标识符使用。
	用处：1. 多个返回值，忽略其他返回值；2. import 下划线+"packageName"，可以忽略对应的包
4. 长度不限制，但不建议太长。
5. 命名规则：
	(1) 报名尽量保持package的名字和目录一致，尽量采取有意义的包名，简短，有意义，不要和标准库冲突
	(2) 变量名、函数名、常量名：采用驼峰法
	(3)	如果变量名、函数名 、常量名首字母大写，则可以被其他包访问；
		如果首字母小写，则只能在本包中使用
注意：
import导入语句通常放在文件开头包声明语句的下面。
导入的包名需要使用双引号包裹起来。
包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔。
main包是一个程序的入口包，所以main函数所在的包建议定义为main包，如果不定义为main包，那么就不能得到可执行文件。
		比如这样建立路径 - project/main/xxx.go (xxx.go 包含package main, func main)

*/

func main3() {
	//	以下定义在函数内，被称为局部变量
	//	变量不能重复声明，变量声明的类型与赋值的类型必须匹配
	//	变量如果只声明未赋值，使用时则使用其默认值，比如int的默认值0
	//	如果声明时没有给定变量类型，声明会根据赋值推断数据类型,赋值后再次赋值不能改变变量的类型。
	//	省略var及类型，注意用":="进行声明赋值，来判断对应类型
	//	支持多变量声明
	var age int //	声明
	age = 18    //	赋值
	fmt.Println("age = ", age)
	var age2 int = 19
	fmt.Println("age2 = ", age2)
	var age3 int
	fmt.Println("age3 = ", age3)
	var name = "tom"
	fmt.Println("name = ", name)
	sex := "MALE"
	fmt.Println("sex = ", sex)
	sex = "FEMALE"
	fmt.Println("sex = ", sex)
	fmt.Println("----------------------------------")
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	var n4, name2, n5 = 10, "Jack", 7.8
	fmt.Println(n4, name2, n5)
	n6, height := 6.9, 100.6
	fmt.Println(n6, height)
	fmt.Println("全局变量")
	fmt.Println(n7, n8)
	fmt.Println(n9, netty)
	fmt.Println(constants.StuName)
}
