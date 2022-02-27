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

func (dr *dayRepo) GetDays(office, time string) ([]dayEntities.Day, error) {
	var days []dayEntities.Day
	convOffice := "%" + office + "%"
	convTime := "%" + time + "%"

	result, err := dr.db.Query(`
	Select days.id, offices.name, days.date, days.quota,
	count(attendances.day_id) AS total_approved, 
	(days.quota-count(attendances.day_id)) AS remaining_quota 
	FROM days
    LEFT JOIN (SELECT attendances.day_id FROM attendances WHERE attendances.status = "approved") attendances ON attendances.day_id = days.id
    LEFT JOIN offices ON offices.id = days.office_id
    WHERE offices.id LIKE ? AND days.date LIKE ?
    GROUP BY days.id ORDER BY days.date ASC`, convOffice, convTime)

	if err != nil {
		return days, err
	}
	for result.Next() {
		var day dayEntities.Day

		errScan := result.Scan(&day.ID, &day.OfficeId.Name, &day.Date, &day.Quota, &day.TotalApproved, &day.RemainingQuota)
		if errScan != nil {
			return days, errScan
		}

		days = append(days, day)
	}

	return days, nil
}

func (dr *dayRepo) UpdateDay(day dayEntities.Day) (dayEntities.Day, error) {
	query := `UPDATE days SET quota = ?, updated_at = now() WHERE id = ?`

	statement, err := dr.db.Prepare(query)
	if err != nil {
		return day, err
	}

	defer statement.Close()

	_, err = statement.Exec(day.Quota, day.ID)
	if err != nil {
		return day, err
	}

	return day, nil
}
