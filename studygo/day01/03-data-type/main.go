package main

import (
	"fmt"
	"math"
)

// 整型
// 无符号
var uint8A uint8 = 255
var uint16A uint16 = 65535
var uint32A uint32 = 4294967295
var uint64A uint64 = 18446744073709551615

// 有符号
var int8A int8 = 127
var int16A int16 = 32767
var int32A int32 = 2147483647
var int64A int64 = 9223372036854775807

// 特殊整型
/* 类型	描述
uint	32位操作系统上就是uint32，64位操作系统上就是uint64
int	32位操作系统上就是int32，64位操作系统上就是int64
uintptr	无符号整型，用于存放一个指针

注意： 在使用int和 uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异。

注意事项 获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等都可以用int来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和 uint。
*/

// 数字字面量语法（Number literals syntax）
/* Go1.13版本之后引入了数字字面量语法，这样便于开发者以二进制、八进制或十六进制浮点数的格式定义数字，例如：

v := 0b00101101， 代表二进制的 101101，相当于十进制的 45。 v := 0o377，代表八进制的 377，相当于十进制的 255。 v := 0x1p-2，代表十六进制的 1 除以 2²，也就是 0.25。

而且还允许我们用 _ 来分隔数字，比如说： v := 123_456 表示 v 的值等于 123456。

我们可以借助fmt函数来将一个整数以不同进制形式展示。
*/

// 十进制
var a int64 = 10

// 八进制 以0开头
var b int64 = 077

// 十六进制 以ox开头
var c int64 = 0xff

// 浮点型
/*
Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准： float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。 float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。

打印浮点数时，可以使用fmt包配合动词%f，代码如下：
*/

// 复数
// complex64和complex128
var c1 complex64
var c2 complex128

// 布尔值
/* Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

注意：

布尔类型变量的默认值为false。
Go 语言中不允许将整型强制转换为布尔型.
布尔型无法参与数值运算，也无法与其他类型进行转换。 */
// true
var b1 bool = true

// 字符串
/* Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：
 */
var s string = "hello"

// 字符串转义符
/* Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

转义符	含义
\r	回车符（返回行首）
\n	换行符（直接跳到下一行的同列位置）
\t	制表符
\'	单引号
\"	双引号
\\	反斜杠
举个例子，我们要打印一个Windows平台下的一个文件路径：
fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
*/

// 多行字符串
/*
Go语言中要定义一个多行字符串时，就必须使用反引号字符：
*/
var s1 string = `
第一行
第二行
第三行
`

// 字符串的常用操作
/*
方法	介绍
len(str)	求长度
+或fmt.Sprintf	拼接字符串
strings.Split	分割
strings.contains	判断是否包含
strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
strings.Index(),strings.LastIndex()	子串出现的位置
strings.Join(a[]string, sep string)	join操作
*/

// byte和rune类型
/* 组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
 */
var byteA byte = 'a'
var runeA rune = '中'

func main() {
	fmt.Println("uint8A=", uint8A)
	fmt.Println("uint16A=", uint16A)
	fmt.Println("uint32A=", uint32A)
	fmt.Println("uint64A=", uint64A)

	// 10
	fmt.Printf("a = %d \n", a)

	// 1010  占位符%b表示二进制
	fmt.Printf("a = %b \n", a)

	fmt.Printf("%o \n", b) // 77

	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	fmt.Printf("Pi = %f\n", math.Pi)
	fmt.Printf("Pi = %.2f\n", math.Pi)

	// (1+2i)
	c1 = 1 + 2i
	// (2+3i)
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(b1)

	fmt.Println("s=", s, "\ns1=", s1)

	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	fmt.Println("s的长度=", len(s))
	fmt.Println("byteA=", byteA, "runeA=", runeA)

	foreachString()
}

// 遍历字符串
/* 	因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。 */
func foreachString() {
	var s string = "hello 街墙"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}

	fmt.Println()

	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}

	fmt.Println()
}

// 修改字符串
/* 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。 */
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// 类型转换
/*
Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

强制类型转换的基本语法如下：

T(表达式)
其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。
*/
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
