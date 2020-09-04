package cookie

import (
	"fmt"
	"library/dam"
	"net/http"
	"time"
)

// 设置客户端Cookie
// 有效期1年
// 包含 [uuid], [username], [nickname]
//     [plant], [rmb_in], [rmb_out], [coin]
func SetUserInfo(w http.ResponseWriter, user *dam.User) {
	// Cookie
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)

	http.SetCookie(w, &http.Cookie{
		Name:    "uuid",
		Value:   fmt.Sprint(user.Uuid),
		Expires: expiration,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   user.Username,
		Expires: expiration,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "nickname",
		Value:   user.Nickname,
		Expires: expiration,
	})

}
