package main

import (
	"errors"
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func sum1(a, b int) int {
	return a + b
}

func sum2(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或b不能是负数")
	}
	return a + b, nil
}

func sum3(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a or b can not is negative number.")
	}
	sum = a + b
	err = nil
	return
}

func sum4(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Age uint

func (age Age) String() {
	println("this age is ", age)
}

func (age *Age) Modify() {
	*age = Age(30)
}

func main() {
	println(sum(1, 2))
	println(sum1(1, 2))
	i, err := sum2(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	result, _ := sum2(-1, 2)
	println(result)

	println(sum4(1, 2, 3, 4, 5))

	sum2 := func(a, b int) int {
		return a + b
	}
	println(sum2(1, 2))

	cl := colsure()
	println(cl())
	println(cl())
	println(cl())

	age := Age(18)
	age.String()
	age.Modify()
	//(&age).Modify()
	age.String()
}
