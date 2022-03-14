package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Unit8FromInt(-1))
	fmt.Println(Unit8FromInt(10))
	fmt.Println(Unit8FromInt(255))
	fmt.Println(Unit8FromInt(256))
}

func Unit8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the unit8 range", n)
}
