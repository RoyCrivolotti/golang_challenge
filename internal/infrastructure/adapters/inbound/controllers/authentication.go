package controllers

import (
	"encoding/json"
	"fmt"
	"golangchallenge/internal/core/domain"
	"golangchallenge/internal/core/ports"
	"golangchallenge/internal/utils"
	"golangchallenge/pkg/web"
	"net/http"
	"net/mail"
	"strings"
)

type IAuthenticationController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

type authenticationController struct {
	service ports.IAuthenticationService
}

func NewAuthenticationController(authenticationService ports.IAuthenticationService) IAuthenticationController {
	return &authenticationController{
		service: authenticationService,
	}
}

func (c *authenticationController) SignUp(w http.ResponseWriter, r *http.Request) {
	var authenticationData domain.AuthenticationData
	if err := json.NewDecoder(r.Body).Decode(&authenticationData); err != nil {
		utils.Logger.Error(fmt.Sprintf("Authentication error: %s", err.Error()))
		http.Error(w, "To sign up you must provide an email and a password", 400)
		return
	}

	if _, err := mail.ParseAddress(authenticationData.Email); err != nil {
		utils.Logger.Error(fmt.Sprintf("Authentication error: %s", err.Error()))
		http.Error(w, "Invalid email address", 400)
		return
	}

	if strings.Trim(authenticationData.Password, " ") == "" {
		http.Error(w, "Enter a non empty password", 400)
		return
	}

	if err, token := c.service.SignUp(r.Context(), authenticationData); err != nil {
		utils.Logger.Error(fmt.Sprintf("Authentication error: %s", err.Error()))
		http.Error(w, err.Error(), 500)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	}
}

func (c *authenticationController) SignIn(w http.ResponseWriter, r *http.Request) {
	var authenticationData domain.AuthenticationData
	if err := json.NewDecoder(r.Body).Decode(&authenticationData); err != nil {
		utils.Logger.Error(fmt.Sprintf("Authentication error: %s", err.Error()))
		http.Error(w, "To sign in you must provide an email and a password", 400)
		return
	}

	if _, err := mail.ParseAddress(authenticationData.Email); err != nil {
		utils.Logger.Error(fmt.Sprintf("Authentication error: %s", err.Error()))
		http.Error(w, "Invalid email address", 400)
		return
	}

	if err, token := c.service.SignIn(r.Context(), authenticationData); err != nil {
		utils.Logger.Error(fmt.Sprintf("Failed to create token: %s", err.Error()))
		http.Error(w, "Failed to log in", 500)
		return
	} else {
		if err := web.EncodeJSON(w, token, http.StatusOK); err != nil {
			utils.Logger.Error(fmt.Sprintf("Unexpected error generating response object: %s", err.Error()))
			http.Error(w, "Unexpected server error", 500)
			return
		}
	}
}
