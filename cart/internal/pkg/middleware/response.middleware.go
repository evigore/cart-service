package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		responseWriter := newResponseWriter(w)
		defer responseWriter.Close()

		next.ServeHTTP(responseWriter, req)

		w.Header().Set("content-type", "applications/json")
		if responseWriter.StatusCode < 400 { // not wrapping non-error responses
			return
		}

		errorMessage := responseWriter.Body.Bytes()
		res := errorResponse{Error: string(errorMessage)}
		bytes, err := json.Marshal(res)
		if err != nil {
			log.Println("failed to return an error:", err)
		}

		responseWriter.Body.Reset()
		responseWriter.Write(bytes)
	})
}
