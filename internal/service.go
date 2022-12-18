package internal

import (
	"context"
)

type Repository interface {
	GetQuote(ctx context.Context, quote GenerateQuote) (*Quote, error)
	GetQuoteById(ctx context.Context, id int) (*Quote, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GenerateQuote(ctx context.Context, input GenerateQuote) (*Quote, error) {
	// sanitize input
	input = input.Sanitize()

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

	// parsing to name
	namedString, err := resp.MatchSentencesWithNames(input.Names)
	if err != nil {
		return nil, err
	}
	resp.Sentences = namedString

	return resp, nil
}

func (s *service) ShuffleQuote(ctx context.Context, input ShuffleQuote) (*Quote, error) {
	// validate input
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// get data from database
	resp, err := s.repository.GetQuoteById(ctx, input.QuoteId)
	if err != nil {
		return nil, err
	}

	// shuffle names
	input, err = input.ShuffleNames()
	if err != nil {
		return nil, err
	}

	// parsing to name
	namedString, err := resp.MatchSentencesWithNames(input.Names)
	if err != nil {
		return nil, err
	}
	resp.Sentences = namedString

	return resp, nil
}
