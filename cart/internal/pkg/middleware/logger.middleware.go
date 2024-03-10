package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		responseWriter := newResponseWriter(w)
		defer responseWriter.Close()

		next.ServeHTTP(responseWriter, req)

		if responseWriter.StatusCode >= 500 {
			res := responseWriter.Body.Bytes()
			log.Printf("%s %s %s - %s", req.Method, req.RequestURI, time.Since(start), res)
		} else {
			log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
		}
	})
}
