package ports

import (
	"context"
	"golangchallenge/internal/core/domain"
	"golangchallenge/internal/errorx"
)

//go:generate mockgen -source=./authentication.go -destination=./mock/authentication_mock.go

type IAuthenticationService interface {
	SignUp(ctx context.Context, authenticationData domain.AuthenticationData) (err errorx.Error, token string)
	SignIn(ctx context.Context, authenticationData domain.AuthenticationData) (err errorx.Error, token string)
}
