package controllers

import (
	"encoding/json"
	"fmt"
	"golangchallenge/internal/core/domain"
	"golangchallenge/internal/core/ports"
	"golangchallenge/internal/utils"
	"golangchallenge/pkg/web"
	"net/http"
)

type ICourseController interface {
	SortCourses(w http.ResponseWriter, r *http.Request)
}

type courseController struct {
	service ports.ICourseService
}

func NewCourseController(courseService ports.ICourseService) ICourseController {
	return &courseController{
		service: courseService,
	}
}

func (c *courseController) SortCourses(w http.ResponseWriter, r *http.Request) {
	var userCourseData domain.UserCourseData
	if err := json.NewDecoder(r.Body).Decode(&userCourseData); err != nil {
		utils.Logger.Error(fmt.Sprintf("Invalid JSON of courses: %s", err.Error()))
		http.Error(w, "Please enter a valid JSON of courses", 400)
		return
	}

	token := c.service.SortCourses(r.Context(), userCourseData)
	web.EncodeJSON(w, token, http.StatusOK)
}
