package repositories

import (
	"fmt"
	_ "github.com/lib/pq"
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

func (ur *UserRepository) CreateUsers() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			UNIQUE (username),
			UNIQUE (email)
		);
	`

	_, err := ur.db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("users テーブルが作成されました")
	return nil
}

func (ur *UserRepository) CreateSession() error {
	query := `
	CREATE TABLE session (
		id SERIAL PRIMARY KEY, -- 自動生成の一意のID
		sessionId VARCHAR(32) NOT NULL, -- セッションID（32文字の文字列）
		-- 他のセッション情報を格納する列を追加できます
		created_at TIMESTAMP DEFAULT current_timestamp
		);
	`

	_, err := ur.db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("session テーブルが作成されました")
	return nil
}
