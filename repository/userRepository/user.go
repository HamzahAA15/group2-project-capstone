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

func (ur *userRepo) CheckEmail(userChecked userEntities.User) (userEntities.User, error) {
	user := userEntities.User{}
	result, err := ur.db.Query("SELECT email FROM users WHERE email=?", userChecked.Email)
	if err != nil {
		return user, err
	}
	defer result.Close()
	if isExist := result.Next(); isExist {
		return user, fmt.Errorf("user already exist")
	}

	if user.Email != userChecked.Email {
		// usernya belum ada
		return user, nil
	}

	return user, nil
}

func (ur *userRepo) Login(identity string) (userEntities.User, error) {
	row := ur.db.QueryRow(`SELECT id, email, password, role FROM users WHERE nik = ? OR email = ? AND deleted_at IS NULL`, identity, identity)

	var user userEntities.User

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepo) GetUser(id string) (userEntities.User, error) {
	var user userEntities.User

	row := ur.db.QueryRow(`
		SELECT 
			users.id, offices.name, users.avatar, users.nik, users.email, users.name, users.phone, users.role, users.created_at 
		FROM 
			users 
		JOIN
			offices ON offices.id = users.office_id
		WHERE 
			users.id = ? AND users.deleted_at IS NULL`, id)

	err := row.Scan(&user.ID, &user.Office.Name, &user.Avatar, &user.Nik, &user.Email, &user.Name, &user.Phone, &user.Role, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepo) CreateUser(user userEntities.User) (userEntities.User, error) {
	query := `INSERT INTO users (id, avatar, nik, email, password, name, phone, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	statement, err := ur.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.ID, user.Avatar, user.Nik, user.Email, user.Password, user.Name, user.Phone, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepo) UpdateUser(user userEntities.User) (userEntities.User, error) {
	query := `UPDATE users SET nik = ?, email = ?, password = ?, name = ?, phone = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL`

	statement, err := ur.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Nik, user.Email, user.Password, user.Name, user.Phone, user.UpdatedAt, user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (us *userRepo) UploadAvatarUser(user userEntities.User) error {
	query := `UPDATE users SET avatar = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL`

	statement, err := us.db.Prepare(query)
	if err != nil {
		return err
	}

	result, errExec := statement.Exec(user.Avatar, user.UpdatedAt, user.ID)
	if errExec != nil {
		return errExec
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil

}
