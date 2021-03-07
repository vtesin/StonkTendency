package ticker

import "github.com/vtesin/StonkTendency/entity"

//Service  interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateStonk Create a ticker detail, and perform sentiment analysis
func (s *Service) CreateStonk(symbol string, sentiment float32) error {
	t, err := entity.NewTicker(symbol)
	if err != nil {
		return err
	}

	t.Sentiment = 0.0

	// TODO add sentiment analysis and store it in Mongo
	return nil
}

// GetStonk Find a ticker
func (s *Service) GetStonk(symbol string) (*entity.Ticker, error) { return nil, nil }

// SearchStonks Search for a ticker based on a criteria
func (s *Service) SearchStonks(query string) ([]*entity.Ticker, error) { return nil, nil }

// ListStonks Lists all tickers in the system
func (s *Service) ListStonks() ([]*entity.Ticker, error) { return nil, nil }

// UpdateStonk Update a ticker
func (s *Service) UpdateStonk(t *entity.Ticker) error { return nil }

// DeleteStonk Deletes a ticker
func (s *Service) DeleteStonk(symbol string) error { return nil }
