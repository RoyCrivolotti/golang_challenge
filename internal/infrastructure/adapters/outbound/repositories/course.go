package repositories

import (
	"golangchallenge/internal/core/ports"

	"github.com/jinzhu/gorm"
)

type courseRepository struct {
	mysqlClient *gorm.DB
}

func NewCourseRepository(mysqlClient *gorm.DB) ports.ICourseRepository {
	return &courseRepository{
		mysqlClient: mysqlClient,
	}
}
