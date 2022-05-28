package unit_test

import (
	"context"
	"golangchallenge/internal/core/domain"
	mock_ports "golangchallenge/internal/core/ports/mock"
	"golangchallenge/internal/infrastructure/adapters/inbound/controllers"
	"golangchallenge/internal/infrastructure/configuration"
	mock_authentication "golangchallenge/internal/infrastructure/configuration/authentication/mock"
	"golangchallenge/internal/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	firebase "firebase.google.com/go"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthenticationService := mock_ports.NewMockIAuthenticationService(ctrl)
	mockAuthenticationService.EXPECT().SignUp(gomock.Any(), domain.AuthenticationData{
		Email:    "test@test.com",
		Password: "123123",
	}).Return(nil, "token").AnyTimes()

	authenticationController := controllers.NewAuthenticationController(mockAuthenticationService)
	utils.InitTest(t)

	router := chi.NewRouter()

	authenticationMiddlewareMock := mock_authentication.NewMockIAuthenticationMiddleware(ctrl)
	authenticationMiddlewareMock.EXPECT().Authenticate(gomock.Any()).AnyTimes()

	firebaseApp, _ := firebase.NewApp(context.Background(), nil)
	firebaseAuth, _ := firebaseApp.Auth(context.Background())
	srv := configuration.NewServer(router)

	db, _, _ := sqlmock.New()
	defer db.Close()
	gormDB, _ := gorm.Open("mysql")

	srv.Initialize(firebaseAuth, authenticationMiddlewareMock, gormDB)

	router.Post("/user/signup", authenticationController.SignUp)

	tests := []struct {
		description string
		input       string
		want        string
	}{
		{
			"TestCase: Happy path",
			`{"email":"test@test.com","password":"123123"}`,
			"token",
		},
		{
			"TestCase: Invalid email",
			`{"email":"asdasd","password":"123123"}`,
			"Invalid email address\n",
		},
		{
			"TestCase: Invalid email",
			`{"email":"test@test.com","password":"  "}`,
			"Enter a non empty password\n",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
			t.Parallel()
			rr := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/user/signup", strings.NewReader(tt.input))
			request.Header.Set("Content-Type", "application/json")
			assert.NoError(t, err)
			router.ServeHTTP(rr, request)

			println("rr.Body.Bytes()")
			println(rr.Body.String())
			assert.Equal(t, tt.want, rr.Body.String())
		})
	}
}
