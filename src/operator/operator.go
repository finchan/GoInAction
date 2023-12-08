package main

/*
		/ 1. 算数运算符: + - * / % ++ --
		/ 2. 赋值运算符: =, +=, -=, *=, /=, %=
		/ 3. 关系运算符：==, !=, >,<, >=, <=
运算符-
		\ 4. 逻辑运算符: &&， ||， ! (注意短路特点)
		\ 5. 位运算符：&, |, ^
		\ 6. 其他运算符：&, *
		\ 7. 位移运算符：<<, >>
*/

import (
	"fmt"
)

func main() {
	var n1 int = 10
	fmt.Println(n1)
	var n2 int = n1 + 10
	fmt.Println(n2)
	var s1 string = "abc" + "def"
	fmt.Println(s1)
	fmt.Println("=================================")
	fmt.Println(10 / 3)
	fmt.Println(10.0 / 3)
	fmt.Println("=================================")
	//a%b === a-a/b*b
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)
	fmt.Println("=================================")
	//	++自增加1操作，--自减，减1操作
	//	Go语言里，++，--只能单独使用，不能参与到运算中，不能用于赋值
	//	Go语言里，++，--只能在变量的后面，不能在变量的前面，--a,++a这样的写法都是错的
	var n3 int = 10
	n3++
	fmt.Println(n3)
	fmt.Println("=================================")
	//	键盘录入年龄、姓名、成绩，是否VIP（bool）
	//	方式1：Scanln - 录入数据的类型一定要匹配
	var age int
	var name string
	var score float32
	var isVIP bool
	fmt.Println("请录入学生的年龄：")
	//	传入age的地址的目的是，在Scanln函数中，对地址中的值进行改变的时候，实际外面的age被影响了
	fmt.Scanln(&age)
	fmt.Println("请录入学生的姓名：")
	fmt.Scanln(&name)
	fmt.Println("请录入学生的成绩：")
	fmt.Scanln(&score)
	fmt.Println("请录入学生是否VIP：")
	fmt.Scanln(&isVIP)
	fmt.Printf("年龄：%v, 姓名： %v, 成绩：%v, 是否VIP：%v", age, name, score, isVIP)
	fmt.Println("=================================")
	//	键盘录入年龄、姓名、成绩，是否VIP（bool）
	//	方式2：Scanf - 录入数据的类型一定要匹配
	fmt.Println("请录入学生的年龄，姓名，成绩，是否VIP，使用空格分隔")
	var age2 int
	var name2 string
	var score2 float32
	var isVIP2 bool
	fmt.Scanf("%d %s %f %t", &age2, &name2, &score2, &isVIP2)
	fmt.Printf("年龄：%v, 姓名： %v, 成绩：%v, 是否VIP：%v", age2, name2, score2, isVIP2)
}
