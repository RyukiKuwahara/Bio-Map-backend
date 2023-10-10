package repositories

import (
	"database/sql"
	"fmt"
	"github.com/RyukiKuwahara/Bio-Map/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func (ur *UserRepository) CheckUser(user models.SigninUser) (string, error) {
	query := "SELECT id FROM users WHERE username = $1 AND password = $2"
	row := ur.db.QueryRow(query, user.Username, user.Password)

	var userId string
	err := row.Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}

	return userId, nil
}

func (ur *UserRepository) RegisterSessionId(sessionId, userId string) error {

	query := "DELETE FROM session WHERE id = $1"
	_, err := ur.db.Exec(query, userId)
	if err != nil {
		return err
	}

	query = "INSERT INTO session (id, sessionId) VALUES ($1, $2)"
	_, err = ur.db.Exec(query, userId, sessionId)
	if err != nil {
		return err
	}
	return nil
}
