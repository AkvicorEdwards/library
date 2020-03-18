package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

// Note: Don't store your key in your source code. Pass it via an
// environmental variable, or flag (or both), and don't accidentally commit it
// alongside your code. Ensure your key is sufficiently random - i.e. use Go's
// crypto/rand or securecookie.GenerateRandomKey(32) and persist the result.
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func Get(r *http.Request, name string) (*sessions.Session, error) {
	return store.Get(r, name)
}

func New(r *http.Request, name string) (*sessions.Session, error)  {
	return store.New(r, name)
}

func Save(w http.ResponseWriter, r *http.Request, session *sessions.Session) error {
	return store.Save(r, w, session)
}

func MaxAge(age int) {
	store.MaxAge(age)
}


//func use(w http.ResponseWriter, r *http.Request) {
//	// Get a session. We're ignoring the error resulted from decoding an
//	// existing session: Get() always returns a session, even if empty.
//	session, _ := store.Get(r, "session-name")
//	// Set some session values.
//	session.Values["foo"] = "bar"
//	session.Values[42] = 43
//	// Save it before we write to the response/return from the handler.
//	err := session.Save(r, w)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
