package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type PostgresDBsqlx struct {
	db *sqlx.DB
}

func NewPostgresDBsqlx() (*PostgresDBsqlx, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	db, err := sqlx.ConnectContext(context.Background(), "postgres", os.Getenv("DATABASE_URL"))

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

	fmt.Println("PostgresSQL connected successfully!")

	return &PostgresDBsqlx{
		db: db,
	}, nil
}

func (db *PostgresDBsqlx) Close() {
	db.db.Close()
}

func (m *PostgresDBsqlx) FindBYMSSV(mssv string) (User, error) {
	query := `SELECT id, name, age, mssv FROM users WHERE mssv = ?`

	var res User

	err := m.db.QueryRowx(query, mssv).StructScan(&res)

	if err != nil {
		return User{}, fmt.Errorf("failed to find user: %w", err)
	}

	return res, nil
}
