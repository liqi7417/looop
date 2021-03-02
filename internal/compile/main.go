package main

import "fmt"

const s = "Go101.org"

func main() {
	var a byte = 1 << len(s) / 128
	var b byte = 1 << len(s[:]) / 128
	fmt.Println("vim-go")
	println(a, b)
}
