package entity

import "time"

// Ticker data
type Ticker struct {
	Symbol    string
	Sentiment float32
	Timestamp time.Time
}

// NewTicker create a new ticker
func NewTicker(symbol string) (*Ticker, error) {
	t := &Ticker{
		Symbol:    symbol,
		Sentiment: 0.0,
		Timestamp: time.Now(),
	}

	err := t.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return t, nil

}

// Validate ticker
func (t *Ticker) Validate() error {
	if t.Symbol == "" {
		return ErrNoTicker
	}

	return nil
}
