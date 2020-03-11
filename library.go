package main

import (
	"fmt"
	"net/http"
	"time"

	"library/config"
	"library/handler"
	"library/models/utilities/database"
)

func main() {
	config.ParseYaml()
	handler.ParsePrefix()
	database.Init()

	server := http.Server {
		Addr:              config.Data.Server.Addr,
		Handler:           &handler.MyHandler{},
		ReadTimeout:       20 * time.Second,
		MaxHeaderBytes:	   8<<20,
	}

	fmt.Println("ListenAndServe: ", config.Data.Server.Addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}


