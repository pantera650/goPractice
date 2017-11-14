package main

import "log"

type myTest struct{
	age int
	name string
}

func main(){
	firstman := myTest{20, "张三"}
	f(&firstman)
	f1(firstman)
	log.Println(firstman)
	f2(&firstman)
	log.Println(firstman)

	second := new(myTest)
	second.age = 35
	second.name = "李四"
	f(second)
	f1(*second)
	log.Println(second)
	f2(second)
	log.Println(second)

	thirdman := &myTest{20, "张三"}
	f(thirdman)
	f1(*thirdman)
	log.Println(thirdman)
	f2(thirdman)
	log.Println(thirdman)
}

func f(mt *myTest){
	log.Println(mt)
}

func f1(mt myTest){
	mt.age = 10000
}

func f2(mt *myTest){
	mt.age=2000
}
