package repositories

import (
	"database/sql"
	"fmt"

	"github.com/RyukiKuwahara/Bio-Map/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func (ur *UserRepository) GetUserId(session_id string) (int, error) {
	query := "SELECT user_id FROM session WHERE session_id = $1"
	row := ur.db.QueryRow(query, session_id)

	var userId int
	err := row.Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, fmt.Errorf("user not found")
		}
		return -1, err
	}

	return userId, nil
}

func (ur *UserRepository) GetSpeciesId(name string) (int, error) {

	query := "SELECT species_id FROM species WHERE species_name = $1"
	row := ur.db.QueryRow(query, name)

	var speciesId int
	err := row.Scan(&speciesId)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, fmt.Errorf("species_name not found")
		}
		return -1, err
	}
	return speciesId, nil
}

func (ur *UserRepository) RegisterPost(pr models.PostRequest, userId, speciesId int, imagePath string) error {
	query := "INSERT INTO session (user_id, species_id, lat, lng, image_path, explain) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := ur.db.Exec(query, userId, speciesId, pr.Lat, pr.Lng, imagePath, pr.Explain)
	if err != nil {
		return err
	}
	return nil
}
