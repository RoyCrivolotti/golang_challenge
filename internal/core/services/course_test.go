package services_test

import (
	"context"
	"golangchallenge/internal/core/domain"
	mock_ports "golangchallenge/internal/core/ports/mock"
	"golangchallenge/internal/core/services"
	"golangchallenge/internal/utils"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSortCourses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	utils.InitTest(t)

	tests := []struct {
		description string
		input       domain.UserCourseData
		want        []string
	}{
		{
			"TestCase: Happy path",
			domain.UserCourseData{
				ID: "30ecc27b-9df7-4dd3-b52f-d001e79bd035",
				Courses: []domain.Course{
					{Name: "PortfolioConstruction", Dependency: "PortfolioTheories"},
					{Name: "InvestmentManagement", Dependency: "Investment"},
					{Name: "Investment", Dependency: "Finance"},
					{Name: "PortfolioTheories", Dependency: "Investment"},
					{Name: "InvestmentStyle", Dependency: "InvestmentManagement"},
				},
			},
			[]string{"Finance", "Investment", "InvestmentManagement", "PortfolioTheories", "InvestmentStyle", "PortfolioConstruction"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
			t.Parallel()
			mockCourseRepository := mock_ports.NewMockICourseRepository(ctrl)
			courseService := services.NewCourseService(mockCourseRepository)
			res := courseService.SortCourses(context.Background(), tt.input)
			assert.Equal(t, tt.want, res)
		})
	}
}
