package checkInOutRepository

import (
	"database/sql"
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
		id, attendance_id, temprature, is_checkins, is_checkouts, created_at, updated_at 
	FROM 
		checkins
	ORDER BY 
		created_at DESC`)
	if err != nil {
		return checkinsouts, err
	}

	for result.Next() {
		var checkinout checkinEntities.Checkin

		errScan := result.Scan(&checkinout.ID, &checkinout.AttendanceID, &checkinout.Temprature, &checkinout.IsCheckIns, &checkinout.IsCheckOuts, &checkinout.CreatedAt, &checkinout.UpdatedAt)

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
		checkins.id, checkins.attendance_id, checkins.temprature, checkins.is_checkins, checkins.is_checkouts,checkins. created_at, checkins.updated_at
	FROM 
		checkins
	JOIN
		attendances ON attendances.id = checkins.attendance_id
	WHERE
		attendances.user_id = ?
	ORDER BY 
		created_at DESC`, userID)
	if err != nil {
		return checkinsouts, err
	}

	for result.Next() {
		var checkinout checkinEntities.Checkin

		errScan := result.Scan(&checkinout.ID, &checkinout.AttendanceID, &checkinout.Temprature, &checkinout.IsCheckIns, &checkinout.IsCheckOuts, &checkinout.CreatedAt, &checkinout.UpdatedAt)

		if errScan != nil {
			return checkinsouts, errScan
		}

		checkinsouts = append(checkinsouts, checkinout)
	}

	return checkinsouts, nil
}
