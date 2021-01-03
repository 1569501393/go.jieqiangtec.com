package main

import "fmt"

// 标准变量声明
// var 变量名 变量类型
// 变量声明以关键字var开头，变量类型放在变量的后面，行尾无需封号。

// var name string
// var age int
var isOK bool

/*
批量声明
每声明一个变量就需要写var关键字比较繁琐,go 语言中还支持批量声明 */
var (
	a string
	b int
	c bool
	d float32
)

/* 变量初始化
go 语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。
每个变量会被初始化成其类型的默认值，例如：
整型和浮点型变量的默认值为 0
字符串变量的默认值为 空字符串
布尔型变量默认值为 false
切片 函数 指针 变量的默认值为 nil

当然我们也可以在声明变量的时候为其指定初始值。
变量的初始化的标准格式如下：
var 变量名 类型 = 表达式
*/

// var name string = 'jieqiang';
// var age int = 18

// 或者一次初始化多个变量
/* 单引号，双引号
在go语法中，双引号是常用的来表达字符串，如果你使用了单引号，编译器会提示出错
invalid character literal (more than one character) */
// var name, age = 'jieqiang', 20
// var name, age = "Q1mi", 20

/*
类型推导
有时候我们会将变量的类型省略，这个时候的编译器会根据等号右边的值来推导变量的类型完成初始化
*/
var name = "jieqiang"
var age = 18

/*
短变量声明
在函数内部，可以使用更加简陋的 := 方式声明并初始化变量
*/
// 全局变量
var m = 100

/*
匿名变量
在使用多重赋值时，如果想要忽略某个值，可以使用 匿名变量（anonymous variable）.
匿名变量用一个下划线 _ 表示，例如
*/
func foo() (int, string) {
	return 10, "Q1mi"
}

/*
// 常量
相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。
常量的声明和变量的声明非常类似，只是把 var 换成 const, 常量在定义的时候必须赋值。
*/
// const pi = 3.1415
// const e = 2.7182

/*
声明了pi和e 这两个常量之后，在整个程序运行期间他们的值都不再发生变化了。
多个常量也可以一起声明。
*/
const (
	pi = 3.1415
	e  = 2.7182
)

/*
const 同时声明多个常量时， 如果省略了值，则表示和上面一行的值相同。例如：
*/

/* const (
	n1 = 100
	n2
	n3
) */

// 上面示例中，常量n1 n2 n3的值都是100
/*
iota 是go语言的常量计数器，只能在常量的表达式中使用。
iota 在const关键字出现时将被重置为0。
const中每新增一行常量声明将使 iota 计数一次
（iota 可理解为const语句块中的行索引）。
使用 iota 能简化定义，在定义枚举时很有用。
*/
/* const (
	// 0
	n1 = iota
	// 1
	n2
	// 2
	n3
	// 3
	n4
) */

// 几个常见的iota示例:
// 使用 _ 跳过某些值
/* const (
	// 0
	n1 = iota
	// 1
	n2
	// 2
	_
	// 3
	n4
) */

// iota声明中间插队
const (
	// 0
	n1 = iota
	// 100
	n2 = 100
	// 2
	n3 = iota
	// 3
	n4
)

const n5 = iota

/* 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
 */

const (
	_ = iota
	// KB ...
	KB = 1 << (10 * iota)
	// MB ...
	MB = 1 << (10 * iota)
	// GB ...
	GB = 1 << (10 * iota)
	// TB ...
	TB = 1 << (10 * iota)
	// PB ...
	PB = 1 << (10 * iota)
)

// 多个iota定义在一行
/* const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
) */

func main() {
	n := 10
	m := 200
	// 200 10
	fmt.Println(m, n)
	fmt.Println(name)
	fmt.Println(age)

	x, _ := foo()
	_, y := foo()

	// x= 10
	fmt.Println("x=", x)
	// y= Q1mi
	fmt.Println("y=", y)

	// n1= 100 , n2= 100 , n3= 100
	fmt.Println("n1=", n1, ", n2=", n2, ", n4=", n4)
	fmt.Println("KB=", KB, ",MB=", MB)
}
