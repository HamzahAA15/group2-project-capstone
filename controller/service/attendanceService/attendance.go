package attendanceService

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/repository/attendanceRepository"
	"sirclo/project-capstone/repository/userRepository"
	"sirclo/project-capstone/utils/request/attendanceRequest"
	"time"

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

func (as *attendanceService) GetAttendancesById(attID string) (string, string, error) {
	userId, employeeName, err := as.attRepo.GetAttendancesById(attID)
	return userId, employeeName, err
}

func (as *attendanceService) GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order string) ([]attendanceEntities.Attendance, error) {

	t := time.Now()
	if dateStart == "" {
		dateStart = time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location()).String()
	}
	if dateEnd == "" {
		dateEnd = time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location()).String()
	}

	attendances, err := as.attRepo.GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order)
	return attendances, err
}

func (as *attendanceService) GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error) {
	attendances, err := as.attRepo.GetAttendancesCurrentUser(userId, status, order)
	return attendances, err
}

func (as *attendanceService) IsCheckins() ([]attendanceEntities.Attendance, error) {
	attendances, err := as.attRepo.IsCheckins()
	return attendances, err
}

func (as *attendanceService) CreateAttendance(loginId string, input attendanceRequest.CreateAttRequest) (attendanceEntities.Attendance, error) {
	var attendance attendanceEntities.Attendance
	attendance.ID = uuid.New().String()
	attendance.Day.ID = input.Day
	attendance.Employee.ID = loginId

	createAttendance, err := as.attRepo.CreateAttendance(attendance)
	return createAttendance, err
}

func (as *attendanceService) UpdateAttendance(loginId string, input attendanceRequest.UpdateAttRequest) (attendanceEntities.Attendance, error) {
	var attendance attendanceEntities.Attendance

	attendance.ID = input.ID
	attendance.Status = input.Status
	attendance.Notes = input.Notes
	attendance.Admin.ID = loginId

	updateAttendance, err := as.attRepo.UpdateAttendance(attendance)
	return updateAttendance, err
}
