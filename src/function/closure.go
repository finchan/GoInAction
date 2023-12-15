package function

/*
闭包是返回匿名函数+匿名函数以外的变量组合的整体，匿名函数中引用的变量会一直保存在内存中，会一直使用
本质是匿名函数，不过是引用外面的变量/参数。闭包= 匿名函数+引用的变量/参数
闭包对内存消耗大
*/
import "fmt"

// getSum的返回值是一个函数(参数int，返回值int)
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}
func Closure() {
	f := getSum()
	fmt.Println(f(1)) //	1
	fmt.Println(f(2)) //	3
}
