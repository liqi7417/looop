package main

import (
	"fmt"
	"github.com/liqi7417/looop/src/mocktest/service"
	"time"
)

func main() {
	fmt.Println("vim-go")
	var s service.SI
	s = service.NewRPCService()
	s.Run()
	fmt.Println("sleep 3 sec")
	time.Sleep(3 * time.Second)
}
