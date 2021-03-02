package main

import (
	"fmt"
	"time"
)

func sum(a, b int) int {
	t := time.Now().UnixNano()
	fmt.Printf("sum, a:%d, b:%d, t:%d\n", a, b, t)
	t = time.Now().UnixNano()
	fmt.Printf("sum, a:%d, b:%d, t:%d\n", a, b, t)
	t = time.Now().UnixNano()
	fmt.Printf("sum, a:%d, b:%d, t:%d\n", a, b, t)
	t = time.Now().UnixNano()
	fmt.Printf("sum, a:%d, b:%d, t:%d\n", a, b, t)
	t = time.Now().UnixNano()
	fmt.Printf("sum, a:%d, b:%d, t:%d\n", a, b, t)
	return a + b
}

func cal(a, b int) int {
	defer func(a, b int) {
		t := time.Now().UnixNano()
		fmt.Printf("cal defer, a:%d, b:%d, t:%d\n", a, b, t)
	}(a, b)
	t := time.Now().UnixNano()
	fmt.Printf("cal, a+b:%d, t:%d\n", a+b, t)
	return sum(a, b)
}

func main() {
	num := cal(3, 2)
	t := time.Now().UnixNano()
	fmt.Printf("num:%d, t:%d\n", num, t)
}
