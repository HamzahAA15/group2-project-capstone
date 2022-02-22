package attendanceRequest

type CreateAttRequest struct {
	Day string `json:"day" form:"day"`
}

type UpdateAttRequest struct {
	ID     string `json:"id" form:"id"`
	Status string `json:"status" form:"status"`
	Notes  string `json:"notes" form:"notes"`
	Admin  string `json:"admin" form:"admin"`
}
