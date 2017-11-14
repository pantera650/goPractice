package main

import "log"

func main(){
	log.Print(f1())
	log.Print(f2())
	log.Print("s1")
	defer func(){
		if err:=recover();err!=nil{
			log.Print(err)
		}
	}()
	log.Print("s2")

	f()
	log.Print("s3")
}

func f(){
	log.Print("f1")
	panic(55)
	log.Print("f2")
}

func f1() (result int) {
	defer func() {

		result++

	}()

	return 0
}

func f2() (result int) {
	return 0

	defer func() {

		result++

	}()
	return 0
}
