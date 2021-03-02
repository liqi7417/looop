package main

import (
	"fmt"
	"reflect"
	"time"
)

type FACC struct {
	slit string
}

func main() {
	fc := make(map[string]interface{})
	ak, exist := fc["testkey"]
	as, ok := ak.(string)
	fmt.Println("%v %v %v %v", ak, exist, as, ok)
	fmt.Println(time.Now().UTC())
	a := time.Now().UTC().AddDate(0, 0, 0).Format("20060102")
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)
}
