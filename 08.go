package main

import "fmt"

func main() {
	//go fmt.Println("xxx")
	//fmt.Println("main goroutine")
	//time.Sleep(time.Second)

	ch := make(chan string)
	fmt.Println(ch)

	go func() {
		fmt.Println("xxx")
		ch <- "goroutine 完成"
	}()
	fmt.Println("main goroutine")

	v := <-ch
	fmt.Println("接收到的chain值为：", v)

	ints := make(chan int, 5)
	fmt.Println(ints)
	ints <- 2
	ints <- 3
	fmt.Println(cap(ints), len(ints))

	close(ch)
	close(ints)

	onlySend := make(chan<- int)
	onlyReceive := make(<-chan int)
	defer close(onlySend)
	println(onlyReceive)

	select {
	case i1 := <-onlyReceive:
		println(i1)
	case c2 := <-ch:
		println(c2)
	default:
		println("没有可用路线")
	}
}
