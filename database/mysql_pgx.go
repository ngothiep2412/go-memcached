package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Mssv string `json:"mssv"`
}

type MySQLDB struct {
	db *sql.DB
}

func NewMySQLDB() (*MySQLDB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("MySQL connected successfully!")

	return &MySQLDB{
		db: db,
	}, nil
}

func (db *MySQLDB) Close() {
	db.db.Close()
}

func (m *MySQLDB) FindBYMSSV(mssv string) (User, error) {
	query := `SELECT id, name, age, mssv FROM users WHERE mssv = ?`

	var res User

	err := m.db.QueryRow(query, mssv).Scan(&res.Id, &res.Name, &res.Age, &res.Mssv)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user with mssv '%s' not found", mssv)
		}
		return User{}, fmt.Errorf("failed to find user: %w", err)
	}

	return res, nil
}
