package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s_origin interface{}

	s_origin = "tokenerror"

	ss, _ := json.Marshal(s_origin)

	ss_now := string(ss)
	fmt.Println(ss_now)
	fmt.Println(s_origin)
}
