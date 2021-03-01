package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	ip := &i
	// var fp *float64 = (*float64)(ip) 这是不可以的
	fmt.Println(ip)

	fp := (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	println(i)

	main1()

	println(unsafe.Sizeof(int8(0)))
}

func main1() {
	p := new(person)
	pName := (*string)(unsafe.Pointer(p))
	*pName = "xxx"

	pAge := (*int)(unsafe.Pointer((uintptr(unsafe.Pointer(p))) + unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)
}

type person struct {
	Name string
	Age  int
}
