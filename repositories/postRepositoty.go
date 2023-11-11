package repositories

import (
	"database/sql"
	"fmt"

	"github.com/RyukiKuwahara/Bio-Map/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func (ur *UserRepository) GetUserId(session_id string) (int, error) {
	fmt.Println(session_id)
	query := "SELECT user_id FROM session WHERE session_id = $1"
	row := ur.db.QueryRow(query, session_id)

	var userId int
	err := row.Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, fmt.Errorf("session not found")
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
			return -1, fmt.Errorf("登録された名前はデータベースに存在しません")
		}
		return -1, err
	}
	return speciesId, nil
}

func (ur *UserRepository) RegisterPost(pr models.PostRequest, userId, speciesId int, imagePath string) error {
	query := "INSERT INTO posts (user_id, species_id, lat, lng, image_path, explain) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := ur.db.Exec(query, userId, speciesId, pr.Lat, pr.Lng, imagePath, pr.Explain)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) CountOverlapping(userId, speciesId int) (int, error) {

	query := `SELECT COUNT(user_id) FROM posts WHERE user_id = $1 and species_id = $2`
	row := ur.db.QueryRow(query, userId, speciesId)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (ur *UserRepository) GetGenreId(speciesId int) (int, error) {

	query := `SELECT genre_id FROM species WHERE species_id = $1`
	row := ur.db.QueryRow(query, speciesId)

	var genreId int
	err := row.Scan(&genreId)
	if err != nil {
		return -1, err
	}
	return genreId, nil
}

func (ur *UserRepository) CountPosts(userId, genreId int) (int, error) {

	query := `
		SELECT COUNT(DISTINCT posts.species_id) FROM posts
		INNER JOIN species ON posts.species_id = species.species_id
		WHERE user_id = $1 and genre_id = $2
	`
	row := ur.db.QueryRow(query, userId, genreId)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (ur *UserRepository) RegisterBadge(userId, badgeId int) error {
	query := "INSERT INTO user_badge_history (user_id, badge_id) VALUES ($1, $2)"
	_, err := ur.db.Exec(query, userId, badgeId)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetBadgesPath(badgeId int) (string, error) {

	query := `SELECT badge_path FROM badge WHERE badge_id = $1`
	row := ur.db.QueryRow(query, badgeId)

	var badgePath string
	err := row.Scan(&badgePath)
	if err != nil {
		return "", err
	}
	return badgePath, nil
}
