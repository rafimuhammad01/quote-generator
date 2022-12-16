package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r repository) GetQuote(ctx context.Context, generator GenerateQuote) (*Quote, error) {
	var quote Quote

	queryStr := "SELECT * FROM quote WHERE number_of_people = $1"

	err := r.db.QueryRowContext(ctx, queryStr, generator.NumberOfPeople).Scan(&quote)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, GenerateError(ErrDataNotFound, fmt.Sprintf("data with number of people %d doesn't exist", generator.NumberOfPeople))
		}

		return nil, GenerateError(ErrInternalServer, err.Error())
	}

	return &quote, nil
}
