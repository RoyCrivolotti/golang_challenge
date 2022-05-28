package domain

type AuthenticationData struct {
	Email    string `json:"email" gorm:"primary_key"`
	Password string `json:"password"`
}
