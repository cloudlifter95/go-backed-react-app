package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

// ResponseWriter Wrapper to Capture Response Body
type responseRecorder struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)                  // Capture response body
	return r.ResponseWriter.Write(b) // Write to actual response
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode // Capture status code
	r.ResponseWriter.WriteHeader(statusCode)
}

// LoggingMiddleware logs requests and responses
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Read and log request body
		var requestBody bytes.Buffer
		if r.Body != nil {
			bodyBytes, _ := io.ReadAll(r.Body)
			requestBody.Write(bodyBytes)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset body for next handler
		}
		log.Printf("Request: %s %s | Body: %s", r.Method, r.URL.Path, requestBody.String())

		// Capture Response
		recorder := &responseRecorder{
			ResponseWriter: w,
			body:           new(bytes.Buffer),
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(recorder, r)

		// Log Response
		duration := time.Since(start)
		log.Printf("Response: %d | Duration: %v | Body: %s\n", recorder.statusCode, duration, recorder.body.String())
	})
}
