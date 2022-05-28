package handlers

import (
	"golangchallenge/internal/core/ports"
	"golangchallenge/internal/core/services"
	"golangchallenge/internal/infrastructure/adapters/outbound/repositories"

	"firebase.google.com/go/auth"
	"github.com/jinzhu/gorm"
)

type ServiceHandler struct {
	CourseService         ports.ICourseService
	AuthenticationService ports.IAuthenticationService
}

func NewServiceHandler(firebaseAuth *auth.Client, db *gorm.DB) *ServiceHandler {
	authenticationService := services.NewAuthenticationService(firebaseAuth)

	courseRepository := repositories.NewCourseRepository(db)
	courseService := services.NewCourseService(courseRepository)

	return &ServiceHandler{
		CourseService:         courseService,
		AuthenticationService: authenticationService,
	}
}
