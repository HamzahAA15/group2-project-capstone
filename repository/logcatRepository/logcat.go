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
