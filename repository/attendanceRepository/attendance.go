package attendanceRepository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sirclo/project-capstone/entities/attendanceEntities"
)

type attendanceRepo struct {
	db *sql.DB
}

func NewMySQLDayRepository(db *sql.DB) AttendanceRepoInterface {
	return &attendanceRepo{
		db: db,
	}
}

func (ar *attendanceRepo) GetAttendances(employee, date, status, office, order string) ([]attendanceEntities.Attendance, error) {
	var attendances []attendanceEntities.Attendance
	convEmployee := "%" + employee + "%"
	convTime := "%" + date + "%"
	convStatus := "%" + status + "%"
	convOffice := "%" + office + "%"

	var query string
	query = `
		SELECT
			attendances.id, day.date AS date, office.id AS office_id, office.name, user.avatar, user.email, user.nik, user.name as employee, attendances.status, (COALESCE(NULLIF(attendances.notes,''), '-')) AS notes, (COALESCE(NULLIF(admin.name,''), '-')) AS admin 
		FROM 
			attendances
		LEFT JOIN
			days AS day ON day.id = attendances.day_id
		LEFT JOIN
			offices AS office ON office.id = day.office_id
		LEFT JOIN
			users AS user ON user.id = attendances.user_id
		LEFT JOIN
			users AS admin ON admin.id = attendances.admin_id
		WHERE
			user.name LIKE ? AND day.date LIKE ? AND attendances.status LIKE ? AND office.name LIKE ?
		ORDER BY attendances.created_at %s`
	result, err := ar.db.Query(fmt.Sprintf(query, order), convEmployee, convTime, convStatus, convOffice)
	if err != nil {
		return attendances, err
	}

	for result.Next() {
		var attendance attendanceEntities.Attendance

		errScan := result.Scan(&attendance.ID, &attendance.Day.Date, &attendance.OfficeId, &attendance.Office, &attendance.Employee.Avatar, &attendance.Employee.Email, &attendance.Employee.Nik, &attendance.Employee.Name, &attendance.Status, &attendance.Notes, &attendance.Admin.Name)

		if errScan != nil {
			return attendances, errScan
		}
		attendances = append(attendances, attendance)

	}
	return attendances, nil
}

func (ar *attendanceRepo) GetAttendancesRangeDate(employee, dateStart, dateEnd, status, office, order string) ([]attendanceEntities.Attendance, error) {
	var attendances []attendanceEntities.Attendance
	convEmployee := "%" + employee + "%"
	convStatus := "%" + status + "%"
	convOffice := "%" + office + "%"

	var query string
	query = `
		SELECT
			attendances.id, day.date AS date, office.id, office.name, user.avatar, user.email, user.nik, user.name as employee, attendances.status, (COALESCE(NULLIF(attendances.notes,''), '-')) AS notes, (COALESCE(NULLIF(admin.name,''), '-')) AS admin 
		FROM 
			attendances
		LEFT JOIN
			days AS day ON day.id = attendances.day_id
		LEFT JOIN
			offices AS office ON office.id = day.office_id
		LEFT JOIN
			users AS user ON user.id = attendances.user_id
		LEFT JOIN
			users AS admin ON admin.id = attendances.admin_id
		WHERE
			user.name LIKE ? AND (day.date BETWEEN ? AND ?) AND attendances.status LIKE ? AND office.name LIKE ?
		ORDER BY attendances.created_at %s`
	result, err := ar.db.Query(fmt.Sprintf(query, order), convEmployee, dateStart, dateEnd, convStatus, convOffice)
	if err != nil {
		return attendances, err
	}

	for result.Next() {
		var attendance attendanceEntities.Attendance

		errScan := result.Scan(&attendance.ID, &attendance.Day.Date, &attendance.OfficeId, &attendance.Office, &attendance.Employee.Avatar, &attendance.Employee.Email, &attendance.Employee.Nik, &attendance.Employee.Name, &attendance.Status, &attendance.Notes, &attendance.Admin.Name)

		if errScan != nil {
			return attendances, errScan
		}
		attendances = append(attendances, attendance)

	}
	return attendances, nil
}

func (ar *attendanceRepo) GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error) {
	var attendances []attendanceEntities.Attendance
	convStatus := "%" + status + "%"
	var query string
	query = `
		SELECT
			attendances.id, day.date AS date, office.name, user.avatar, user.email, user.nik, user.name as employee, attendances.status, (COALESCE(NULLIF(attendances.notes,''), '-')) AS notes, (COALESCE(NULLIF(admin.name,''), '-')) AS admin 
		FROM 
			attendances
		LEFT JOIN
			days AS day ON day.id = attendances.day_id
		LEFT JOIN
			offices AS office ON office.id = day.office_id
		LEFT JOIN
			users AS user ON user.id = attendances.user_id
		LEFT JOIN
			users AS admin ON admin.id = attendances.admin_id
		WHERE
			user.id = ? AND attendances.status LIKE ?
		ORDER BY attendances.created_at %s`
	result, err := ar.db.Query(fmt.Sprintf(query, order), userId, convStatus)
	if err != nil {
		return attendances, err
	}

	for result.Next() {
		var attendance attendanceEntities.Attendance

		errScan := result.Scan(&attendance.ID, &attendance.Day.Date, &attendance.Office, &attendance.Employee.Avatar, &attendance.Employee.Email, &attendance.Employee.Nik, &attendance.Employee.Name, &attendance.Status, &attendance.Notes, &attendance.Admin.Name)

		if errScan != nil {
			return attendances, errScan
		}
		attendances = append(attendances, attendance)

	}
	return attendances, nil
}

func (ar *attendanceRepo) CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
	var checkDate string
	var checkDouble int
	errNoRows := ar.db.QueryRow(`SELECT days.date FROM days WHERE days.id = ? AND days.date > now()`, att.Day.ID).Scan(&checkDate)
	switch {
	case errNoRows == sql.ErrNoRows:
		return att, errors.New("day has passed")
	case errNoRows != nil:
		log.Fatal(errNoRows)
	}

	errDouble := ar.db.QueryRow(`SELECT count(id) as id FROM attendances WHERE day_id = ? AND user_id = ?`, att.Day.ID, att.Employee.ID).Scan(&checkDouble)
	if errDouble != nil {
		log.Fatal(errDouble)
	}
	if checkDouble != 0 {
		return att, errors.New("you already sign for current date")
	}

	query := `INSERT INTO attendances (id, day_id, user_id, created_at) VALUES (?, ?, ?, now())`

	statement, err := ar.db.Prepare(query)
	if err != nil {
		return att, err
	}

	_, errExec := statement.Exec(att.ID, att.Day.ID, att.Employee.ID)
	if errExec != nil {
		return att, errExec
	}
	return att, nil
}

func (ar *attendanceRepo) UpdateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
	query := `UPDATE attendances SET status = ?, notes = ?, admin_id = ?, updated_at = now() WHERE id = ?`

	statement, err := ar.db.Prepare(query)
	if err != nil {
		return att, err
	}

	_, errExec := statement.Exec(att.Status, att.Notes, att.Admin.ID, att.ID)
	if errExec != nil {
		return att, errExec
	}
	return att, nil
}
