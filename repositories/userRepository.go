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
func (ur *UserRepository) SaveUser(user models.SignupUser) error {

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	_, err := ur.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
