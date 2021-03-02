package main

import (
	"fmt"
)

func getNum() int64 {
	return 1024
}

func getNum2() int64 {
	return 500
}

func main() {
	a := getNum()
	b := getNum2()
	a_c := a/1000 + 1
	b_c := b/1000 + 1
	fmt.Println(a_c, b_c)
}
