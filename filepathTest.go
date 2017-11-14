package main

import (
	"log"
	"path/filepath"
	"os"
	"github.com/astaxie/beego/utils"
)

func main(){
	p,e:=os.Getwd()
	if e!=nil{
		log.Print(e)
	}

	log.Print(p)
	path := os.Args[0]
	log.Print(path)

	log.Print(utils.FileExists(filepath.Join(path, "filepathtest.go")))
	log.Print(utils.FileExists(filepath.Join(path, "filepathTest.go")))

	var err error
	path, err = filepath.Abs(filepath.Dir(path))
	if(err != nil){
		log.Print(err)
	}
	log.Println(path)

}
