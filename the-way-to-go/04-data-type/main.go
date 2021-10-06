package main

import (
	fm "fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)
import "fmt"

func init() {
	println("begin===")
}

func main() {
	println("Hello, World")
	fmt.Println("Hello, world222")
	alias()
	var sum int
	sum = Sum(1, 2)
	fmt.Println(sum)

	// 基础数据类型
	BasicType()
	ConstType()
	VarType()
	StringFunctions()
	TimeFunctions()
	PointFunctions()
}

func alias() {
	fm.Println("Hello, Alias")
}

func Sum(a, b int) int {
	return a + b
}

func BasicType() {
	a := 5.0
	b := int(a)
	fmt.Println("b==", b)
}

func ConstType() {
	const PI = 3.141592654
	// fmt.Sprintf("Pi=%v", PI)
	fmt.Println("Pi=", PI)

	// 在 Go 语言中，你可以省略类型说明符 [type]，因为编译器可以根据常量的值来推断其类型。
	// 显式类型定义： const b string = "abc"
	// 隐式类型定义： const b = "abc"
	const str string = "abc"
	fmt.Println(fmt.Sprintf("str= %s, %T, %v", str, str, str))

	// 常量也允许使用并行赋值的形式
	// const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday = 1, 2, 3, 4, 5, 6, 7
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	println("Sunday=", Sunday)
	println("Monday=", Monday)
	println("Tuesday=", Tuesday)
	println("Wednesday=", Wednesday)
	println("Thursday=", Thursday)
	println("Friday=", Friday)
	println("Saturday=", Saturday)

	const (
		unKnown = 0
		Female  = 1
		Male    = 2
	)
	println("UnKnown=", unKnown)

	// iota 可以被用作枚举值
	const (
		ZERO = iota
		ONE
		TWO
	)
	println("ZERO=", ZERO)
	println("ONE=", ONE)
	println("TWO=", TWO)
}

func VarType() {
	// var a int
	// var b bool
	// var str string

	var (
		a   int
		b   bool
		str string
	)

	println("a=", a, reflect.TypeOf(a))
	println("b=", b, reflect.TypeOf(b))
	println("str=", str, reflect.TypeOf(str))

	fmt.Printf("version=%s \n", runtime.Version())

	// 格式化说明符
	// 在格式化字符串里，%d 用于格式化整数（%x 和 %X 用于格式化 16 进制表示的数字），%g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法），%0d 用于规定输出定长的整数，其中开头的数字 0 是必须的。
	// %n.mg 用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 或者 f，例如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00。
	// ————————————————
	// 原文作者：Go &#25216;&#26415;&#35770;&#22363;&#25991;&#26723;&#65306;&#12298;Go &#20837;&#38376;&#25351;&#21335;&#65288;&#65289;&#12299;
	// 转自链接：https://learnku.com/docs/the-way-to-go/basic-types-and-operators/3586
	// &#29256;&#26435;&#22768;&#26126;&#65306;&#32763;&#35793;&#25991;&#26723;&#33879;&#20316;&#26435;&#24402;&#35793;&#32773;&#21644; LearnKu &#31038;&#21306;&#25152;&#26377;&#12290;&#36716;&#36733;&#35831;&#20445;&#30041;&#21407;&#25991;&#38142;&#25509;

	// 位左移 <<：
	// 用法：bitP << n。
	// bitP 的位向左移动 n 位，右侧空白部分使用 0 填充；如果 n 等于 2，则结果是 2 的相应倍数，即 2 的 n 次方。例如：
	fmt.Println("1 << 10=", 1<<10)
	fmt.Println("1 << 20=", 1<<20)

	// 位左移常见实现存储单位的用例
	// 使用位左移与 iota 计数配合可优雅地实现存储单位的常量枚举：
	type ByteSize float64
	const (
		_           = iota // 通过赋值给空白标识符来忽略值
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		XB
		YB
	)
	fmt.Println("KB=", KB)
	fmt.Println("YB=", YB)

	// 在通讯中使用位左移表示标识的用例
	type BitFlag int
	const (
		Active BitFlag = 1 << iota // 1 << 0 == 1
		Send
		Receive
	)

	flag := Active | Send
	println("flag=", flag)

	// 04-data-type/main.go:149:12: division by zero
	// println(1 / 0)

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("a=%d / \n", a)
	}

	for i := 0; i < 10; i++ {
		r := rand.Intn(8)
		fmt.Printf("r=%d / \n", r)
	}

	timeNs := int64(time.Now().Nanosecond())
	rand.Seed(timeNs)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f/ \n", 100*rand.Float32())
	}

	// 字符串
	// 解释字符串：
	// 该类字符串使用双引号括起来，其中的相关的转义字符将被替换，这些转义字符包括：
	// \n：换行符
	// \r：回车符
	// \t：tab 键
	// \u 或 \U：Unicode 字符
	// \\：反斜杠自身
	//
	//
	// 非解释字符串：
	// 该类字符串使用反引号括起来，支持换行，例如：
	//
	// ————————————————
	// 原文作者：Go &#25216;&#26415;&#35770;&#22363;&#25991;&#26723;&#65306;&#12298;Go &#20837;&#38376;&#25351;&#21335;&#65288;&#65289;&#12299;
	// 转自链接：https://learnku.com/docs/the-way-to-go/character-string/3587
	// &#29256;&#26435;&#22768;&#26126;&#65306;&#32763;&#35793;&#25991;&#26723;&#33879;&#20316;&#26435;&#24402;&#35793;&#32773;&#21644; LearnKu &#31038;&#21306;&#25152;&#26377;&#12290;&#36716;&#36733;&#35831;&#20445;&#30041;&#21407;&#25991;&#38142;&#25509;

	// var strRaw = `this is a raw thing \n`
	strRaw := `this is a raw thing \n`
	fmt.Printf("原样输出：%s \n", strRaw)
	fmt.Printf("字符串 strRaw 的第 1 个字节：%v \n", strRaw[0])
	fmt.Printf("字符串 strRaw 的最后 1 个字节：%v \n", strRaw[len(strRaw)-1])

	// 字符串拼接符 +
	// 	两个字符串 s1 和 s2 可以通过 s := s1 + s2 拼接在一起。
	// s2 追加在 s1 尾部并生成一个新的字符串 s。
	// 你可以通过以下方式来对代码中多行的字符串进行拼接：

	strConcat := "Beginning of the string" + ", second part of the thing"
	println("strConcat=", strConcat)
}

