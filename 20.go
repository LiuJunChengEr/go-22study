package main

import (
	"fmt"
	"os"
)

func init(){
	fmt.Println("init in main.go")
}


func main() {
	fmt.Println("先导入fmt包，才能使用")
	println(os.Getuid())
}