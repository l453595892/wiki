package main

import "net/http"

// NewCustomResponseWriter new responsewriter
func NewCustomResponseWriter(w http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{w: w}
}

// CustomResponseWriter custom response writer
type CustomResponseWriter struct {
	w    http.ResponseWriter
	done bool
}

// Complete if set , can not write into body
func (w *CustomResponseWriter) Complete() {
	w.done = true
}

func (w *CustomResponseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *CustomResponseWriter) Write(bts []byte) (int, error) {
	if w.done {
		return 0, nil
	}
	return w.w.Write(bts)
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.w.WriteHeader(statusCode)
}
