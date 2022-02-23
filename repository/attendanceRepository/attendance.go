package attendanceRepository

import (
	"database/sql"
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

func (ar *attendanceRepo) GetAttendances() ([]attendanceEntities.Attendance, error) {
	var attendances []attendanceEntities.Attendance

	result, err := ar.db.Query(`SELECT id, day_id, user_id, status, notes, admin_id FROM attendances`)
	if err != nil {
		return attendances, err
	}

	for result.Next() {
		var attendance attendanceEntities.Attendance

		errScan := result.Scan(&attendance.ID, &attendance.Day.ID, &attendance.Employee.ID, &attendance.Status, &attendance.Notes, &attendance.Admin.ID)

		if errScan != nil {
			return attendances, errScan
		}
		attendances = append(attendances, attendance)
	}
	return attendances, nil
}

func (ar *attendanceRepo) CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
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
