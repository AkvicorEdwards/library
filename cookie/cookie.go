package cookie

import (
	"library/db"
	"library/def"
	"net/http"
	"time"
)

// 设置客户端Cookie
// 有效期1年
func SetUserInfo(w http.ResponseWriter, user *db.User) {
	// Cookie
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)

	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   user.Username,
		Path:    def.SessionPath,
		Domain:  def.SessionDomain,
		Expires: expiration,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "nickname",
		Value:   user.Nickname,
		Path:    def.SessionPath,
		Domain:  def.SessionDomain,
		Expires: expiration,
	})
}
