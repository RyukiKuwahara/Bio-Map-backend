package repositories

import (
	"database/sql"
	"log"
	"os"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// UserRepository handles user data operations
type UserRepository struct {
	db *sql.DB
}

// GetMaxUserID returns the maximum user ID from the database
func (ur *UserRepository) GetMaxUserID() (int, error) {
	var maxID int
	query := "SELECT MAX(id) FROM users"

	err := ur.db.QueryRow(query).Scan(&maxID)
	if err != nil {
		// If there are no users in the database, return 0 as the maximum ID
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return maxID, nil
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository() (*UserRepository, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Establish a database connection
	db, err := sql.Open("postgres", "host="+host+" port="+port+" user="+dbUser+" password="+password+" dbname="+dbName)
	if err != nil {
		log.Fatal(err)
	}

	// UserRepository objectを生成して返す
	return &UserRepository{
		db: db,
	}, nil
}

// SaveUser saves a user in the database
func (ur *UserRepository) SaveUser(user models.User) error {
	maxID, err := ur.GetMaxUserID() // Get the maximum user ID
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)"
	_, err = ur.db.Exec(query, maxID+1, user.Username, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("User saved successfully")

	return nil
}
