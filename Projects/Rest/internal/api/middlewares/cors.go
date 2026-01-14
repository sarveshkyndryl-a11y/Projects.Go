package middlewares

import "net/http"
func Cors(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if isOriginAllowed(origin){
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}else{
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next.ServeHTTP(w, r)
	})
}

func isOriginAllowed(origin string) bool {
	allowedOrigins := []string{"http://localhost:3000", "http://example.com"}
	for _, o := range allowedOrigins {
		if o == origin {
			return true
		}
	}
	return false
}