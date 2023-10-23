package repositories

import (
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
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
		session_id VARCHAR(32) PRIMARY KEY, -- セッションID（32文字の文字列）
		user_id INT,
		created_at TIMESTAMP DEFAULT current_timestamp,
		FOREIGN KEY (user_id) REFERENCES users (user_id)
		);
	`

	_, err := ur.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("session テーブルが作成されました")
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

	fmt.Println("genresの行が追加されました")
}

func (ur *UserRepository) CreateSpecies() {
	query := `
		CREATE TABLE IF NOT EXISTS species (
			species_id SERIAL PRIMARY KEY,
			species_name TEXT,
			genre_id INT,
			FOREIGN KEY (genre_id) REFERENCES genres (genre_id)
		);
	`

	_, err := ur.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("species テーブルが作成されました")

	file, err := os.Open("./setups/species.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	query = "INSERT INTO species (species_name, genre_id) VALUES ($1, $2);"

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		species_name := record[0]
		genre_id := record[1]

		_, err = ur.db.Exec(query, species_name, genre_id)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("species の行が追加されました")
}

func (ur *UserRepository) CreatePosts() {
	query := `
		CREATE TABLE IF NOT EXISTS posts (
			post_id SERIAL PRIMARY KEY,
			user_id INT,
			species_id INT,
			lat DOUBLE PRECISION,
			lng DOUBLE PRECISION,
			image_path VARCHAR(255),
			explain TEXT,
			time_stamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users (user_id),
			FOREIGN KEY (species_id) REFERENCES species (species_id)
		);
	`

	_, err := ur.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("posts テーブルが作成されました")
}
