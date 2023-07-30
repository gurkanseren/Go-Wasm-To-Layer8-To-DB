package middleware

import (
	"log"
	"net/http"
	"time"
)

// statusCapturingResponseWriter is a wrapper around an http.ResponseWriter that captures the status code written to it.
type statusCapturingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code written to the response writer.
func (w *statusCapturingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// LogRequest is a middleware that logs incoming requests and their responses.
func LogRequest(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Use the custom response writer to capture the status code
		crw := &statusCapturingResponseWriter{w, http.StatusOK}
		// Call the next handler in the chain with the custom response writer
		next.ServeHTTP(crw, r)
		// Log information about the request
		clientIP := r.RemoteAddr
		method := r.Method
		uri := r.RequestURI
		statusCode := crw.statusCode
		processingTime := time.Since(start).Milliseconds()

		log.Printf("%s - %s %s - %d - %dms", clientIP, method, uri, statusCode, processingTime)
	}
}
