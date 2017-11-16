package main

import (
	"log"
	"runtime"
	"time"
)

func main(){
	c := make(chan int, 0)
	maxProcs := runtime.NumCPU()   //获取cpu个数
	runtime.GOMAXPROCS(maxProcs)  //限制同时运行的goroutines数量
	log.Print("hello")
	go func(s string){
		log.Println(s)
		c <- 100
	}("world")

	time.Sleep(10*time.Second)
	log.Println(<-c)

}
