package ticker

import (
	"github.com/vtesin/StonkTendency/entity"
	"github.com/vtesin/StonkTendency/usecase/sentiment"
)

//Service  interface
type Service struct {
	repo             Repository
	sentimentService sentiment.UseCase
}

//NewService create new use case
func NewService(r Repository, s sentiment.UseCase) *Service {
	return &Service{
		repo:             r,
		sentimentService: s,
	}
}

// CreateStonk Create a ticker detail, and perform sentiment analysis
func (s *Service) CreateStonk(symbol string, sentiment float64) error {
	t, err := entity.NewTicker(symbol)
	if err != nil {
		return err
	}

	//TODO remove sentiment param - not needed
	compound, err := s.sentimentService.GetSentiment(symbol)

	if err != nil {
		return err
	}

	t.Sentiment = compound
	err = s.repo.Update(t)

	return err
}

// GetStonk Find a ticker
func (s *Service) GetStonk(symbol string) (*entity.Ticker, error) {
	t, err := s.repo.Get(symbol)
	return t, err
}

// SearchStonks Search for a ticker based on a criteria
func (s *Service) SearchStonks(query string) ([]*entity.Ticker, error) { return nil, nil }

// ListStonks Lists all tickers in the system
func (s *Service) ListStonks() ([]*entity.Ticker, error) { return nil, nil }

// UpdateStonk Update a ticker
func (s *Service) UpdateStonk(t *entity.Ticker) error { return nil }

// DeleteStonk Deletes a ticker
func (s *Service) DeleteStonk(symbol string) error { return nil }
