package middlewares

import "net/http"

// ApplyMiddleware applies a list of middlewares to an HTTP handler function
func ApplyMiddleware(h http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
