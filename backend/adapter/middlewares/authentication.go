package middlewares

import (
	"backend/adapter/modules/authentication"
	"log"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, "Start authenticating...")

		authentication.Authen()

		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, "Done authenticate.")
	})

}
