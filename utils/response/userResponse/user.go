package userResponse

import (
	"sirclo/project-capstone/entities/userEntities"
)

type AuthFormatter struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

func FormatAuth(token, role string) AuthFormatter {
	formatter := AuthFormatter{
		Token: token,
		Role:  role,
	}

	return formatter
}

type UserFormatter struct {
	ID         string `json:"id"`
	OfficeName string `json:"office_name"`
	Avatar     string `json:"avatar"`
	Nik        string `json:"nik"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
}

func FormatUser(user userEntities.User) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		OfficeName: user.Office.Name,
		Avatar:     user.Avatar,
		Nik:        user.Nik,
		Email:      user.Email,
		Name:       user.Name,
		Phone:      user.Phone,
		Role:       user.Role,
	}

	return formatter
}
