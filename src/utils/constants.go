package util

import "fmt"

var StuName string = "Jack" // 变量名首字母大写则可通过包进行访问
var Age int
var Sex string
var Name string

func init() {
	fmt.Println("constants init is invoked")
	Age = 19
	Sex = "男"
	Name = "张三"
}
