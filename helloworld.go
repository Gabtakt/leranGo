// Package main is the main program
package main

// Package implementing formatted I/O.

// main is the beginning of the program
// func main() {
// 	fmt.Println("hello world")
// 	print("ABC")
// 	fmt.Println(add(1, 2))
// 	fmt.Println(rev(1, 2))
// 	type x int
// 	type y string
// 	var a x = 5
// 	var b y = "abc"
// 	fmt.Println(a, b)
// }

func add(a int, b int) int {
	return a + b
}

func rev(a int, b int) (c int, d int) {
	return b, a
}
