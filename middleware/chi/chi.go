package chi

import (
	"bufio"
	"context"
	"errors"
	"net"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/slok/go-http-metrics/middleware"
)

// Handler returns a chi measuring middleware.
func Handler(handlerID string, m middleware.Middleware, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		wi := &responseWriterInterceptor{
			statusCode:     http.StatusOK,
			ResponseWriter: w,
		}
		reporter := &reporter{
			w: wi,
			r: req,
		}

		m.Measure(handlerID, reporter, func() {
			h.ServeHTTP(wi, req)
		})
	}
}

type reporter struct {
	w *responseWriterInterceptor
	r *http.Request
}

func (s *reporter) Method() string { return s.r.Method }

func (s *reporter) Context() context.Context { return s.r.Context() }

func (s *reporter) URLPath() string {
	path := s.r.URL.Path

	if ctx := chi.RouteContext(s.r.Context()); ctx != nil {
		return strings.TrimRight(ctx.RoutePattern(), "/")
	}

	return path
}

func (s *reporter) StatusCode() int { return s.w.statusCode }

func (s *reporter) BytesWritten() int64 { return int64(s.w.bytesWritten) }

// responseWriterInterceptor is a simple wrapper to intercept set data on a
// ResponseWriter.
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterInterceptor) Write(p []byte) (int, error) {
	w.bytesWritten += len(p)
	return w.ResponseWriter.Write(p)
}

func (w *responseWriterInterceptor) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("type assertion failed http.ResponseWriter not a http.Hijacker")
	}
	return h.Hijack()
}

func (w *responseWriterInterceptor) Flush() {
	f, ok := w.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}

	f.Flush()
}
