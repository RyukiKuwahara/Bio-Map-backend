package repositories

import (
	"github.com/RyukiKuwahara/Bio-Map/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func (ur *UserRepository) GetPosts(name string) ([]models.Post, error) {
	query := `
        SELECT posts.post_id, species.species_name, posts.image_path, posts.explain, posts.lat, posts.lng
        FROM posts
        INNER JOIN species ON posts.species_id = species.species_id
        WHERE species.species_name = $1
    `
	rows, err := ur.db.Query(query, name)
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
