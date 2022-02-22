package attendanceService

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/repository/attendanceRepository"
	"sirclo/project-capstone/utils/request/attendaceRequest"
)

type attendanceService struct {
	attRepo attendanceRepository.AttendanceRepoInterface
}

func NewAttendanceService(attRepo attendanceRepository.AttendanceRepoInterface) AttServiceInterface {
	return &attendanceService{
		attRepo: attRepo,
	}
}

func (as *attendanceService) CreateAttendace(input attendaceRequest.CreateAttRequest) (attendanceEntities.Attendance, error) {
	var attendance attendanceEntities.Attendance
	attendance.ID = input.ID
	attendance.Day.ID = input.Day
	attendance.Employee.ID = input.Employee

	createAttendance, err := as.attRepo.CreateAttendance(attendance)
	if err != nil {
		return attendance, err
	}
	return createAttendance, nil
}
