package function

import (
	"fmt"
	"time"
)

/*
	package time
1. 获取当前时间 time.Now() - 返回结构体，结构体对应的类型time.Time
2. 自定义格式 - datestr2 := current.Format("2006/01/02 15/04/05")
				这个时间数字是固定的，不能变化，是Go语言成立时编写到这里想到的模式串
*/

func DateTimeInvoked() {
	fmt.Println("======================")
	current := time.Now()
	fmt.Printf("%T, %v", current, current)
	// 调用结构体对应的各项值
	fmt.Printf("\n年：%v", current.Year())
	fmt.Printf("\n月：%v", int(current.Month()))
	fmt.Printf("\n日：%v", current.Day())
	// 这种新用法，新思路啊，Sprintf返回值
	datestr := fmt.Sprintf("当前：%d-%d-%d %d:%d:%d", current.Year(), current.Month(), current.Day(), current.Hour(), current.Minute(), current.Second())
	fmt.Println(datestr)
	//	格式化日期
	datestr2 := current.Format("2006/01/02 15/04/05")
	fmt.Println(datestr2)
	datestr3 := current.Format("2006/01/02 15:04")
	fmt.Println(datestr3)
	fmt.Println("\n======================")
}
