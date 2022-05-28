package domain

type UserCourseData struct {
	ID      string   `json:"userId" gorm:"primary_key"`
	Courses []Course `json:"courses"`
}

type Course struct {
	Name       string `json:"desiredCourse"`
	Dependency string `json:"requiredCourse"`
}
