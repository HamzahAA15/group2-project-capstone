package attendanceRequest

type CreateAttRequest struct {
	ID       string `json:"id" form:"id"`
	Day      string `json:"day" form:"day"`
	Employee string `json:"employee" form:"employee"`
}
