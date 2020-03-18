package main

import (
	"fmt"
	"library/config"
	"library/handler"
	"library/mysql"
	"library/parameter"
	"library/tpl"
	"net/http"
	"time"
)

func main() {
	tpl.Parse()
	parameter.AddBasicArgs()
	config.AddParseModule()
	config.AddParseServer()
	config.AddParseMySQL()
	config.AddParseCert()
	config.AddParsePath()
	parameter.ParseArgs()
	mysql.SetDEFAULT(config.MySQL)
	handler.ParsePrefix()

	server := http.Server {
		Addr:           config.Server.Addr,
		Handler:        &handler.MyHandler{},
		ReadTimeout:    20 * time.Second,
		MaxHeaderBytes: 8<<20,
	}

	fmt.Println("ListenAndServe: ", config.Server.Addr)
	//if err := server.ListenAndServe(); err != nil {
	//	panic(err)
	//}
	if err := server.ListenAndServeTLS(config.Cert.Cert, config.Cert.Key); err != nil {
		panic(err)
	}
}
