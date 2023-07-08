package repositories

import (
	"database/sql"
	"log"
	"os"

	"github.com/RyukiKuwahara/Bio-Map/backend/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// SaveUser saves a user in the database
func SaveUser(user models.User) error {
	// Implement the logic to save the user in the database

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	db, err := sql.Open("postgres", "host="+host+" port="+port+" user="+dbuser+" password="+password+" dbname="+dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(query, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("User saved successfully")
	}

	return nil
}
