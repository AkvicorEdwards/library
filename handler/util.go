package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func TimeUnixFormat(t int64) string {
	if t == 0 {
		return ""
	}
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Int64(i string) int64 {
	res, _ := strconv.ParseInt(i, 10, 64)
	return res
}

func Int(i string) int {
	res, _ := strconv.Atoi(i)
	return res
}

func Fprint(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprint(w, a...)
}

func Fprintf(w http.ResponseWriter, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, format, a...)
}

func Fprintln(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprintln(w, a...)
}

const StatusOK = 0
const StatusER = 1

const ResponseOK = `{"status":"0"}`
const ResponseER = `{"status":"1"}`

type ResponseStruct struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

func Marshal(data interface{}) string {
	res := ResponseStruct{
		Status: StatusOK,
		Data:   data,
	}
	str, err := json.Marshal(res)
	if err != nil {
		return fmt.Sprintf(`{"status":"%d"}`, StatusER)
	}
	return string(str)
}


func TplRedirect(url string) string {
	tpl := `<!DOCTYPE html>
<script type="text/javascript">
function url_confirm() {
`
	if url == "history" {
		tpl += "window.history.back();"
	} else if len(url) != 0 {
		tpl += fmt.Sprintf(`window.location.href="%s";`, url)
	}
	tpl += `
}
url_confirm()
</script>`
	return tpl
}