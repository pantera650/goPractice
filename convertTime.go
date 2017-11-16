package main

import (
	"log"
	"time"
	"fmt"
	"strconv"
)

func main(){
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	log.Println(time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18"))
	t := time.Now()
	fmt.Println(t)

	fmt.Println(t.UTC().Format(time.UnixDate))

	fmt.Println(t.Unix())

	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	fmt.Println(timestamp)
	timestamp = timestamp[:10]
	fmt.Println(timestamp)
}
