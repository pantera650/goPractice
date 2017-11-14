package main

import "fmt"

func main(){
	var a =[...]int{1, 2, 3, 4, 5};
	var s = a[1:4];
	fmt.Println(s);
	fmt.Println(s[:cap(s)]);

	var a1 =[...]int{1, 2, 3, 4, 5};
	var s1 = a1[1:4:4];
	fmt.Println(s1);
	fmt.Println(s1[:cap(s1)]);
}
