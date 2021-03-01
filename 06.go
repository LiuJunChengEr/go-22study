package main

import "fmt"

type person struct {
	name string
	age  int
}

type person1 struct {
	name string
	age  int
	addr address
}

type address struct {
	province string
	city     string
}

type Stringer interface {
	String() string
}

//func (p *person) String() string {
func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d\n", p.name, p.age)
}

func printString(s Stringer) {
	println(s.String())
}

func NewPerson(name string) *person {
	return &person{name: name}
}

type person2 struct {
	name string
	age  uint
	address
}

func main() {
	//var p person
	p := person{"xxx", 18}
	fmt.Println(p.name, p.age)

	p1 := person{age: 18, name: "xxx"}
	fmt.Println(p1.age, p1.name)

	p2 := person1{
		age:  30,
		name: "xxx",
		addr: address{
			province: "北京",
			city:     "北京",
		},
	}
	println(p2.addr.province)

	printString(person{"xxx", 18})
	printString(&person{"xxx", 18})

	println(NewPerson("yyy").name)

	p3 := person2{
		age:  30,
		name: "xxx",
		address: address{
			province: "北京",
			city:     "北京",
		},
	}
	println(p3.province)
}
