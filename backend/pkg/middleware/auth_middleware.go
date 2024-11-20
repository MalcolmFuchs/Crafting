package middleware

import (
	"net/http"
	"strings"

	"github.com/MalcolmFuchs/Crafting/internal/auth"
)

func AuthenticationMiddleware(next http.Handler, secretKey []byte) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: Kein Authorization-Header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitAfterN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Unauthorized: Ungültiges Authorization-Format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		_, err := auth.VerifyAccessToken(tokenString, secretKey)
		if err != nil {
			http.Error(w, "Unauthorized: Ungültiges Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
