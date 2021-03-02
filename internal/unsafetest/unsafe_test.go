package unsafetest

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Person struct {
	Name  string //名字
	Hobit string //爱好
	Age   int8   //年龄
}

// 字节对齐

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
	p_size := unsafe.Sizeof(personA)
	fmt.Println("personA size is: ", p_size)
	//Sizeof 示例
	p_age_size := unsafe.Sizeof(personA.Age)
	fmt.Println("personA Age size is: ", p_age_size)
	//Offsetof 示例
	name_offset := unsafe.Offsetof(personA.Name)
	fmt.Println("Name offset is : ", name_offset)
	//Offsetof 示例
	age_offset := unsafe.Offsetof(personA.Age)
	fmt.Println("Age offset is : ", age_offset)

	fmt.Println("Sizeof struct : ", unsafe.Sizeof(struct {
		i8  int8
		i64 int64
		i16 int16
		s   string
	}{}))

	//Pointer 示例
	nameaddr := (uintptr)(unsafe.Pointer(&(personA.Name)))
	fmt.Println("nameaddr is: ", nameaddr)
	personAaddr := (uintptr)(unsafe.Pointer(&(personA)))
	fmt.Println("personAaddr is: ", personAaddr)
	personA.Age = 100
	personA.Hobit = "run"

	//指针操作 示例  //错误
	personAaddr2 := nameaddr - name_offset
	fmt.Println("personAaddr2 is: ", personAaddr2)
	personB := (*Person)(unsafe.Pointer(personAaddr2))
	fmt.Println("personB.Age is :", personB.Age)

	//异常情况 示例
	//... 中间逻辑使personAaddr2指向不合法位置
	//personB = (*Person)(unsafe.Pointer(uintptr(0)))
	//fmt.Println("personB.Age is :", personB.Age)
}

func TestPointer(t *testing.T) {
	a := 5
	fmt.Println("a, type:", reflect.TypeOf(a), ", value:", reflect.ValueOf(a))
	b := (*int)(unsafe.Pointer(&a))
	fmt.Println("aaddr:", &a, ", baddr:", b, ", bvalue:", *b)
}
