package internalhttp

import (
	"fmt"
	"net/http"
	"time"
)

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		s.app.Logger.Info(fmt.Sprintf("%s [%s] %s %s %d %s",
			r.RemoteAddr,
			time.Now(),
			r.Method,
			r.RequestURI,
			http.StatusOK,
			r.UserAgent(),
		))
	})
}
