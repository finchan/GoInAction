package function

/*
为了在函数执行完后，及时释放资源，Go设计者提供defer关键字
应用场景：比如像关闭使用的资源，在使用的时候随手defer，因为defer有延迟执行机制（函数执行完毕在执行defer压入栈的语句），
所以用完随手写了关闭，省心省事。
*/

import "fmt"

func DeferTest(num1 int, num2 int) int {
	//	Go语言中，程序遇到defer关键字，不会立即执行defer后的语句,而是将defer后的语句压入一个栈中，
	//			也会将相关的值同时拷贝入栈中，不会随着函数后面的变化而变化。
	//			继续执行函数后面的语句。
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)
	//	栈的特点是先进后出，在函数执行完后（return），从栈中去除语句执行，取的顺序是先进后出
	var sum int = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}
