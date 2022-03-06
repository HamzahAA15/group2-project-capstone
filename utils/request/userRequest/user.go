package userRequest

type CreateUserInput struct {
	ID       string `json:"id" form:"id"`
	OfficeID string `json:"office_id" form:"office_id"`
	Nik      string `json:"nik" form:"nik"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Role     string `json:"role" form:"role"`
}

type UpdateUserInput struct {
	Nik      string `json:"nik" form:"nik"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
}

type LoginUserInput struct {
	Identity string `json:"identity" form:"identity"`
	Password string `json:"password" form:"password"`
}
