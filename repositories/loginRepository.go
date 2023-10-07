package repositories

import (
	"fmt"
	"github.com/RyukiKuwahara/Bio-Map/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func (ur *UserRepository) CheckUser(user models.SigninUser) error {
	var count int
	query := "SELECT (*) FROM user WHERE username = ? AND password = ?"
	rows, err := ur.db.Query(query, user.Username, user.Password)
	fmt.Println(rows, err)
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
	fmt.Println(count)

	if count == 1 {
		return nil
	} else {
		return err
	}
}
