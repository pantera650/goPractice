package main

import (
	"log"
	"os"
	"text/template"
	htemplate "html/template"
)

type Note struct {
	Title       string
	Description string
}

const templ = `Note - Title:{{.Title}},Description:{{.Description}}`

func main() {
	log.Print("this", " is", " test")
	note := Note{
		"我是标题",
		"我是描述",
		}

	//create a template with name
	t := template.New("note")

	//解析内容到模板
	t, err := t.Parse(templ)
	if err != nil {
		log.Fatal("Parse:", err)
		return
	}

	h := htemplate.New("网页输出测试")

	h,err = h.Parse(`
	<html>
		<head>
			<title>{{.title}}</title>
		</head>
		<body>
			{{.content}}
		</body>
	</html>
	`);

	if err != nil {
		log.Fatal("Parse:", err)
		return
	}



	//将数据用到模板中
	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Execute:", err)
		return
	}

}
