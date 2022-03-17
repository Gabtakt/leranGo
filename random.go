package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main2() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}
	for i := 0; i < 10; i++ {
		r := rand.Intn(8)
		fmt.Println(r)
	}
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int())
	}
}
