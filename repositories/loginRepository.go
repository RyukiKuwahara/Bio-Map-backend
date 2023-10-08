package repositories

import (
	"fmt"
	"github.com/RyukiKuwahara/Bio-Map/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func (ur *UserRepository) CheckUser(user models.SigninUser) error {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2"
	rows, err := ur.db.Query(query, user.Username, user.Password)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return err
		}
	}

	if count == 1 {
		return nil
	} else {
		return fmt.Errorf("user not found or multiple users exist")
	}
}
