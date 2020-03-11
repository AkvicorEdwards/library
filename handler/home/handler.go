package home

import (
	"fmt"
	"html/template"
	"net/http"
	"library/config"
)

func Page(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles(config.Data.Path.Theme + "home.tpl")

	phrase := map[string]interface{}{
		"lang": "zh-cn",
		"title": config.Data.Server.Title,
	}

	if err := t.Execute(w, phrase); err != nil {
		fmt.Println(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
}

