package main

import "fmt"

/*
====================
if 条件表达式 {
	逻辑代码
}
注意：1. 条件表达式左右（）可以不写，也建议不写 2. if和表达式之间，一定要有空格 3. {}是必须的
=====================
if 条件表达式 {
	逻辑代码
} else {
	逻辑代码2
}
=====================
if 条件表达式 {
	逻辑代码
} else if 条件表达式2 {
	逻辑代码
}
...
 else {
	逻辑代码2
}
====================
switch 表达式 {
	case 值1, 值2, ...:
		语句块1
	case 值3, 值4....:
		语句块2
	...
	default:
		语句块
}
注意：1. switch后是一个表达式（常量值、变量、一个有返回的函数等都可以）
	2. case后的各个值的数据类型，必须和switch的表达式数据类型一致
	3. case后面多个表达式使用逗号分隔
	4. case后面表达式如果是常量值（字面量），要求不能重复
	5. case后面不需要带break，执行完自动退出，如果一个都匹配不到执行default
	6. default不是必须的，位置也是随意的
	7. switch后可以不代表达式，当作if分支来使用 (不推荐)
	8. switch后面可以直接声明/定义一个变量，分号结束 (不推荐)
	9. switch穿透，利用fallthrough关键字，如果case语句块后增加fallthrough，则会继续执行下一个case
====================
for 初始表达式; 布尔表达式; 迭代因子  {
	循环体
}
//	以下类似于while循环
i := 1
for i <= 5 {
	i++
}
//	死循环
for {

}
//	死循环2
for ;; {

}
====================
for range结构可以遍历数组、切片、字符串、map及通道，类似于其他语言的foreach
for key, val := range coll {

}
====================
几个用于循环的关键字： 1. break 2. continue 这两个关键词同其他语言，都是针对直接包含它们的循环体进行处理。如果要跨循环，使用label标签。
					3. goto实现程序的跳转（一般与条件语句配合使用，跳转到指定的label（标签），但不建议使用）
					4. return可以用于for循环，其实它是直接结束当前的函数的，其实是结束函数
*/

func main() {
	var count int = 20
	if count < 30 {
		fmt.Printf("库存不足")
	}
	if count2 := 20; count2 < 30 {
		fmt.Printf("库存2不足")
	}
	if count3 := 20; count3 < 30 {
		fmt.Printf("库存3不足")
	} else {
		fmt.Printf("库存3充足")
	}
	//	switch转成if的一种特殊情况见下例
	var a int = 1
	switch {
	case a == 1:
		fmt.Println("aaa")
	case a == 2:
		fmt.Println("bbb")
	}
	//	switch后面可以直接声明/定义一个变量，分号结束
	switch b := 7; {
	case b > 6:
		fmt.Println("大于6")
	case b <= 6:
		fmt.Println("小于等于6")
	}
	//	for
	var sum int = 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)
	//	for
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	//	for range
	var str string = "hello golang你好"
	//	方式1： 普通for循环
	fmt.Println("不支持中文，或者Unicode码，仅支持ASCII")
	for i := 0; i < len(str); i++ {
		//	仅能通过字节遍历，不支持中文，需要通过切片进行遍历，否则出现乱码
		fmt.Printf("%c", str[i])
	}
	fmt.Println()
	//	方式2： for range //	支持中文，是对字符进行遍历的
	for i, value := range str {
		fmt.Printf("索引为%d,具体值为：%c\n", i, value)
	}
	//	break 指定的层级
label2:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i: %v, j: %v\n", i, j)
			if i == 2 && j == 2 {
				break label2
			}
		}
	}

}
