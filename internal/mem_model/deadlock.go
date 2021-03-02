package main

import (
    "runtime"
    //"time"
)

type T struct {
    msg string
}

var t *T

func setup() {
    //time.Sleep(3 * time.Second)
    print("setup\n")
    print("setup\n")
    print("setup\n")
    print("setup\n")
    runtime.Gosched()
    //time.Sleep(1 * time.Second)
    t = &T{
        msg: "start",
    }
}

func main() {
    go setup()
    var i int = 0
    for t == nil {
        i++
        print(i)
        print("\n")
        //runtime.Gosched()
    }
    print("test nil done, ", i, "\n")
    print(t.msg)
}
