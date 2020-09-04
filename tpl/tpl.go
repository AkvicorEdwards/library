package tpl

import (
	"html/template"
	"log"
)

var err error
var errCnt = 0

var Index *template.Template
var Login *template.Template
var Register *template.Template
var BookAdd *template.Template
var BookIndex *template.Template
var BookComplex *template.Template
var Book *template.Template

func Parse() {
	err = nil
	errCnt = 0

	Index = addFromFile("./tpl/index.html")
	Login = addFromFile("./tpl/login.html")
	Register = addFromFile("./tpl/register.html")
	BookAdd = addFromFile("./tpl/book-add.html")
	BookIndex = addFromFile("./tpl/book-index.html")
	BookComplex = addFromFile("./tpl/book-complex.html")
	Book = addFromFile("./tpl/book.html")

	log.Printf("Parsing the html template was completed with %d errors\n", errCnt)
}

func add(name, tpl string) (t *template.Template) {
	t, err = template.New(name).Parse(tpl)
	if err != nil {
		errCnt++
		log.Println(err)
	}
	return
}

func addFromFile(file string) (t *template.Template) {
	t, err = template.ParseFiles(file)
	if err != nil {
		errCnt++
		log.Println(err)
	}
	return
}