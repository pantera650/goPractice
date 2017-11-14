package main

import (
	"unsafe"
	"fmt"
)

func main() {
	d := struct {
		s string
		x   int
		y   int
	}{"abc", 100, 1000}
	p := uintptr(unsafe.Pointer(&d))
	p += unsafe.Offsetof(d.y)
	p2 := unsafe.Pointer(p)
	px := (*int)(p2)
	*px = 200
	fmt.Printf("%#v\n", d)
}
