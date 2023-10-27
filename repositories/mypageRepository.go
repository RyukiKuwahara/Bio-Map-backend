package repositories

import (
	"github.com/RyukiKuwahara/Bio-Map/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func (ur *UserRepository) GetName(userId int) (string, error) {
	query := "SELECT users.username FROM users WHERE user_id = $1"

	row := ur.db.QueryRow(query, userId)

	var name string
	err := row.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (ur *UserRepository) GetUserPosts(userId int) ([]models.Post, error) {
	query := `
        SELECT posts.post_id, species.species_name, posts.image_path, posts.explain, posts.lat, posts.lng
        FROM posts
        INNER JOIN species ON posts.species_id = species.species_id
        WHERE posts.user_id = $1
    `
	rows, err := ur.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		err := rows.Scan(
			&post.PostId,
			&post.SpeciesName,
			&post.ImagePath,
			&post.Explain,
			&post.Lat,
			&post.Lng,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
