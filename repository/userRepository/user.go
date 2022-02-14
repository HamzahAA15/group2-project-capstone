package userRepository

import (
	"database/sql"
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

func (ur *userRepo) GetUser(id string) (userEntities.User, error) {
	var user userEntities.User

	row := ur.db.QueryRow(`SELECT id, username, email, password, avatar, created_at FROM users WHERE id = ?`, id)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
