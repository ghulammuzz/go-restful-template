package cors

import (
	"net/http"
)

type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

func CORS(config CORSConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", getAllowedOrigin(r, config.AllowedOrigins))
			w.Header().Set("Access-Control-Allow-Methods", joinMethods(config.AllowedMethods))
			w.Header().Set("Access-Control-Allow-Headers", joinHeaders(config.AllowedHeaders))

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getAllowedOrigin(r *http.Request, allowedOrigins []string) string {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return ""
	}
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return origin
		}
	}
	return ""
}

func joinMethods(methods []string) string {
	if len(methods) == 0 {
		return ""
	}
	return methods[0]
}

func joinHeaders(headers []string) string {
	if len(headers) == 0 {
		return ""
	}
	return headers[0]
}
