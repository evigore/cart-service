package middleware

import (
	"bytes"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	Body       *bytes.Buffer
	StatusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		Body:           &bytes.Buffer{},
		StatusCode:     200,
	}
}

func (w *responseWriter) Write(data []byte) (int, error) {
	return w.Body.Write(data)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
}

func (w *responseWriter) Close() {
	res := w.Body.Bytes()
	w.ResponseWriter.WriteHeader(w.StatusCode)
	w.ResponseWriter.Write(res)
}
