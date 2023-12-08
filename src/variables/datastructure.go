package main

//https://studygolang.com/
import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
				/整数型(int, int8, int16, int32, uint, uint8, uint16, uint32, uint64, byte)
				/浮点型(float32,float64)
			/数值型
			/字符型(没有单纯的字符型，使用byte来保存单个字母字符)
			/布尔型(bool)
			/字符串(string)
	/基本数据类型

数据类型

	\派生数据类型/复杂数据类型
			\指针
			\数组
			\结构体
			\管道
			\函数
			\切片
			\接口
			\map

	 1. 整型(超出会报overflow异常)：
	    int8	- 有符号 - 1字节 - -2^7-2^7-1(-128~127) -128 = 10000000, 127 = 01111111
		int16	- 有符号 - 2字节 - -2^15-2^15-1(-32768~32767)
		int32	- 有符号 - 4字节 - -2^31-2^31-1(-2147483648~2147483647)
		int64	- 有符号 - 8字节 - -2^63-2^63-1
		uint8	- 无符号 - 1字节 - 0~255
		uint16	- 无符号 - 2字节 - 0~2^16-1
		uint32	- 无符号 - 4字节 - 0~2^32-1
		uint64	- 无符号 - 8字节 - 0~2^64-1
		int		- 有符号 - 32位系统4字节 - -2^31-2^31-1
		int		- 有符号 - 64位系统8字节 - -2^63-2^63-1
		uint	- 无符号 - 32位系统4字节 - 0-2^32-1
		uint	- 无符号 - 64位系统8字节 - 0-2^64-1
		rune	- 有符号 - 等价于int32   - -2^31-2^31-1(-2147483648~2147483647)
		byte	- 无符号 - 等价于uint8 - 0~255
	2. 浮点类型(符号位+指数位+尾数为，尾数为只存一个大概，会有精度的损失),通常情况下建议使用float64(默认)
		float32		4字节	-3.403E38~3.403E38 (E或者e均可)
		float64		8字节	-1.798E308~1.798E308
	3. 字符型 (byte, uint) - Golang使用Unicode字符集，使用utf-8这种编码方案。
		所以ASCII码字符对应数字可以使用byte，超过255，使用uint
		转义字符 - 	\n \b \t(一个制表位，和前面字符凑8个，不足补空格)
					\r (将光标回到本行开头后续输入就会替换原有的字符)
					\", \'
	4. 布尔型 bool - 只允许true和false，只占用1个字节
	5. 字符串 string - 字符串不可变 - 字符串一旦定义好，其中的字符不可变 - s2[2] = 't' 报错
					如果字符串中有特别字符，不能使用双引号，而是反引号 ``
					字符串拼接 "abc" + "def" (拼接时加号保留在最后)
					或者 s5 += "abc"
	6. 基本数据类型的默认值： 整数类型：0，浮点类型：0，布尔类型：false，字符串类型：""
	7. 基本数据类型之间的转换，只能显式转换(即便式从小到大的转换) - T(v)把v转换成T类型，
					大类型到小类型的转换，不会报错，但会数据溢出，数不对了。以上常用于整形和浮点之间的互转。
					常用的数据类型转string，或者string转数据类型。
						基本类型转string类型：方式1：fmt.Springf("%参数%",表达式）；方式2：strconv函数 - strconv.Format***
						string类型转基本类型：strconv.parse***，确保string类型式有效的转换数据，否则按照基本类型的默认值输出
*/
func main2() {
	var num = 28
	fmt.Printf("num type is : %T", num)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num), " bytes")
	var num1 float32 = 3.14
	fmt.Println(num1)
	var num2 float32 = 314e+2
	var num3 float64 = 314e-2
	fmt.Println(num2, num3)
	var num4 float32 = 256.00000916
	var num5 float64 = 256.00000916
	fmt.Println(num4, num5)
	var num6 = 3.14
	fmt.Printf("%T", num6)
	fmt.Println()
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '6'
	fmt.Println(c2)
	var c3 byte = '{'
	fmt.Println(c3)
	//	汉字字符对应unicode码值，比如“中”对应20013
	var c4 uint = '中'
	fmt.Println(c4)
	var c5 byte = 'A'
	fmt.Printf("c4对应具体的字符%c", c4)
	fmt.Printf("c5对应具体的字符%c", c5)
	fmt.Println()
	fmt.Println("aaa\nb\bb")
	fmt.Println("aaa\rbbbb")
	fmt.Println("aaa\tbbb")
	var flag1 bool = true
	var flag2 bool = false
	fmt.Println(flag1, flag2)
	flag3 := 5 < 9
	fmt.Println(flag3)
	var s1 string = "hello"
	fmt.Println(s1)
	var s2 string = `fmt.Printf("c4对应具体的字符%c", c4)
fmt.Printf("c5对应具体的字符%c", c5)`
	fmt.Println(s2)
	var s3 = "abc" + "def"
	s3 += "hij"
	fmt.Println(s3)
	var nn1 int = 100
	var nn2 float32 = float32(nn1)
	fmt.Println(nn1, nn2)
	fmt.Printf("%T", nn2)
	var cn1 int = 19
	var cn2 float32 = 4.78
	var cn3 bool = false
	var cn4 byte = 'a'
	fmt.Println()
	var cs1 string = fmt.Sprintf("%d", cn1)
	fmt.Printf("cs1的类型：%T, cs1 = %q\n", cs1, cs1)
	var cs2 string = fmt.Sprintf("%f", cn2)
	fmt.Printf("cs2的类型：%T, cs2 = %q\n", cs2, cs2)
	var cs3 string = fmt.Sprintf("%t", cn3)
	fmt.Printf("cs3的类型：%T, cs3 = %q\n", cs3, cs3)
	var cs4 string = fmt.Sprintf("%c", cn4)
	fmt.Printf("cs4的类型：%T, cs4 = %q\n", cs4, cs4)
	var cn5 int = 18
	var cs5 string = strconv.FormatInt(int64(cn5), 10)
	fmt.Printf("cs5的类型：%T, cs5 = %q\n", cs5, cs5)
	var cn6 float64 = 4.29
	var cs6 string = strconv.FormatFloat(cn6, 'f', 9, 64)
	fmt.Printf("cs5的类型：%T, cs6 = %q\n", cs6, cs6)
	var cn7 bool = true
	var cs7 string = strconv.FormatBool(cn7)
	fmt.Printf("cs7的类型：%T, cs7 = %q\n", cs7, cs7)
	var sc1 = "true"
	var scb bool
	scb, _ = strconv.ParseBool(sc1)
	fmt.Printf("%T, %v\n", scb, scb)
	var sc2 string = "19"
	var sci int64
	sci, _ = strconv.ParseInt(sc2, 10, 64)
	fmt.Printf("%T, %v\n", sci, sci)
	var sc3 string = "3.14"
	var scf float64
	scf, _ = strconv.ParseFloat(sc3, 64)
	fmt.Printf("%T, %v\n", scf, scf)
}
