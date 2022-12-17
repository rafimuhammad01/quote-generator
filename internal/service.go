package internal

import (
	"context"
)

type Repository interface {
	GetQuote(ctx context.Context, quote GenerateQuote) (*Quote, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GenerateQuote(ctx context.Context, input GenerateQuote) (*Quote, error) {
	// validate input
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// get matched value num of people and names
	input, err := input.MatchNumberOfPeopleAndNames()
	if err != nil {
		return nil, err
	}

	// get data from database
	resp, err := s.repository.GetQuote(ctx, input)
	if err != nil {
		return nil, err
	}

	// prepare names
	sentence, names := resp.ParseNameToStrFormat(input.Names)
	resp.Sentences = sentence

	// parsing to name
	namedString, err := resp.MatchSentencesWithNames(input.Names, names)
	if err != nil {
		return nil, err
	}
	resp.Sentences = namedString

	return resp, nil
}
