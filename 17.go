package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	ss := []string{"a", "b"}
	fmt.Printf("长度：%d,容量：%d\n", len(ss), cap(ss))
	ss = append(ss, "c", "d")
	fmt.Printf("长度：%d,容量：%d\n", len(ss), cap(ss))

	a1 := [2]string{"a", "b"}
	s1 := a1[0:1]
	s2 := a1[:]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

	s := "xxx"
	b := []byte(s)
	//s3:=string(b)
	s4 := (*string)(unsafe.Pointer(&b))
	println(*s4)
}
