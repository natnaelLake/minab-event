package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

// Logger is a middleware function that logs HTTP requests
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Create a response writer that captures the status code
		lrw := &loggingResponseWriter{w, http.StatusOK}

		// Call the next handler in the chain
		next(lrw, r)

		// Log the request details
		logEntry := fmt.Sprintf(
			"%s %d %s %s %s\n",
			r.RemoteAddr,
			lrw.statusCode,
			r.Method,
			r.RequestURI,
			time.Since(startTime),
		)

		fmt.Print(logEntry)
	}
}

// loggingResponseWriter wraps the http.ResponseWriter to capture the status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code for logging
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
