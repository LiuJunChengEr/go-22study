package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c1"}
	println(array[0])
	println(array[2])

	array1 := [...]string{"1", "2", "3"}
	println(&array1)

	array2 := [5]string{1: "b", 3: "d"}
	println(&array2)

	for i, v := range array {
		println(i, v)
	}

	for _, v := range array {
		println(v)
	}

	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice := array3[2:5]
	println(slice)

	slice1 := array3[2:5]
	slice1[1] = "f"
	fmt.Println(array3)

	// 创建了一个元素类型为string的切片，长度为4
	slice2 := make([]string, 4)
	// 新创建的切片[]string容量为8
	slice3 := make([]string, 4, 8)
	fmt.Println(slice2, slice3)

	slice4 := []string{"a", "b", "c", "d", "e"}
	println(len(slice4), cap(slice4))

	slice5 := append(slice4, "f")
	fmt.Println(slice5)
	slice6 := append(slice4, "f", "g")
	fmt.Println(slice6)
	slice7 := append(slice4, slice1...)
	fmt.Println(slice7)

	nameAgeMap := make(map[string]int)
	nameAgeMap["年龄"] = 18
	fmt.Println(nameAgeMap)

	nameAgeMap1 := map[string]int{"年龄": 18}
	fmt.Println(nameAgeMap1)

	println(nameAgeMap1["年龄"])

	age, ok := nameAgeMap1["年龄"]
	if ok {
		println(age)
	}

	delete(nameAgeMap1,"年龄")

	println(len(nameAgeMap))

	s:="abcde"
	bs:=[]byte(s)
	fmt.Println(bs)

	println(s[0],s[2])
}
