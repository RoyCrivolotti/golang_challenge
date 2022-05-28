package handlers

import (
	"golangchallenge/internal/infrastructure/adapters/inbound/controllers"
)

type ControllerHandler struct {
	CourseController         controllers.ICourseController
	AuthenticationController controllers.IAuthenticationController
}

func NewControllerHandler(serviceHandler *ServiceHandler) *ControllerHandler {
	return &ControllerHandler{
		CourseController:         controllers.NewCourseController(serviceHandler.CourseService),
		AuthenticationController: controllers.NewAuthenticationController(serviceHandler.AuthenticationService),
	}
}
