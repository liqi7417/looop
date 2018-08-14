package unsafetest

import (
    "fmt"
    "testing"
    "unsafe"
)

type Person struct {
    Age   int32  //年龄
    Name  string //名字
    Hobit string //爱好
}

func TestMytest(t *testing.T) {
    var personA Person
    var a byte
    //Alignof 示例
    align := unsafe.Alignof(a)
    fmt.Println("align is: ", align)
    //Sizeof 示例
    a_size := unsafe.Sizeof(a)
    fmt.Println("a size is: ", a_size)
    //Sizeof 示例
    p_size := unsafe.Sizeof(&personA)
    fmt.Println("personA size is: ", p_size)
    //Offsetof 示例
    offset := unsafe.Offsetof(personA.Name)
    fmt.Println("offset is : ", offset)

    //Pointer 示例
    nameaddr := (uintptr)(unsafe.Pointer(&(personA.Name)))
    fmt.Println("nameaddr is: ", nameaddr)
    personAaddr := (uintptr)(unsafe.Pointer(&(personA)))
    fmt.Println("nameaddr is: ", personAaddr)
    personA.Age = 100
    personA.Hobit = "run"

    //指针操作 示例
    personAaddr2 := nameaddr - offset
    fmt.Println("personAaddr2 is: ", personAaddr2)
    personB := (*Person)(unsafe.Pointer(personAaddr2))
    fmt.Println("personB.Age is :", personB.Age)

    //异常情况 示例
    //... 中间逻辑使personAaddr2指向不合法位置
    personB = (*Person)(unsafe.Pointer(uintptr(0)))
    fmt.Println("personB.Age is :", personB.Age)
}
