package csrf

import (
	"net/http"
)

// CSRFMiddleware is a middleware function that handles CSRF protection
func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate CSRF token and set it as a cookie
		token, err := generateCSRFToken()
		if err != nil {
			http.Error(w, "Error generating CSRF token", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    token,
			HttpOnly: true,
		})

		// Check for valid CSRF token
		if r.Method == http.MethodPost {
			var token string
			if r.Header.Get("X-CSRF-Token") != "" {
				token = r.Header.Get("X-CSRF-Token")
			} else {
				token = r.FormValue("csrf_token")
			}
			if token != cookie.Value {
				http.Error(w, "Invalid CSRF token", http.StatusBadRequest)
				return
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
