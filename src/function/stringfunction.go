package function

import (
	"fmt"
	"strconv"
	"strings"
)

/*
内置函数(builtin)可以不用导包，直接使用
字符串相关函数
1. 统计字符串的长度，按字节进行统计：len(str)
2. 字符串遍历：r:= []rune(str)
3. 字符串转整数： n, err := strconv.Atoi("66")
4. 整数转字符串： str = strconv.ltoa(6887)
5. 查找字符串是否在指定的字符串中： strings.Contains("javagolang", "go")
6. 统计一个字符串有几个指定的子串： strings.Count("javaandgolang", "a")
7. 不区分大小写的字符串比较： fmt.Println(strings.EqualFold("go","Go"))
8. 返回字串在字符串第一次出现的索引值，如果没有返回-1：strings.Index("javaandgolang", "a")
9. 字符串替换： strings.Replace("goandjavagogo", "go", "golang", n)
			n可以指定你希望替换几个，如果n=-1表示全部替换，替换两个n就是2
10. 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
			strings.Split("go-python-java", "-")
11. 将字符串的字母进行大小写的转换
			strings.ToLower("Go")
			strings.ToUpper("go")
12. 将字符串左右两边的空格去掉： strings.TrimSpace("   go and java   ")
13. 将字符串左右两边指定的字符去掉： strings.Trim("~golang~", "~");
14. 将字符串左边指定的字符去掉：strings.TrimLeft("~golang~","~");
15. 京字符串右边指定的字符串去掉：strings.TrimRight("~gonglang~", "~");
16. 判断字符串是否指定的字符串开头: strings.HasPrefix("https://java.sun.com/jsp/jstl/fmt", "http")
17. 判断字符串是否以指定的字符串结束：


*/

func StringFunctions() {
	str := "golang你好"     // 一个汉字3个字节
	fmt.Println(len(str)) //12

	//遍历方式一：for range键值循环
	for i, value := range str {
		fmt.Printf("Index: %d, Value: %c\n", i, value)
	}
	/*
		Index: 0, Value: g
		Index: 1, Value: o
		Index: 2, Value: l
		Index: 3, Value: a
		Index: 4, Value: n
		Index: 5, Value: g
		Index: 6, Value: 你
		Index: 9, Value: 好
	*/
	//遍历方式二：利用[]rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Print(r[i])
		fmt.Printf("  %c", r[i])
		fmt.Println()
	}
	num1, _ := strconv.Atoi("666")
	fmt.Println(num1)

	str1 := strconv.Itoa(88)
	fmt.Println(str1)

	count := strings.Count("golangandjava", "an")
	fmt.Println(count)

	strEquals := strings.EqualFold("go", "Go")
	fmt.Println(strEquals)

	fmt.Println("go" == "Go")
	fmt.Println(strings.Index("goandgoland", "bd"))
}
