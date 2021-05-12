package main

import (
	"fmt"
	"github.com/AkvicorEdwards/arg"
	"library/def"
	"library/handler"
	"library/repair"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	EnableOption()
	repair.EnableShutDownListener()
	handler.ParsePrefix()

	addr := fmt.Sprintf("%s:%d", def.ADDR, def.PORT)
	server := http.Server{
		Addr:              addr,
		Handler:           &handler.MyHandler{},
		ReadTimeout:       20 * time.Minute,
	}
	log.Println(addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func EnableOption() {
	arg.RootCommand.Size = 0
	arg.RootCommand.Executor = func([]string) error { return nil }
	option("init", 0, func(str []string) error {
		repair.InitDatabase()
		os.Exit(0)
		return nil
	})
	option("-db", 1, func(str []string) error {
		def.DBFilename = str[1]
		return nil
	})
	option("-port", 1, func(str []string) error {
		var err error = nil
		def.PORT, err = strconv.Atoi(str[1])
		if err != nil || def.PORT <= 0 || def.PORT > 65535 {
			log.Println("PORT ERROR")
			os.Exit(-1)
		}
		return nil
	})
	option("-sd", 1, func(str []string) error {
		def.SessionDomain = str[1]
		return nil
	})
	option("-sp", 1, func(str []string) error {
		def.SessionPath = str[1]
		return nil
	})
	option("-sn", 1, func(str []string) error {
		def.SessionName = str[1]
		return nil
	})

	arg.EnableOptionCombination()

	wrap(arg.Parse())
}

func option(opt string, size int, f arg.FuncExecutor) {
	wrap(arg.AddOption([]string{opt}, 0, size, 0, "",
		"", "", "", f, nil))
}

func options(opt1, opt2 string, size int, f arg.FuncExecutor) {
	wrap(arg.AddOption([]string{opt1, opt2}, 0, size, 0, "",
		"", "", "", f, nil))
}

func command(cmd string, size int, f arg.FuncExecutor) {
	wrap(arg.AddCommand([]string{cmd}, 0, size, "",
		"", "", "", f, nil))
}

func wrap(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
}