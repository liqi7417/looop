package main

import (
	"fmt"
)

var ss interface{}

func main() {
	ss = nil
	fmt.Printf("ss:%v\n", ss)
}
