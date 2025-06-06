package middlewares

import "net/http"

const origin = "http://localhost:5173"

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		// w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		// w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method=="OPTIONS"{
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w,r)
	})
} 