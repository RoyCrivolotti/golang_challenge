package unit_test

import (
	"context"
	"encoding/json"
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

func TestSortCourses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCourseService := mock_ports.NewMockICourseService(ctrl)
	coursesData := domain.UserCourseData{ID: "30ecc27b-9df7-4dd3", Courses: []domain.Course{{Name: "PortfolioConstruction", Dependency: "PortfolioTheories"}}}
	mockCourseService.EXPECT().SortCourses(gomock.Any(), coursesData).Return([]string{"Finance", "Investment"}).AnyTimes()

	courseController := controllers.NewCourseController(mockCourseService)
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

	router.Post("/courses/sort", courseController.SortCourses)

	tests := []struct {
		description string
		input       string
		want        *[]string
	}{
		{
			"TestCase: Happy path",
			`{"userId":"30ecc27b-9df7-4dd3","courses":[{"desiredCourse":"PortfolioConstruction","requiredCourse":"PortfolioTheories"}]}`,
			&[]string{"Finance", "Investment"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
			t.Parallel()
			rr := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/courses/sort", strings.NewReader(tt.input))
			request.Header.Set("Content-Type", "application/json")
			assert.NoError(t, err)
			router.ServeHTTP(rr, request)

			body := new([]string)
			json.Unmarshal(rr.Body.Bytes(), body)
			assert.Equal(t, tt.want, body)
		})
	}
}
