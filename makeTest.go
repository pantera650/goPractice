package main

import "log"

func main(){
	a := []int{1,2,3,4}
	log.Print(a)

	b := make([]int, 4)
	log.Print(b)

	c := map[string]int{"tom":100, "jerry":20}
	log.Println(c)

	d := make(map[string]int)
	//d = c
	log.Print(d)
}
