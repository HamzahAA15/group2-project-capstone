package dayRequest

type DayUpdateRequest struct {
	ID    string `json:"id" form:"id"`
	Quota int    `json:"quota" form:"quota"`
}
