package attendanceService

import (
	"fmt"
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/repository/attendanceRepository"
	"sirclo/project-capstone/repository/userRepository"
	"sirclo/project-capstone/utils/request/attendanceRequest"

	"github.com/google/uuid"
)

type attendanceService struct {
	attRepo  attendanceRepository.AttendanceRepoInterface
	userRepo userRepository.UserRepoInterface
}

func NewAttendanceService(attRepo attendanceRepository.AttendanceRepoInterface, userRepo userRepository.UserRepoInterface) AttServiceInterface {
	return &attendanceService{
		attRepo:  attRepo,
		userRepo: userRepo,
	}
}

func (as *attendanceService) GetAttendances(employee, date, status, office, order string) ([]attendanceEntities.Attendance, error) {
	attendances, err := as.attRepo.GetAttendances(employee, date, status, office, order)
	if err != nil {
		return attendances, err
	}
	return attendances, nil
}

func (as *attendanceService) GetAttendancesRangeDate(employee, dateStart, dateEnd, status, office, order string) ([]attendanceEntities.Attendance, error) {
	attendances, err := as.attRepo.GetAttendancesRangeDate(employee, dateStart, dateEnd, status, office, order)
	if err != nil {
		return attendances, err
	}
	return attendances, nil
}

func (as *attendanceService) GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error) {
	attendances, err := as.attRepo.GetAttendancesCurrentUser(userId, status, order)
	if err != nil {
		return attendances, err
	}
	return attendances, nil
}

func (as *attendanceService) CreateAttendance(loginId string, input attendanceRequest.CreateAttRequest) (attendanceEntities.Attendance, error) {
	var attendance attendanceEntities.Attendance
	attendance.ID = uuid.New().String()
	attendance.Day.ID = input.Day
	attendance.Employee.ID = loginId

	createAttendance, err := as.attRepo.CreateAttendance(attendance)
	if err != nil {
		return attendance, err
	}
	fmt.Println("err svc: ", err)
	return createAttendance, nil
}

func (as *attendanceService) UpdateAttendance(loginId string, input attendanceRequest.UpdateAttRequest) (attendanceEntities.Attendance, error) {
	var attendance attendanceEntities.Attendance

	attendance.ID = input.ID
	attendance.Status = input.Status
	attendance.Notes = input.Notes
	attendance.Admin.ID = loginId

	updateAttendance, err := as.attRepo.UpdateAttendance(attendance)
	if err != nil {
		return attendance, err
	}
	return updateAttendance, nil
}

func (as *attendanceService) CheckUserRole(loginId string) string {
	currentUser, _ := as.userRepo.GetUser(loginId)
	if currentUser.Role != "admin" {
		return currentUser.Role
	}
	return currentUser.Role
}
