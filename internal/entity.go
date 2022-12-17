package internal

import (
	"fmt"
	"strings"
)

var (
	ErrValidationError = fmt.Errorf("input validation error")
	ErrDataNotFound    = fmt.Errorf("data not found")
	ErrBadRequest      = fmt.Errorf("bad request")
	ErrInternalServer  = fmt.Errorf("internal server error")
)

type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type GenerateQuote struct {
	NumberOfPeople int      `json:"number_of_people" query:"number_of_people"`
	Names          []string `json:"names" query:"names"`
}

func (g GenerateQuote) Validate() error {
	if g.NumberOfPeople == 0 {
		return GenerateError(ErrValidationError, "number of people cannot be 0")
	}

	if len(g.Names) == 0 {
		return GenerateError(ErrValidationError, "names cannot be empty")
	}

	if len(g.Names) < g.NumberOfPeople {
		return GenerateError(ErrValidationError, "names must be equal with number of people")
	}

	return nil
}

func (g GenerateQuote) MatchNumberOfPeopleAndNames() (GenerateQuote, error) {
	var data GenerateQuote
	if len(g.Names) < g.NumberOfPeople {
		return data, GenerateError(ErrValidationError, "names must be equal with number of people")
	}

	data = GenerateQuote{
		NumberOfPeople: g.NumberOfPeople,
		Names:          g.Names[:g.NumberOfPeople],
	}

	return data, nil
}

type Quote struct {
	ID             int    `json:"id" db:"id"`
	Sentences      string `json:"sentences" db:"sentences"`
	NumberOfPeople int    `json:"number_of_people" db:"number_of_people"`
}

func (q Quote) MatchSentencesWithNames(names []string) (string, error) {
	if q.NumberOfPeople != len(names) {
		return "", GenerateError(ErrValidationError, "names must be equal with number of people")
	}

	for i, v := range names {
		q.Sentences = strings.ReplaceAll(q.Sentences, fmt.Sprintf("[p%d]", i+1), v)
	}

	return q.Sentences, nil
}
