package sessions

import "github.com/gorilla/sessions"

// Хранилище для данных сессии
var (
	cookiePassword = "password"
	CookieStore = sessions.NewCookieStore([]byte(cookiePassword))
)