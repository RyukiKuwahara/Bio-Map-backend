package repositories

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
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

func (ur *UserRepository) CreateUsers() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			user_id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			UNIQUE (username),
			UNIQUE (email)
		);
	`

	_, err := ur.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("users テーブルが作成されました")
}

func (ur *UserRepository) CreateSession() {
	query := `
	CREATE TABLE session (
		user_id SERIAL PRIMARY KEY, -- 自動生成の一意のID
		session_id VARCHAR(32) NOT NULL, -- セッションID（32文字の文字列）
		-- 他のセッション情報を格納する列を追加できます
		created_at TIMESTAMP DEFAULT current_timestamp
		);
	`

	_, err := ur.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("session テーブルが作成されました")
}

func (ur *UserRepository) CreatePosts() {
	query := `
		CREATE TABLE IF NOT EXISTS posts (
			post_id SERIAL PRIMARY KEY,
			user_id INT,
			species_id INT,
			image_path VARCHAR(255),
			explain TEXT,
			time_stamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := ur.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("posts テーブルが作成されました")
}

func (ur *UserRepository) CreateGenres() {
	query := `
		CREATE TABLE IF NOT EXISTS genres (
			genre_id SERIAL PRIMARY KEY,
			genre_name TEXT
		);
	`

	_, err := ur.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("genres テーブルが作成されました")

	query = `
		INSERT INTO genres (genre_id, genre_name) VALUES
		(1, '虫'),
		(2, '魚'),
		(3, '植物');
	`

	_, err = ur.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("行が追加されました")
}
