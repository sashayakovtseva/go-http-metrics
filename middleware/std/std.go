// Package std is a helper package to get a standard `http.Handler` compatible middleware.
package std

import (
	"bufio"
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/slok/go-http-metrics/middleware"
)

// Handler returns an measuring standard http.Handler.
func Handler(handlerID string, m middleware.Middleware, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wi := &ResponseWriterInterceptor{
			statusCode:     http.StatusOK,
			ResponseWriter: w,
		}
		reporter := &Reporter{
			w: wi,
			r: r,
		}

		m.Measure(handlerID, reporter, func() {
			h.ServeHTTP(wi, r)
		})
	})
}

// HandlerProvider is a helper method that returns a handler provider. This kind of
// provider is a defacto standard in some frameworks (e.g: Gorilla, Chi...).
func HandlerProvider(handlerID string, m middleware.Middleware) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return Handler(handlerID, m, next)
	}
}

type Reporter struct {
	w *ResponseWriterInterceptor
	r *http.Request
}

func (s *Reporter) Method() string { return s.r.Method }

func (s *Reporter) Context() context.Context { return s.r.Context() }

func (s *Reporter) URLPath() string { return s.r.URL.Path }

func (s *Reporter) StatusCode() int { return s.w.statusCode }

func (s *Reporter) BytesWritten() int64 { return int64(s.w.bytesWritten) }

// ResponseWriterInterceptor is a simple wrapper to intercept set data on a
// ResponseWriter.
type ResponseWriterInterceptor struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
}

func (w *ResponseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *ResponseWriterInterceptor) Write(p []byte) (int, error) {
	w.bytesWritten += len(p)
	return w.ResponseWriter.Write(p)
}

func (w *ResponseWriterInterceptor) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("type assertion failed http.ResponseWriter not a http.Hijacker")
	}
	return h.Hijack()
}

func (w *ResponseWriterInterceptor) Flush() {
	f, ok := w.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}

	f.Flush()
}

// Check interface implementations.
var (
	_ http.ResponseWriter = &ResponseWriterInterceptor{}
	_ http.Hijacker       = &ResponseWriterInterceptor{}
	_ http.Flusher        = &ResponseWriterInterceptor{}
)
