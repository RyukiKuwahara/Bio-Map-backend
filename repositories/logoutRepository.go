package repositories

import (
	_ "github.com/lib/pq"
)

func (ur *UserRepository) RemoveSessionId(sessionId string) error {

	query := "DELETE FROM session WHERE session_id = $1"
	_, err := ur.db.Exec(query, sessionId)
	if err != nil {
		return err
	}
	return nil
}
