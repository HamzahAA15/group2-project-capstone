package schedule

import (
	"log"
	"sirclo/project-capstone/database"
	"sirclo/project-capstone/entities/officeEntities"
	"time"

	"github.com/google/uuid"
)

func GenerateCalender() {
	db := database.MySQLConnection()

	var offices []officeEntities.Office
	querySelect := `SELECT offices.id FROM offices`

	result, err := db.Query(querySelect)
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	for result.Next() {
		var office officeEntities.Office

		errScan := result.Scan(&office.ID)
		if errScan != nil {
			log.Fatal(errScan)
		}

		offices = append(offices, office)
	}

	query := "INSERT INTO days (id, office_id, date, quota, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	t := time.Now()
	firstOfMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	totalDays := lastOfMonth.Day() - firstOfMonth.Day()

	totalWorker := len(offices) * (totalDays + 1)

	wg.Add(totalWorker)
	syn := func(uuid string, officeID string, date time.Time, quota int, createdAt time.Time, updatedAt time.Time) {
		defer wg.Done()
		db.Exec(query, uuid, officeID, date, quota, createdAt, updatedAt)
	}

	for office := 0; office < len(offices); office++ {
		for i := firstOfMonth.Day(); i <= lastOfMonth.Day(); i++ {
			n := time.Date(t.Year(), t.Month()+1, i, 00, 00, 0, 0, t.Location())
			go syn(uuid.New().String(), offices[office].ID, n, 50, time.Now(), time.Now())
		}
	}
	wg.Wait()
}
