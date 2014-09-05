package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Temp struct {
	Name string
}

func mkComponent(item Temp) {

	file := fmt.Sprintf("%s/%s.html", item.Name, item.Name)
	os.Mkdir(item.Name, 0644)
	f, _ := os.Create(file)
	defer f.Close()

	t := template.Must(template.ParseFiles("dummy/c-dummy.html"))
	err := t.Execute(f, item)
	if err != nil {
		fmt.Println("err")
	}
}
func mkDemo(item Temp) {

	f, _ := os.Create(item.Name + "Demo.html")
	defer f.Close()

	t := template.Must(template.ParseFiles("dummy/d-dummy.html"))
	err := t.Execute(f, item)
	if err != nil {
		fmt.Println("err")
	}
}

func main() {
	cname := flag.String("c", "nil", "component name")
	flag.Parse()
	if *cname == "nil" {
		fmt.Println("This is Polymer Component Maker")
		fmt.Println("You wnat to make component")
		fmt.Println("You should typing \nlike this \"go run polyPack.go -c=test-component\"")
		fmt.Println("\n\n\nCopyright (c) 2014 YuJM")
		fmt.println("zsxzsx1010@gmail.com")
	} else {
		if strings.Contains(*cname, "-") {
			tt := Temp{*cname}
			mkComponent(tt)
			fmt.Printf("Make <%s> component\n", *cname)
			mkDemo(tt)
			fmt.Printf("Make %sDemo\n", *cname)
			println("Successful!")
		} else {
			fmt.Println("Component name is contain '-'")
			fmt.Println("Exam :'test-compo'")
		}
	}
}
