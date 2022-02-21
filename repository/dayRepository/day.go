package dayRepository

import (
	"database/sql"
	"sirclo/project-capstone/entities/dayEntities"
)

type dayRepo struct {
	db *sql.DB
}

func NewMySQLDayRepository(db *sql.DB) DayRepoInterface {
	return &dayRepo{
		db: db,
	}
}

func (dr *dayRepo) GetDays() ([]dayEntities.Day, error) {
	var days []dayEntities.Day

	result, err := dr.db.Query(`SELECT id, office_id, date, quota FROM days`)
	if err != nil {
		return days, err
	}

	for result.Next() {
		var day dayEntities.Day

		errScan := result.Scan(&day.ID, &day.OfficeId, &day.Date, &day.Quota)
		if errScan != nil {
			return days, errScan
		}
		days = append(days, day)
	}
	return days, nil
}
