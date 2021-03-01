package main

func main() {
	var s string
	println(s)

	var sp *string
	println(sp)
	sp = new(string)
	*sp = "xxx"
	println(*sp)

	m := make(map[string]int, 10)
	println(m)
}
