package database

import (
	"context"
	database "main/database/db"

	_ "github.com/lib/pq"
)

//go:generate sqlc generate
type PostgreSQLC struct {
	db *database.Queries
}

func NewPostgreSQLC(db *database.Queries) (*PostgreSQLC, error) {
	return &PostgreSQLC{
		db: db,
	}, nil
}

func (p *PostgreSQLC) FindUserByMSSV(mssv string) (User, error) {
	row, err := p.db.FindUserByMssv(context.Background(), mssv)

	if err != nil {
		return User{}, err
	}

	return User{
		Id:   int(row.ID),
		Name: row.Name,
		Age:  int(row.Age),
		Mssv: row.Mssv,
	}, nil
}
