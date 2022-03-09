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
func (ar *attendanceRepo) GetAttendancesById(attID string) (string, string, error) {
	var userId string
	var name string
	query := `SELECT attendances.user_id, users.name FROM attendances LEFT JOIN users ON users.id = attendances.user_id WHERE attendances.id = ?`

	err := ar.db.QueryRow(query, attID).Scan(&userId, &name)
	fmt.Println(userId, name)

	return userId, name, err
}

func (ar *attendanceRepo) GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order string) ([]attendanceEntities.Attendance, error) {
	var attendances []attendanceEntities.Attendance
	convEmployee := "%" + employeeEmail + "%"
	convStatus := "%" + status + "%"
	convOfficeId := "%" + officeId + "%"
	query := `
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
			user.email LIKE ? AND (day.date BETWEEN ? AND ?) AND attendances.status LIKE ? AND office.id LIKE ?
		ORDER BY attendances.created_at %s`
	result, err := ar.db.Query(fmt.Sprintf(query, order), convEmployee, dateStart, dateEnd, convStatus, convOfficeId)
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
	query := `
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

func (ar *attendanceRepo) IsCheckins() ([]attendanceEntities.Attendance, error) {
	var attendances []attendanceEntities.Attendance

	result, err := ar.db.Query(`
	SELECT 
		attendances.id, users.email, users.name, users.avatar, offices.name, "Checkin" AS is_checkins
		FROM 
			attendances
		LEFT JOIN
			days ON days.id = attendances.day_id
		LEFT JOIN
			offices ON offices.id = days.office_id
		LEFT JOIN
			users ON users.id = attendances.user_id
		WHERE
			attendances.id IN (
				SELECT checkins.attendance_id FROM checkins
			) AND days.date = DATE(NOW()) AND attendances.status = "approved"
		GROUP BY 
			attendances.id
	UNION
	SELECT 
		attendances.id, users.email, users.name, users.avatar, offices.name, "Not Checkin" AS is_checkins
		FROM 
			attendances
		LEFT JOIN
			days ON days.id = attendances.day_id
		LEFT JOIN
			offices ON offices.id = days.office_id
		LEFT JOIN
			users ON users.id = attendances.user_id
		WHERE
			attendances.id NOT IN (
				SELECT checkins.attendance_id FROM checkins
			) AND days.date = DATE(NOW()) AND attendances.status = "approved"
		GROUP BY 
			attendances.id`)
	if err != nil {
		return attendances, err
	}

	for result.Next() {
		var attendance attendanceEntities.Attendance

		errScan := result.Scan(
			&attendance.ID,
			&attendance.Employee.Email, &attendance.Employee.Name, &attendance.Employee.Avatar,
			&attendance.Day.OfficeId.Name, &attendance.StatusCheckin,
		)

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

	errDouble := ar.db.QueryRow(`
	SELECT 
		count(attendances.id) as id 
	FROM 
		attendances
	LEFT JOIN 
		days ON 
			days.id = attendances.day_id
	WHERE 
		days.date = (SELECT 
			days.date FROM days WHERE days.id = ?) AND user_id = ?`, att.Day.ID, att.Employee.ID).Scan(&checkDouble)
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
