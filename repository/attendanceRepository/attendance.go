package attendanceRepository

import (
	"database/sql"
	"sirclo/project-capstone/entities/attendanceEntities"
)

type attendanceRepo struct {
	db *sql.DB
}

func NewMySQLDayRepository(db *sql.DB) attendanceRepoInterface {
	return &attendanceRepo{
		db: db,
	}
}

func (ar *attendanceRepo) CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
	query := `INSERT INTO attendances (id, day_id, user_id, created_at) VALUES (?, ?, ?, now())`

	statement, err := ar.db.Prepare(query)
	if err != nil {
		return att, err
	}

	_, errExec := statement.Exec(att.ID, att.Day.ID, att.Employee.ID, att.Status, att.Notes, att.CreatedAt)
	if errExec != nil {
		return att, errExec
	}
	return att, nil
}
