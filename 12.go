package main

import "fmt"

func main() {
	name := "xxx"
	nameP := &name
	//var nameP *string = &name
	fmt.Println(name, nameP)

	intP := new(int)
	println(intP)

	nameV := *nameP
	println("nameP指针指向的值为：", nameV)

	*nameP = "yyy"
	println("修改后nameP指针指向的值为：", *nameP)
	println("修改指针后name的值为：", name)

	age := 18
	modifyAge(&age)
	println(age)
}

func modifyAge(age *int) {
	*age = 20
}
