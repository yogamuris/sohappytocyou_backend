package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"github.com/yogamuris/sohappytocyou/helper"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("sohappytocyou_token")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := token.Value

		claims := &web.Claims{}

		var jwtKey = []byte(helper.GetEnv(".env", "JWT_KEY"))

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
