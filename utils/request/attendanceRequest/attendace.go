package attendanceRequest

type CreateAttRequest struct {
	ID       string `json:"id" form:"id"`
	Day      string `json:"day" form:"day"`
	Employee string `json:"employee" form:"employee"`
}

type UpdateAttRequest struct {
	Status string `json:"status" form:"status"`
	Notes  string `json:"notes" form:"notes"`
}
