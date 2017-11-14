package main

import (
	"html/template"
	"log"
	"os"
)

func main(){
	t := template.New("我的网页")
	t, err := t.Parse(`
	<html>
		<head>
			<title>{{.title}}</title>
		</head>
		<body>
			{{.content}}
		</body>
	</html>
	`)

	if err != nil{
		log.Fatal("网页分析错误")
	}

	if err:=t.Execute(os.Stdout, map[string]interface{}{"title":"我的标题", "content":"我的内容"}); err!=nil{
		log.Fatal("输出错误")
	}
}
