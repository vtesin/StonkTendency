package sentiment

// UseCase interface
type UseCase interface {
	GetSentiment(symbol string) (float64, error)
}
