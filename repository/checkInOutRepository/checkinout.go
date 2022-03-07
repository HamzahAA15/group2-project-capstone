package checkInOutRepository

import (
	"database/sql"
	"errors"
	"log"
	"sirclo/project-capstone/entities/checkinEntities"
)

type checkInOutRepo struct {
	db *sql.DB
}

func NewMySQLCheckInOutRepository(db *sql.DB) CheckInOutRepoInterface {
	return &checkInOutRepo{
		db: db,
	}
}

func (cr *checkInOutRepo) Gets() ([]checkinEntities.Checkin, error) {
	var checkinsouts []checkinEntities.Checkin

	result, err := cr.db.Query(`
	SELECT
		checkins.id, checkins.attendance_id, checkins.temprature, checkins.is_checkins, checkins.is_checkouts, checkins.created_at, checkins.updated_at,
		users.name, users.avatar, users.email,
		offices.name
	FROM
		checkins
	JOIN 
		attendances ON attendances.id = checkins.attendance_id
	JOIN 
		days ON days.id = attendances.day_id
	JOIN 
		users ON users.id = attendances.user_id
	JOIN 
		offices ON offices.id = days.office_id
	ORDER BY
		checkins.created_at DESC`)
	if err != nil {
		return checkinsouts, err
	}

	for result.Next() {
		var checkinout checkinEntities.Checkin

		errScan := result.Scan(
			&checkinout.ID, &checkinout.Attendance.ID, &checkinout.Temprature, &checkinout.IsCheckIns, &checkinout.IsCheckOuts, &checkinout.CreatedAt, &checkinout.UpdatedAt,
			&checkinout.Attendance.Employee.Name, &checkinout.Attendance.Employee.Avatar, &checkinout.Attendance.Employee.Email,
			&checkinout.Attendance.Day.OfficeId.Name,
		)

		if errScan != nil {
			return checkinsouts, errScan
		}

		checkinsouts = append(checkinsouts, checkinout)
	}

	return checkinsouts, nil
}

func (cr *checkInOutRepo) GetByUser(userID string) ([]checkinEntities.Checkin, error) {
	var checkinsouts []checkinEntities.Checkin

	result, err := cr.db.Query(`
	SELECT
		checkins.id, checkins.attendance_id, checkins.temprature, checkins.is_checkins, checkins.is_checkouts, checkins.created_at, checkins.updated_at,
		users.name, users.avatar, users.email,
		offices.name
	FROM
		checkins
	JOIN 
		attendances ON attendances.id = checkins.attendance_id
	JOIN 
		days ON days.id = attendances.day_id
	JOIN 
		users ON users.id = attendances.user_id
	JOIN 
		offices ON offices.id = days.office_id
	WHERE
		attendances.user_id = ?
	ORDER BY
		checkins.created_at DESC`, userID)
	if err != nil {
		return checkinsouts, err
	}

	for result.Next() {
		var checkinout checkinEntities.Checkin

		errScan := result.Scan(
			&checkinout.ID, &checkinout.Attendance.ID, &checkinout.Temprature, &checkinout.IsCheckIns, &checkinout.IsCheckOuts, &checkinout.CreatedAt, &checkinout.UpdatedAt,
			&checkinout.Attendance.Employee.Name, &checkinout.Attendance.Employee.Avatar, &checkinout.Attendance.Employee.Email,
			&checkinout.Attendance.Day.OfficeId.Name,
		)

		if errScan != nil {
			return checkinsouts, errScan
		}

		checkinsouts = append(checkinsouts, checkinout)
	}

	return checkinsouts, nil
}

func (cr *checkInOutRepo) CheckRequest(attendanceID string) (checkinEntities.Checkin, error) {
	var checkRequest checkinEntities.Checkin

	row := cr.db.QueryRow(`
	SELECT
		attendances.id, attendances.user_id, attendances.status
	FROM
		checkins
	RIGHT JOIN
		attendances ON attendances.id = checkins.attendance_id
	WHERE
			attendances.id = ?
	GROUP BY
		attendances.id`, attendanceID)

	err := row.Scan(&checkRequest.Attendance.ID, &checkRequest.Attendance.Employee.ID, &checkRequest.Attendance.Status)
	if err != nil {
		return checkRequest, err
	}

	return checkRequest, nil
}

func (cr *checkInOutRepo) CheckData(userID string, attendanceID string) int {
	row, err := cr.db.Query(`
	SELECT
		COUNT(attendances.id)
	FROM
		attendances
	JOIN
		checkins ON checkins.attendance_id = attendances.id
	WHERE
			attendances.user_id = ?
		AND
			checkins.attendance_id = ?
		AND 
			attendances.status = "approved"`, userID, attendanceID)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	var count int

	for row.Next() {
		if err := row.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count
}

func (cr *checkInOutRepo) CheckIn(checkin checkinEntities.Checkin) (checkinEntities.Checkin, error) {
	query := `INSERT INTO checkins (id, attendance_id, temprature, is_checkins, is_checkouts, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	statement, err := cr.db.Prepare(query)
	if err != nil {
		return checkin, err
	}

	_, errExec := statement.Exec(checkin.ID, checkin.Attendance.ID, checkin.Temprature, checkin.IsCheckIns, checkin.IsCheckOuts, checkin.CreatedAt, checkin.UpdatedAt)
	if errExec != nil {
		return checkin, errExec
	}

	return checkin, nil
}

func (cr *checkInOutRepo) CheckOut(userID string, checkinout checkinEntities.Checkin) (checkinEntities.Checkin, error) {
	query := `
	UPDATE
		checkins
	SET
		is_checkouts = ?,
		updated_at = ?
	WHERE
		id = ?
	AND 
		attendance_id IN(
			SELECT
				id FROM attendances
			WHERE
				id = ?
			AND 
				user_id = ?)`

	statement, err := cr.db.Prepare(query)
	if err != nil {
		return checkinout, err
	}

	result, errExec := statement.Exec(checkinout.IsCheckOuts, checkinout.UpdatedAt, checkinout.ID, checkinout.Attendance.ID, userID)
	if errExec != nil {
		return checkinout, errExec
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return checkinout, errors.New("error")
	}

	return checkinout, nil
}
