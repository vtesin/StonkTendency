// Deals with sentiment analysis of a ticker
package sentiment

import (
	"github.com/jonreiter/govader"
)

// Service collector
type Service struct {
	// Sources
	src []*Source
	// Analyzer
	anz *govader.SentimentIntensityAnalyzer
}

// NewService create a new sentiment service, with multiple options to read from as well as the sentiment implementation to use
func NewService(sources []*Source, analyzer *govader.SentimentIntensityAnalyzer) *Service {

	return &Service{
		src: sources,
		anz: analyzer,
	}
}

// GetSentiment parses the sources injected
func (svc *Service) GetSentiment(symbol string) (float64, error) {
	sourceCount := 0
	scoreCompound := 0.0

	for i := 0; i < len(svc.src); i++ {
		score, err := svc.src[i].Visit(symbol, svc.anz)

		if err != nil {
			return 0, err
		}

		scoreCompound += score
		sourceCount++
	}

	return scoreCompound / float64(sourceCount), nil
}
