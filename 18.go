package main

func main() {

}

//F(0)=0
//F(1)=1
//F(n)=F(n - 1)+F(n - 2)
func Fibonacci(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
