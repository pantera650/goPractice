package main

import (
	"fmt"
	"log"
)

type A struct{
	B
}
type B struct{}
func (b *B) say (s string){
	fmt.Println(s)
}

type test struct{
	a int
}

func (t *test2)my(s string){
	fmt.Print(s)
}

type test2 struct{
	data map[int]string
}

func main() {
	log.Fatal(outputType("hello"))
	a := &A{B{}}
	a.say("hello world")
	log.Fatal("end")
	panic(make(map[string]map[string]string))
	v := test{a:100}
	fmt.Println(v)
	modifySliceData(&v)
	fmt.Println(v)

	s := "unmodify"
	s = "modify";
	fmt.Print(s)
	modify(s)
	fmt.Print(s)
}

type myInterface interface{
	my (s string)
}

func getObj() myInterface{
	return &test2{
		map[int]string{100:"hell"},
	}
}

func modify(s string){
	s="modify"
}

func modifySliceData(t *test) {
	t.a = 1000
}

func outputType(i interface{}) string{
	var r string = ""
	switch i.(type) {      //多选语句switch
		case string:
			r="string"
		//是字符时做的事情
		case int:
			r="int"
		//是整数时做的事情
	}
	return r
}