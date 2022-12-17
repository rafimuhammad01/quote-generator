package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r repository) GetQuote(ctx context.Context, generator GenerateQuote) (*Quote, error) {
	var quote Quote

	queryStr := "SELECT id, sentences, number_of_people FROM quote WHERE number_of_people = $1 ORDER BY RANDOM() LIMIT 1"

	err := r.db.Get(&quote, queryStr, generator.NumberOfPeople)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, GenerateError(ErrDataNotFound, fmt.Sprintf("data with number of people %d doesn't exist", generator.NumberOfPeople))
		}

		return nil, GenerateError(ErrInternalServer, err.Error())
	}

	return &quote, nil
}
