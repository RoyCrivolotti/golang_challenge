package ports

import (
	"context"
	"golangchallenge/internal/core/domain"
)

//go:generate mockgen -source=./course.go -destination=./mock/course_mock.go

type ICourseService interface {
	SortCourses(ctx context.Context, courses domain.UserCourseData) []string
}

type ICourseRepository interface {
}
