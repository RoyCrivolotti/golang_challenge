package authentication

import (
	"context"
	"fmt"
	"golangchallenge/internal/utils"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
)

type IAuthenticationMiddleware interface {
	Authenticate(next http.Handler) http.Handler
}

type authenticationMiddleware struct {
	firebaseAuth *auth.Client
}

//go:generate mockgen -source=./middleware.go -destination=./mock/middleware_mock.go

func NewAuthenticationMiddleware(firebaseAuth *auth.Client) IAuthenticationMiddleware {
	return &authenticationMiddleware{
		firebaseAuth: firebaseAuth,
	}
}

type ctxKey int

const userContextKey ctxKey = iota

func (m *authenticationMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authorizationHeader := r.Header.Get("Authorization")
		var bearerToken string

		authHeaderParts := strings.Split(authorizationHeader, " ")
		if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], "bearer") {
			http.Error(w, "Bearer token cannot be empty", 400)
			return
		} else {
			bearerToken = authHeaderParts[1]
		}

		token, err := m.firebaseAuth.VerifyIDToken(ctx, bearerToken)

		if err != nil {
			utils.Logger.Debug(fmt.Sprintf("Failed to verify JWT: Authentication header '%s' - %s", authorizationHeader, err.Error()))
			http.Error(w, "Failed to authenticate user", 401)
			return
		}

		ctx = context.WithValue(ctx, userContextKey, map[string]string{"UUID": token.UID})

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
