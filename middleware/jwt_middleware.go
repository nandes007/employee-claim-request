package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/web"
	"net/http"
	"os"
	"strings"
)

func JwtAuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenSigning := os.Getenv("JWT_SIGNING")
		bearerToken := r.Header.Get("Authorization")
		tokenString := strings.Replace(bearerToken, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(tokenSigning), nil
		})

		if err != nil || !token.Valid {
			apiResponse := web.ApiResponse{
				Code:   401,
				Status: "Failed",
				Data:   "Unauthorized",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			helper.WriteToResponseBody(w, apiResponse)
			return
		}

		next(w, r, ps)
	}
}
