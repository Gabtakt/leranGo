package main

import "fmt"

const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009
const Log2E = 1 / Ln2 // this is a precise reciprocal
const Billion = 1e9   // float constant
const hardEight = (1 << 100) >> 97

const (
	a = iota // a = 0
	b        // b = 1
	c        // c = 2
	d = 5    // d = 5
	e        // e = 5
)

// 赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
const (
	Apple, Banana     = iota + 1, iota + 2 // Apple=1 Banana=2
	Cherimoya, Durian                      // Cherimoya=2 Durian=3
	Elderberry, Fig                        // Elderberry=3, Fig=4

)

var n int

func P(a int) {
	fmt.Println(a)
}

// func main() {
// 	P(a)
// 	P(b)
// 	P(c)
// 	P(d)
// 	P(e)
// }
