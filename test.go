package main

import (
	"fmt"
	"math"
)

func main() {
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

	var num int = 1
	switch num {
	case 1, 2, 3:
		fmt.Println(num)
	case 4:
		fmt.Println("OK")
	case 100:
		fmt.Println(num)
	}

	switch {
	case num > 0:
		fmt.Println("num > 0")
	case num == 0:
		fmt.Println("num = 0")
	case num < 0:
		fmt.Println("num < 0")
	default:
		fmt.Println("num is not a number")
	}

	switch num2 := 10; {
	case num2 > 0:
	case num2 == 0:
	case num2 < 0:
	default:
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
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
