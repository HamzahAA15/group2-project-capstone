package schedule

import (
	"log"
	"sirclo/project-capstone/database"
	"time"
)

func CheckoutAttendance() {
	db := database.MySQLConnection()

	var countData string = `SELECT COUNT(*) FROM checkins WHERE is_checkouts = FALSE`

	result, err := db.Query(countData)
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	var totalWorker int

	for result.Next() {
		if err := result.Scan(&totalWorker); err != nil {
			log.Fatal(err)
		}
	}

	var query string = `UPDATE checkins SET is_checkouts = ?, updated_at = ? WHERE is_checkouts = ?`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	wg.Add(totalWorker)
	syn := func(status bool, time time.Time, checkout bool) {
		defer wg.Done()
		db.Exec(query, status, time, checkout)
	}

	for i := 0; i < totalWorker; i++ {
		go syn(true, time.Now(), false)
	}
	wg.Wait()
}
