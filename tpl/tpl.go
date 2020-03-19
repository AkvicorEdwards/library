package tpl

import (
	"fmt"
	"html/template"
)

var err error
var errCnt = 0

var Home *template.Template
var Login *template.Template
var Register *template.Template

var Book *template.Template
var BookAdd *template.Template
var BookIndex *template.Template
var BookComplex *template.Template


func Parse() {
	Home = addFromFile("./tpl/ori/home.tpl")
	Login = addFromFile("./tpl/ori/login.tpl")
	Register = addFromFile("./tpl/ori/register.tpl")

	Book = addFromFile("./tpl/ori/book.tpl")
	BookAdd = addFromFile("./tpl/ori/book-add.tpl")
	BookIndex = addFromFile("./tpl/ori/book-index.tpl")
	BookComplex = addFromFile("./tpl/ori/book-complex.tpl")

	fmt.Printf("Parsing the html template was completed with %d errors\n", errCnt)
}

func add(name, tpl string) (t *template.Template) {
	t, err = template.New(name).Parse(tpl)
	if err != nil {
		errCnt++
		fmt.Println(err)
	}
	return
}

func addFromFile(file string) (t *template.Template) {
	t, err = template.ParseFiles(file)
	if err != nil {
		errCnt++
		fmt.Println(err)
	}
	return
}