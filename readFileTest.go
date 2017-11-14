package main

import (
	"io/ioutil"
	"os"
	"log"
	"path/filepath"
	"beego_tmp/utils"
	"bytes"
	"bufio"
//	"runtime"
	"reflect"
	"io"
	"strings"
)

func main(){
	curPath,err := os.Getwd()

	if err!=nil{
		log.Fatal(err)
	}

	//curFilePath := filepath.Dir(curPath)

	var p string
	var err2 error
	if p,err2 = filepath.Abs(filepath.Join(curPath, "chan.go"));err2!=nil{
		log.Fatal(err2)
	}

	if utils.FileExists(p){

		b, err3 := ioutil.ReadFile(p)

		if err3!=nil{
			log.Fatal(err3)
		}
		/*
		文件头两个字节是255 254，为Unicode编码；
        文件头三个字节  254 255 0，为UTF-16BE编码；
        文件头三个字节  239 187 191，为UTF-8编码；
		*/
		buf := bufio.NewReader(bytes.NewBuffer(b))

		head, err := buf.Peek(3)
		log.Println(reflect.TypeOf(head))
		if err == nil && head[0] == 239 && head[1] == 187 && head[2] == 191 {
			for i := 1; i <= 3; i++ {
				buf.ReadByte()
			}
		}

		for{
			line, _, err := buf.ReadLine()
			if err == io.EOF {
				break
			}

			//It might be a good idea to throw a error on all unknonw errors?
			if _, ok := err.(*os.PathError); ok {
				log.Fatal("文件路径错误")
			}

			if bytes.Equal(line, []byte{}) {
				continue
			}

			var bComment []byte
			//log.Fatal(bComment == nil) 	//true
			switch {
				case bytes.HasPrefix(line, []byte{'#'}):
					bComment = []byte{'#'}
				case bytes.HasPrefix(line, []byte{';'}):
					bComment = []byte{';'}
			}

			var comment bytes.Buffer
			if nil != bComment{
				line = bytes.TrimLeft(line, string(bComment))
				// Need append to a new line if multi-line comments.
				if comment.Len() > 0 {
					comment.WriteByte('\n')
				}
				comment.Write(line)
				continue

			}
			//
			//if bytes.HasPrefix(line, sectionStart) && bytes.HasSuffix(line, sectionEnd) {
			//	section = strings.ToLower(string(line[1 : len(line)-1])) // section name case insensitive
			//	if comment.Len() > 0 {
			//		cfg.sectionComment[section] = comment.String()
			//		comment.Reset()
			//	}
			//	if _, ok := cfg.data[section]; !ok {
			//		cfg.data[section] = make(map[string]string)
			//	}
			//	continue
			//}

			log.Println(line)
			//log.Println(strings.ToLower(string(line[1 : len(line)-1])))
			log.Println(strings.ToLower(string(line)))
		}
	}else{
		log.Println("文件" + p + "不存在")
	}

	keyValue := bytes.SplitN([]byte("weisheme这是个实物  = zheshige"), []byte{'='}, 2)
	log.Println(string(bytes.TrimSpace(keyValue[0]))+ "|")

	log.Println(strings.HasPrefix("include(this is test)", "include"))
	log.Fatal(strings.Fields("  foo bar  baz   "))
}
