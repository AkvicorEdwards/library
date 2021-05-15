package handler

import (
	"encoding/json"
	"fmt"
	"library/db"
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

type LoginResponseInfo struct {
	UUID int64 `json:"uuid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	ProfilePhoto string `json:"profile_photo"`
	LastLogin int64 `json:"last_login"`
	Key string `json:"key"`
}

func MarshalLoginResponseInfo(key string, user *db.User) string {
	return Marshal(LoginResponseInfo{
		UUID:         user.UUID,
		Username:     user.Username,
		Nickname:     user.Nickname,
		ProfilePhoto: user.ProfilePhoto,
		LastLogin:    user.LastLogin,
		Key:          key,
	})
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