package middlewares

import "net/http"

//middleware skeleton
func Securityheaders(next http.Handler)http.Handler{
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Security-Policy", "default-src 'self'")
w.Header().Set("X-Content-Type-Options", "nosniff")
w.Header().Set("X-Frame-Options", "DENY")
w.Header().Set("X-XSS-Protection", "1; mode=block")

next.ServeHTTP(w, r)
})
}

// func securityheaders(next http.Handler)http.Handler{
// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// next.ServeHTTP(w, r)
// })
// }