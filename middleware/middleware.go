package middleware

import (
	"../sessions"
	"net/http"
)

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var (
			session, _ = sessions.CookieStore.Get(request, "session")
			_, ok = session.Values["user_id"]
		)
		if !ok {
			http.Redirect(writer, request, "/login", 302)
			return
		}
		handler.ServeHTTP(writer, request)
	}
}