/**
	定义了包名，必须在源文件中非注释的第一行指明这个文件属于哪个包，
	package main 表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包
 */
package main

/**
	import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数或其他元素），fmt 包实现了格式化IO的函数。
 */
import "fmt"
import "math"
import "runtime"

/**
	func main() 是程序开始执行的函数，main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数。
	如果有 init() 函数，则会先执行该函数。
 */

 var a, b int

 var (
 	c int
 	d bool
 )

 var e, f int = 1, 2
 var g, h = 123, "hello"

 func max(x, y int) int {
 	if x > y {
 		return x
	} else {
		return y
	}
 }

 func swap(x *int, y *int) {
 	var temp int
 	temp = *x
 	*x = *y
 	*y = temp
 }

 /**
 	定义结构体 Circle
  */
 type Circle struct {
 	radius float64
 }

 func (c Circle) getArea() float64 {
 	return 3.14 * c.radius * c.radius
 }

 func Factorial(n uint64) (result uint64) {
 	if n > 0 {
 		return n * Factorial(n - 1)
	}
	return 1
 }

 func Fibonacci(n int) int {
 	if n < 2 {
 		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
 }

 func say(s string) {
 	for i:= 0; i < 5; i++ {
 		// runtime.Gosched() 表示让 CPU 把时间片让给别人，下次某个时候继续恢复执行该 goroutine
 		runtime.Gosched()
 		fmt.Println(s)
	}
 }

 func xsum(a []int, c chan int) {
 	total := 0
 	for _, v := range a {
 		total += v
	}
	c <- total
 }



func main() {
	// fmt.Println() 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
	fmt.Println("Hello World!")
	fmt.Print("Hello World!\n")

	var x = "啦啦啦"
	var y string = "xdhuxc"
	// bool 的默认值为：false
	var z bool
	/**
		变量的类型将由编译器自动推断。
		出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误。
		只能在函数体中出现，不可以用于全局变量的声明和赋值。
	 */
	t := 20

	fmt.Println(a, b, c, d, e, f, g, h)
	fmt.Println(x, y, z, t)

	/**
		空白标识符 _ 也被用于抛弃值，例如：值 5 在 _, b = 5, 7 中被抛弃。
		_ 实际上是一个只写变量，不能得到它的值，这样做是因为 Go 语言中必须使用所有被声明的变量，但有时并不需要使用从一个函数中得到的所有返回值。
	 */

	 const (
	 	a1 = iota
	 	a2
	 	a3
	 	a4 = "ha"
	 	a5
	 	a6 = 100
	 	a7
	 	a8 = iota
	 	a9
	 )
	 // 0 1 2 ha ha 100 100 7 8
	 fmt.Println(a1, a2, a3, a4, a5, a6, a7, a8, a9)

	 const (
	 	i = 1 << iota // i = 1 << 0
	 	j = 3 << iota // j = 3 << 1 = 3 * 2^1
	 	k             // k = 3 << 2 = 3 * 2^2
	 	l             // l = 3 << 3 = 3 * 2^3
	 )
	 fmt.Println(i, j, k, l)


	 a = 0
	 b = 15
	 numbers := [6]int{1, 2, 3, 5}

	 for a = 0; a < 10; a++ {
	 	fmt.Printf("a 的值为：%d\n", a)
	 }

	 for a < b {
	 	a++
	 	fmt.Printf("a 的值为：%d\n", a)
	 }

	 for m, n := range numbers {
	 	fmt.Printf("第 %d 位 x 的值为：%d\n", m, n)
	 }
	 var p, q int

	 for p = 2; p < 100; p++ {
	 	for q = 2; q <= (p/q); q++ {
	 		if p % q == 0 {
	 			break
			}
		}
		if q > (p/q) {
			fmt.Printf("%d 是素数。\n", p)
		}
	 }

	 s := 100
	 t = 200
	 fmt.Printf("交换前，s 的值为：%d\n", s)
	 fmt.Printf("交换前，k 的值为：%d\n", t)

	 swap(&s, &t)

	fmt.Printf("交换前，s 的值为：%d\n", s)
	fmt.Printf("交换前，k 的值为：%d\n", t)

	// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	// 使用函数
	fmt.Println(getSquareRoot(9))

	var c1 Circle
	c1.radius = 10.0
	fmt.Println("c1 的面积为：", c1.getArea())

	// 声明一维数组 n
	var n [10] int
	for i := 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j := 0; j < 10; j++ {
		fmt.Printf("%d\n", n[j])
		fmt.Printf("%x\n", &n[j])
	}

	var ptr *float64
	fmt.Printf("ptr 的值为：%x\n", ptr)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s ---> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Printf("%d ---> %c\n", i, c)
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	// 递归函数
	var r uint64 = 15
	fmt.Printf("%d 的阶乘是：%d\n", r, Factorial(r))

	for e := 1; e < 10; e++ {
		fmt.Printf("%d\t", Fibonacci(e))
	}

	// 开一个新的 Goroutine 执行
	go say("World")
	// 当前 Goroutine 执行
	say("hello")

	f := []int{1, 4, 6, 12, 57, -9}
	d1 := make(chan int)
	go xsum(f[:len(f)/2], d1)
	go xsum(f[len(f)/2:], d1)
	c2, c3 := <-d1, <-d1
	fmt.Println(c2, c3, c2 + c3)

	e1 := make(chan int, 2)
	e1 <- 1
	e1 <- 2
	e1 <- 3
	fmt.Println(<-e1)
	fmt.Println(<-e1)
	fmt.Println(<-e1)
}

/**
	当标识符（包括常量、变量、类型、函数名、结构字段等）以一个大写字母开头，如，Xdhuxc，那么使用这种形式的标识符的对象就可以被外部包的代码所使用，客户端程序需要先导入这个包，这被称为导出，
	标识符如果以小写字母开头，则对包外是不可见的，但是它们在整个包的内部是可见并且可用的。
 */

