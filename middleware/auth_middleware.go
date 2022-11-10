package middleware

import (
	"go-salaries-app/helper"
	"go-salaries-app/model/web"
	"net/http"
)

type AuthMiddleware struct {
	http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "rahasia" == request.Header.Get("x-api-key") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResp := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, &webResp)
	}

}
