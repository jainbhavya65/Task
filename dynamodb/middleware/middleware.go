package middleware

import (
	"encoding/json"
	"context"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	
)

var jwtKey = []byte("my_secret_key")

func AlreadyLoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "{\"error\": \"Unauthorized\"}", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		tknStr := c.Value

		token, err := jwt.Parse(tknStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "{\"error\": \"Unauthorized\"}", http.StatusUnauthorized)
				return
			}
			http.Error(w, "{\"error\": \"Unauthorized\"}", http.StatusUnauthorized)
			return	
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "props", claims)
			// Access context values in handlers like this
			// props, _ := r.Context().Value("props").(jwt.MapClaims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "{\"error\": \"Unauthorized\"}", http.StatusUnauthorized)
			return
		}
	})
}

func CheckBodyJson(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var t interface{}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err == nil {
			ctx := context.WithValue(r.Context(), "req", t)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "{\"error\": \"Invalid JSON\"}", http.StatusBadRequest)
		}
	})
}