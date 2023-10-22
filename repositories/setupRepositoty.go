package repositories

import (
	_ "github.com/lib/pq"
)

func (ur *UserRepository) TableExits(tableName string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)"
	var exists bool
	err := ur.db.QueryRow(query, tableName).Scan(&exists)

	if err != nil {
		return false, err
	}
	return exists, nil
}

func (ur *UserRepository) CreateUsers() error {
	return nil
}
