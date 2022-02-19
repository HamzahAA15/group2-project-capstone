package officeRepository

import (
	"database/sql"
	"sirclo/project-capstone/entities/officeEntities"
)

type officeRepo struct {
	db *sql.DB
}

func NewMySQLOfficeRepository(db *sql.DB) OfficeRepoInterface {
	return &officeRepo{
		db: db,
	}
}

func (or *officeRepo) GetOffices() ([]officeEntities.Office, error) {
	var offices []officeEntities.Office

	result, err := or.db.Query(`SELECT id, name, created_at, updated_at FROM offices`)
	if err != nil {
		return offices, err
	}

	for result.Next() {
		var office officeEntities.Office

		err = result.Scan(&office.ID, &office.Name, &office.CreatedAt, &office.UpdatedAt)
		if err != nil {
			return offices, err
		}

		offices = append(offices, office)
	}
	return offices, nil
}
