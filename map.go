package main

import "fmt"

func main(){
	m := map[string]int{"tom":10, "jack":20, "dick":15}

	fmt.Println(m)

	v,ok := m["tom"]

	fmt.Println(v, ok)

	v,ok = m["w"]

	fmt.Println(v, ok)
}
