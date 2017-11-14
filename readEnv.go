package main

import (
	"os"
	"log"
)

func main(){
	log.Fatal(os.Getenv("JAVA_HOME"))
}
