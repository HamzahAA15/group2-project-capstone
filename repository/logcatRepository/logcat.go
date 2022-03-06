package logcatRepository

import (
	"database/sql"
	"sirclo/project-capstone/entities/logcatEntities"
)

type logcatRepo struct {
	db *sql.DB
}

func NewMySQLLogcatRepository(db *sql.DB) LogcatRepoInterface {
	return &logcatRepo{
		db: db,
	}
}

func (lr *logcatRepo) CreateLogcat(lc logcatEntities.Logcat) (logcatEntities.Logcat, error) {
	query := `INSERT INTO logcats (id, user_id, message, category, created_at) VALUES (?, ?, ?, ?, now())`

	statement, err := lr.db.Prepare(query)
	if err != nil {
		return lc, err
	}
	_, errExec := statement.Exec(lc.ID, lc.User.ID, lc.Message, lc.Category)
	if errExec != nil {
		return lc, errExec
	}
	return lc, nil
}

func (lr *logcatRepo) GetLogcats() ([]logcatEntities.Logcat, error) {
	var logcats []logcatEntities.Logcat

	result, err := lr.db.Query(`
	SELECT 
		id, message, category, created_at 
	FROM 
		logcats
	ORDER BY 
		created_at DESC`)
	if err != nil {
		return logcats, err
	}

	for result.Next() {
		var logcat logcatEntities.Logcat

		err = result.Scan(&logcat.ID, &logcat.Message, &logcat.Category, &logcat.CreatedAt)
		if err != nil {
			return logcats, err
		}

		logcats = append(logcats, logcat)
	}

	return logcats, nil
}

func (lr *logcatRepo) GetLogcatUser(userID string) ([]logcatEntities.Logcat, error) {
	var logcats []logcatEntities.Logcat

	result, err := lr.db.Query(`
	SELECT 
		id, message, category, created_at 
	FROM 
		logcats
	WHERE
		user_id = ?
	ORDER BY 
		created_at DESC`, userID)
	if err != nil {
		return logcats, err
	}

	for result.Next() {
		var logcat logcatEntities.Logcat

		err = result.Scan(&logcat.ID, &logcat.Message, &logcat.Category, &logcat.CreatedAt)
		if err != nil {
			return logcats, err
		}

		logcats = append(logcats, logcat)
	}

	return logcats, nil
}
