package userResponse

import (
	"sirclo/project-capstone/entities/userEntities"
)

type AuthFormatter struct {
	Token string `json:"token"`
}

func FormatAuth(token string) AuthFormatter {
	formatter := AuthFormatter{
		Token: token,
	}

	return formatter
}

type UserFormatter struct {
	ID       string `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

func FormatUser(user userEntities.User) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Avatar:   user.Avatar,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Phone:    user.Phone,
	}

	return formatter
}
