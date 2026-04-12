package middleware

import "net/http"

// Auth is a stub middleware that will validate JWT Bearer tokens once
// authorization is implemented. Currently it passes all requests through.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: parse Authorization header, validate JWT, inject userID into context
		next.ServeHTTP(w, r)
	})
}
