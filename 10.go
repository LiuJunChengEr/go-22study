package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, stop := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗1】")
	}()

	time.Sleep(5 * time.Second)
	stop()

	wg.Wait()
}
func watchDog(ctx context.Context, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case d:=<-ctx.Done():
			fmt.Printf("%T",d)
			fmt.Println("协程停止")
			return
		default:
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}
