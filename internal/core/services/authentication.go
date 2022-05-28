package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golangchallenge/internal/core/domain"
	"golangchallenge/internal/core/ports"
	"golangchallenge/internal/errorx"
	"net/http"
	"os"
	"strings"

	"firebase.google.com/go/auth"
)

type authenticationService struct {
	firebaseAuth *auth.Client
}

func NewAuthenticationService(firebaseAuth *auth.Client) ports.IAuthenticationService {
	return &authenticationService{
		firebaseAuth: firebaseAuth,
	}
}

func (s *authenticationService) SignUp(ctx context.Context, authenticationData domain.AuthenticationData) (errorx.Error, string) {
	params := (&auth.UserToCreate{}).
		Email(authenticationData.Email).
		EmailVerified(false).
		UID(authenticationData.Email).
		Password(authenticationData.Password).
		Disabled(false)

	_, err := s.firebaseAuth.CreateUser(ctx, params)

	if err != nil {
		return errorx.NewInternalServerError(
				"Failed to register new user",
				err.Error(),
				"AuthenticationService-SignUp"),
			""
	}

	return s.createToken(ctx, authenticationData)
}

func (s *authenticationService) SignIn(ctx context.Context, authenticationData domain.AuthenticationData) (errorx.Error, string) {
	_, err := s.firebaseAuth.GetUserByEmail(ctx, authenticationData.Email)

	if err != nil {
		return errorx.NewInternalServerError(
				"Failed to sign in, make sure you entered the correct email",
				err.Error(),
				"AuthenticationService-SignIn"),
			""
	}

	return s.createToken(ctx, authenticationData)
}

func (s *authenticationService) createToken(ctx context.Context, authenticationData domain.AuthenticationData) (errorx.Error, string) {
	signInUrl := fmt.Sprintf("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=%s", os.Getenv("FIREBASE_API_KEY"))
	reqBody := map[string]interface{}{
		"email":             authenticationData.Email,
		"password":          authenticationData.Password,
		"returnSecureToken": true,
	}

	buffer := new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(reqBody); err != nil {
		return errorx.NewInternalServerError(
				"Unexpected server error",
				err.Error(),
				"AuthenticationService-SignUp"),
			""
	}

	resp, err := http.Post(signInUrl, "application/json", buffer)
	defer resp.Body.Close()

	if err != nil {
		return errorx.NewInternalServerError(
				"Unexpected error when trying to sign up",
				err.Error(),
				"AuthenticationService-SignUp"),
			""
	}

	decoder := json.NewDecoder(resp.Body)
	firebaseAuthResponse := struct {
		Kind         string `json:"kind"`
		LocalId      string `json:"localId"`
		Email        string `json:"email"`
		DisplayName  string `json:"displayName"`
		IdToken      string `json:"idToken"`
		Registered   bool   `json:"registered"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
	}{}

	if err := decoder.Decode(&firebaseAuthResponse); err != nil || strings.Trim(firebaseAuthResponse.IdToken, " ") == "" {
		return errorx.NewInternalServerError(
				"Unexpected error when trying to create bearer token",
				err.Error(),
				"AuthenticationService-createToken"),
			""
	}

	return nil, firebaseAuthResponse.IdToken
}
