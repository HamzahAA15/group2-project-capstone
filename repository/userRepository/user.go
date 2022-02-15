package userRepository

import (
	"database/sql"
	"fmt"
	"sirclo/project-capstone/entities/userEntities"

	_ "github.com/go-sql-driver/mysql"
)

type userRepo struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) UserRepoInterface {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) GetUsers() ([]userEntities.User, error) {
	var users []userEntities.User

	result, err := ur.db.Query(`SELECT id, username, email, password, avatar, created_at FROM users WHERE deleted_at IS NULL`)
	if err != nil {
		return users, err
	}

	for result.Next() {
		var user userEntities.User

		err = result.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.CreatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *userRepo) GetUser(id string) (userEntities.User, error) {
	var user userEntities.User

	row := ur.db.QueryRow(`SELECT id, username, email, password, avatar, created_at FROM users WHERE id = ? AND deleted_at IS NULL`, id)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepo) DeleteUser(loginId string) error {
	query := `UPDATE users SET deleted_at = now() WHERE id = ? AND deleted_at IS NULL`

	statement, err := ur.db.Prepare(query)
	if err != nil {
		return err
	}

	result, errExec := statement.Exec(loginId)
	if errExec != nil {
		return errExec
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("event not found")
	}
	return nil

}