func StringFunctions() {
	// hasPrefixHello:true, type:bool, value:true, value:true
	// prefix := "Hello"

	// hasPrefixHello:false, type:bool, value:false, value:false
	prefix := "hello"
	str := "Hello, world"
	hasPrefixHello := strings.HasPrefix(str, prefix)
	// hasPrefixHello:false, type:bool, value:false, value:false
	fmt.Printf("hasPrefixHello:%v, type:%T, value:%#v, value:%+v \n", hasPrefixHello, hasPrefixHello, hasPrefixHello, hasPrefixHello)

	substr := "H"
	strContains := strings.Contains(str, substr)
	fmt.Printf("%s strContains %s :%v, type:%T, value:%#v, value:%+v \n", str, substr, strContains, strContains, strContains, strContains)

	fmt.Printf("The position of the first instance '%s' of '%s' is:%v \n", substr, str, strings.Index(str, substr))
	fmt.Printf("The position of the Last instance '%s' of '%s' is:%v \n", substr, str, strings.LastIndex(str, substr))
	fmt.Printf("'%s' Replace '%s' to 'h' is:%v \n", str, substr, strings.Replace(str, substr, "h", -1))
	fmt.Printf("'%s' Count '%s'  is:%v \n", str, substr, strings.Count(str, substr))

	repeatCnt := 3
	fmt.Printf("'%s' Repeat %d  is:%#v \n", str, repeatCnt, strings.Repeat(str, repeatCnt))
	fmt.Printf("'%s' ToLower is:%#v \n", str, strings.ToLower(str))
	fmt.Printf("'%s' ToUpper is:%#v \n", str, strings.ToUpper(str))
	fmt.Printf("'%s' TrimSpace is:%#v \n", str, strings.TrimSpace(str))
	fmt.Printf("'%s' Fields is:%#v \n", str, strings.Fields(str))
	fmt.Printf("'%s' Split is:%v, type:%T, value:%#v, value:%+v \n", str, strings.Split(str, ","), strings.Split(str, ","), strings.Split(str, ","), strings.Split(str, ","))

	i := 99
	strI := "99"
	fmt.Printf("'%d' Itoa is:%#v \n", i, strconv.Itoa(i))

	atoi, err := strconv.Atoi(strI)
	if err != nil {
		fmt.Printf("'%s' Atoi is:%#v \n", strI, err)
		return
	}
	fmt.Printf("'%s' Atoi is:%#v \n", strI, atoi)
	fmt.Printf("The size of ints is:%+v \n", strconv.IntSize)
}

func TimeFunctions() {
	fmt.Printf("time.Now() is %v \n", time.Now())
	fmt.Printf("time.Now().Day() is %v \n", time.Now().Day())
	fmt.Printf("time.Now().Format() is %v \n", time.Now().Format("20060102"))
	fmt.Printf("time.Now().Format() is %v \n", time.Now().Format("2006-01-02 03:04:05"))
}

func PointFunctions() {
	i := 5
	fmt.Printf("An Integer: %d, it's location in memory: %p \n", i, &i)
	fmt.Printf("An Integer: %d, it's location in memory: %#v \n", i, &i)

	// var intP *int
	// intP = &i
	intP := &i
	fmt.Printf("intP: %v,type:%T, it's location in memory: %#v : %+v: , The value at memory location %p is %+v \n", intP, intP, &intP, &intP, &intP, *intP)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p, addres:%v \n", p, &p)   // prints address	0xc000056030
	fmt.Printf("Here is the string *p: %s, addres:%v \n", *p, &*p) // prints string	ciao
	fmt.Printf("Here is the string s: %s, addres:%v \n", s, &s)    // prints same string	ciao

	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x1 addr=0x0 pc=0xe1da87]
	// 	var pStr *int = nil
	// *pStr = 0
}
