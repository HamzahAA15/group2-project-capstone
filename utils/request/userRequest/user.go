package userRequest

type CreateUserInput struct {
	ID       string `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateUserInput struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginUserInput struct {
	Identity string `json:"identity" form:"identity"`
	Password string `json:"password" form:"password"`
}
