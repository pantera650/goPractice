package main

func main(){
	c := make(chan string, 5)

	c <- "hello world"
}
