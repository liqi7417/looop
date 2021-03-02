package main

import (
	"time"
)

type people struct {
	name string
}

func main() {
	p := &people{
		name: "test_people",
	}
	go func(p *people) {
		name_go := p.name
		print(name_go)
	}(p)
	//p.name = "nihao"
	print(p.name)
	time.Sleep(5 * time.Second)
}
