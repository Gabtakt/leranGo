package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

func main() {

	/*
		data-type test
	*/

	// fmt.Println(Unit8FromInt(-1))
	// fmt.Println(Unit8FromInt(10))
	// fmt.Println(Unit8FromInt(255))
	// fmt.Println(Unit8FromInt(256))
	// fmt.Println(IntFromFloat64(float64(1.5)))
	// aVar := uint8(15)
	// bVar := uint8(7)
	// fmt.Println(aVar &^ bVar)
	// fmt.Println(^uint8(255))
	// var ch int = '\u0041'
	// var ch2 int = '\u03B2'
	// var ch3 int = '\U00101234'
	// fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	// fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	// fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	// fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point

	/*
		switch test
	*/

	// var num int = 1
	// switch num {
	// case 1, 2, 3:
	// 	fmt.Println(num)
	// case 4:
	// 	fmt.Println("OK")
	// case 100:
	// 	fmt.Println(num)
	// }

	// switch {
	// case num > 0:
	// 	fmt.Println("num > 0")
	// case num == 0:
	// 	fmt.Println("num = 0")
	// case num < 0:
	// 	fmt.Println("num < 0")
	// default:
	// 	fmt.Println("num is not a number")
	// }

	// switch num2 := 10; {
	// case num2 > 0:
	// case num2 == 0:
	// case num2 < 0:
	// default:
	// }

	// k := 6
	// switch k {
	// case 4:
	// 	fmt.Println("was <= 4")
	// 	fallthrough
	// case 5:
	// 	fmt.Println("was <= 5")
	// 	fallthrough
	// case 6:
	// 	fmt.Println("was <= 6")
	// 	fallthrough
	// case 7:
	// 	fmt.Println("was <= 7")
	// 	fallthrough
	// case 8:
	// 	fmt.Println("was <= 8")
	// 	fallthrough
	// default:
	// 	fmt.Println("default case")
	// }

	/*
		for-loop test
	*/

	// for i := 1; i <= 15; i++ {
	// 	fmt.Println(i)
	// }

	// 	i := 1
	// LOOP:
	// 	fmt.Println(i)
	// 	i += 1
	// 	if i <= 15 {
	// 		goto LOOP
	// 	}

	// for i := 0; i < 25; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("G")
	// 	}
	// 	fmt.Println("")
	// }

	// for i, str := 1, "G"; i <= 25; i++ {
	// 	fmt.Println(str)
	// 	str += "G"
	// }

	// for i := 1; i <= 100; i++ {
	// 	switch {
	// 	case i%3 == 0 && i%5 == 0:
	// 		fmt.Printf("FizzBuzz ")
	// 	case i%3 == 0:
	// 		fmt.Printf("Fizz ")
	// 	case i%5 == 0:
	// 		fmt.Printf("Buzz ")
	// 	default:
	// 		fmt.Printf("%d ", i)
	// 	}
	// }

	// var i int = 5
	// for i > 0 {
	// 	i--
	// 	fmt.Printf("%d ", i)
	// }

	/*
		func test
	*/

	// num2, num3 := getX2AndX3(2)
	// fmt.Printf("func call : %d, %d\n", num2, num3)
	// num2, num3 = getX2AndX3_2(3)
	// fmt.Printf("func call : %d, %d\n", num2, num3)

	// A, B := 10, 5
	// ans1, ans2, ans3 := addMulSub_1(A, B)
	// fmt.Printf("ans: %d %d %d\n", ans1, ans2, ans3)
	// ans1, ans2, ans3 = addMulSub_2(A, B)
	// fmt.Printf("ans: %d %d %d\n", ans1, ans2, ans3)

	// var input float64 = 5
	// ans, err := MySqrt_2(input)
	// if err != nil {
	// 	fmt.Printf("Error: %s\n", err)
	// } else {
	// 	fmt.Printf("The squre root of %f is %.16f\n", input, ans)
	// }

	// n := 0
	// reply := &n
	// Multiply(10, 5, reply)
	// fmt.Println("Multiply:", n)
	// Multiply(10, 6, &n)
	// fmt.Println("Multiply:", n)

	// x := Mymin(1, 2, 3, 4, 0, -1)
	// fmt.Println("min", x)
	// y := []int{1, 2, 3, 4, 0, -2}
	// x = Mymin(y...)
	// fmt.Println("min", x)

	// F1("A", "B", "C")

	// deferTestFunction1()
	// deferTestFunction2()

	/*
		内置函数
			名称					说明
			close					用于管道通信
			len、cap				len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道，不能用于 map）
			new、make				new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）new() 是一个函数，不要忘记它的括号
			copy、append			用于复制和连接切片
			panic、recover			两者均用于错误处理机制
			print、println			底层打印函数（详见第 4.2 节），在部署环境中建议使用 fmt 包
			complex、real imag		用于创建和操作复数（详见第 4.5.2.2 节）
	*/

	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("fibonacci(%d) is: %v\n", i, fibonacci(i))
	// }

	// for i := 0; i <= 10; i++ {
	// 	res1, res2 := fibonacci2(i)
	// 	fmt.Printf("fibonacci(%d) is: %v %v\n", i, res1, res2)
	// }

	// print1To10(10)

	for i := 0; i <= 30; i++ {
		input := big.NewInt(int64(i))
		output := getFact(input)
		fmt.Printf("fact(%d) is: %v \n", i, output.String())
	}
}

func Unit8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the unit8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

func getX2AndX3(input int) (int, int) {
	return input * 2, input * 3
}

func getX2AndX3_2(input int) (x1 int, x2 int) {
	x1 = input * 2
	x2 = input * 3
	return
}

func addMulSub_1(A, B int) (int, int, int) {
	return A + B, A * B, A - B
}

func addMulSub_2(A, B int) (sum int, mul int, sub int) {
	sum = A + B
	mul = A * B
	sub = A - B
	return
}

func MySqrt_1(input float64) (float64, error) {
	if input < 0 {
		return math.NaN(), errors.New("squre root of negative number")
	}
	return math.Sqrt(input), nil
}

func MySqrt_2(input float64) (output float64, err error) {
	if input < 0 {
		output = math.NaN()
		err = errors.New("squre root of negative number")
	} else {
		output = math.Sqrt(input)
		err = nil
	}
	return
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func Mymin(arg ...int) int {
	if len(arg) == 0 {
		return 0
	}
	min := arg[0]
	for _, v := range arg {
		if v < min {
			min = v
		}
	}
	return min
}

func F1(s ...string) {
	F3(s)
	F2(s...)
	F3(s)
}

func F2(s ...string) {
	for i := 0; i < len(s); i++ {
		s[i] = "X"
	}
}

func F3(s []string) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}

func deferTestFunction1() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func deferTestFunction2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	return
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func fibonacci2(n int) (res1 int, res2 int) {
	if n <= 1 {
		res2 = 1
	} else if n == 2 {
		res2 = 2
	} else {
		_, fib1 := fibonacci2(n - 1)
		_, fib2 := fibonacci2(n - 2)
		res2 = fib1 + fib2
	}
	res1 = n
	return res1, res2
}

func print1To10(n int) {
	if n > 0 {
		fmt.Print(n, " ")
		print1To10(n - 1)
	}
}

func getFact(n *big.Int) (res *big.Int) {
	if n.Cmp(big.NewInt(0)) == 0 {
		res = big.NewInt(1)
	} else {
		tmp := new(big.Int)
		tmp.Add(tmp, n)
		tmp.Sub(tmp, big.NewInt(1))
		tmp2 := getFact(tmp)
		res = tmp.Mul(n, tmp2)
	}
	return res
}
