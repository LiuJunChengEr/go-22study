package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var i = 10
	println(i)

	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Printf("f32 is %f,f64 is %f\n", f32, f64)

	var bf bool = false
	var bt bool = true
	fmt.Printf("bf is %t,bt is %t\n", bf, bt)

	var s1 string = "hello"
	var s2 string = "世界"
	fmt.Printf("s1 is %s,s2 is %s\n", s1, s2)
	fmt.Println(s1 + s2)

	j := 10
	println(j)

	pj := &j
	println(pj)

	const name = "xxx"
	println(name)

	//const (
	//	one = 1
	//	two = 2
	//	three = 3
	//	four = 4
	//)

	const (
		one = iota + 1
		two
		three
		four
	)

	itoa := strconv.Itoa(j)
	atoi, err := strconv.Atoi(itoa)
	fmt.Println(itoa, atoi, err)

	j2f := float64(j)
	f2i := int(j2f)
	fmt.Println(j2f, f2i)

	str:= "hello 你好"
	println(str,len(str))
	println(utf8.RuneCountInString(str))
	println(len([]rune(str)))
}
