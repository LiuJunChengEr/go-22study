package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

func main() {
	i := 3
	//iv := reflect.ValueOf(i)
	iv := reflect.ValueOf(&i)
	it := reflect.TypeOf(i)
	fmt.Println(iv, it)

	//i1 := iv.Interface().(int)
	//println(i1)

	iv.Elem().SetInt(4)
	fmt.Println(i)

	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
}

func main1() {
	p := person{"xxx", 18}
	ppv := reflect.ValueOf(&p)
	ppv.Elem().Field(0).SetString("yyy")
	fmt.Println(p)
}

type person struct {
	Name string `json:"name""`
	Age  int `json:"age"`
}

func main2() {
	p := person{"xxx", 20}
	ppv := reflect.ValueOf(&p)
	fmt.Println(ppv.Kind())

	pv := reflect.ValueOf(p)
	fmt.Println(pv.Kind())
}

func main3() {
	p := person{"xxx", 18}
	pt := reflect.TypeOf(p)

	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段：", pt.Field(i).Name)
	}

	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法：", pt.Method(i).Name)
	}
}

func (p person) String() string {
	return fmt.Sprintf("name is %s,age is %d", p.Name, p.Age)
}

func main4() {
	p := person{"xxx", 18}
	pt := reflect.TypeOf(p)
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer：", pt.Implements(stringerType))
	fmt.Println("是否实现了fmt.io.Writer：", pt.Implements(writerType))
}

func main5() {
	p := person{"xxx", 18}
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}

	respJSON := "{\"Name\":\"xxx\",\"Age\":18}"
	err = json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p, err)
}

func main6(){
	p := person{"xxx", 18}
	pt := reflect.TypeOf(p)

	for i:=0;i<pt.NumField();i++{
		sf := pt.Field(i)
		fmt.Printf("字段：%s中，json tag为：%s\n",sf.Name,sf.Tag.Get("json"))
	}
}

func main7(){
	p := person{"xxx", 18}
	pv := reflect.ValueOf(p)
	mPrint := pv.MethodByName("Print")
	args :=[]reflect.Value{reflect.ValueOf("开始")}
	mPrint.Call(args)
}

func (p person) Print(prefix string){
	fmt.Printf("%s:Name is %s,Age is %d\n",prefix,p.Name,p.Age)
}