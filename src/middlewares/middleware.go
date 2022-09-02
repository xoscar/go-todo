package middlewares

import (
	"net/http"
)

type Middleware interface {
	GetMiddleware(next http.Handler) http.Handler
}
