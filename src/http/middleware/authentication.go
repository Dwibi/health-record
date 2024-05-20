package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dwibi/health-record/src/helpers"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the request headers
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parse the JWT token
		token, err := jwt.ParseWithClaims(tokenString, &helpers.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// TODO: create errorResponse for this
		if err != nil {
			helpers.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Verify the token and extract claims
		claims, ok := token.Claims.(*helpers.Claims)

		if !ok || !token.Valid {
			helpers.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Attach user information to the request context
		ctx := context.WithValue(r.Context(), helpers.UserContextKey, claims.UserId)
		handlerFunc(w, r.WithContext(ctx))
	}
}
