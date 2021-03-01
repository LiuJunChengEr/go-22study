package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum = 0
	//mutex sync.Mutex
	mutex sync.RWMutex
)

func main() {
	for i := 0; i < 100; i++ {
		go add(10)
	}
	//time.Sleep(2 * time.Second)
	//fmt.Println("和为：", sum)

	for i := 0; i < 10; i++ {
		go fmt.Println("和为：", readSum1())
	}
	time.Sleep(2 * time.Second)
	fmt.Println("和为：", readSum())

	run()

	doOnce()

	race()
}

func add(i int) {
	mutex.Lock()
	sum += i
	mutex.Unlock()
}

func add1(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

func readSum() int {
	mutex.Lock()
	defer mutex.Unlock()
	b := sum
	return b
}

func readSum1() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

func run() {
	var wg sync.WaitGroup
	wg.Add(110)

	for i := 0; i < 100; i++ {
		go func() {
			// 计数器-1
			defer wg.Done()
			add(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("和为:", readSum1())
		}()
	}

	wg.Wait()
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)

	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			println(num, "号已经就位")
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			println(num, "号开始跑......")
		}(i)
	}

	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		println("裁判已经就位，准备发令枪")
		println("比赛开始，大家准备跑")
		cond.Broadcast()
	}()

	wg.Wait()
}
