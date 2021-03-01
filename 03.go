package main

import "fmt"

func main() {
	if i := 10; i > 10 {
		fmt.Println("i>10")
	} else if i < 10 {
		fmt.Println("i < 10")
	} else {
		fmt.Println("i=10")
	}

	switch j := 6; {
	case j > 10:
		fmt.Println("i>10")
	case j > 5 && j < 10:
		fmt.Println("5<i<=10")
	default:
		fmt.Println("i<=5")
	}

	switch j := 1; j {
	case 1:
		fallthrough
	case 2:
		fmt.Println("1")
	default:
		fmt.Println("没有匹配")
	}

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	println(sum)

	sum1:=0
	i:=1
	for i<100 {
		sum1+=i
		i++
	}
	println(sum1)
}
