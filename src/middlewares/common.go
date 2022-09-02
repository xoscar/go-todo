package middlewares

import (
	"net/http"
)

var CommonMiddleware Middleware = commonMiddleware{}

type commonMiddleware struct{}

func (c commonMiddleware) GetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
