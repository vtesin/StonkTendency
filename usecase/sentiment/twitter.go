package sentiment

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jonreiter/govader"
	"github.com/vtesin/StonkTendency/entity"
)

type Twitter struct {
	Source
}

type TweetStream struct {
	Data []Tweet   `json:"data"`
	Meta TweetMeta `json:"meta"`
}

type Tweet struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	Lang      string `json:"lang"`
	AuthorId  string `json:"author_id"`
}

type TweetMeta struct {
	NewestId    string `json:"newest_id"`
	OldestId    string `json:"oldest_id"`
	ResultCount uint   `json:"result_count"`
	NextToken   string `json:"next_token"`
}

// NewTwitter creates a new twitter source. Uses Twitter API v2 to find tweets matching the query up-to seven days
func NewTwitter(url string, placeholder string, limit int16) *Twitter {
	return &Twitter{
		Source{
			Url:         url,
			Placeholder: placeholder,
			Limit:       limit,
		},
	}
}

// Visit Twitter API and perform the tweet sentiment analysis for a given symbol
func (t *Source) Visit(symbol string, a *govader.SentimentIntensityAnalyzer) (float64, error) {
	// TODO add error handling
	apiUrl, _ := t.createUrl(symbol)

	auth, _ := t.auth()

	client := &http.Client{}

	req, _ := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", strings.Replace("Bearer $BEARER_TOKEN", "$BEARER_TOKEN", auth, 1))

	resp, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	if resp.StatusCode != 200 {
		return 0, entity.ErrTwitterApiBadRequest
	}

	body := ""
	buffer := make([]byte, 2048)

	for {
		n, err := resp.Body.Read(buffer)

		if n == 0 && err == io.EOF {
			break // end of our response payload
		} else if err != nil {
			return 0, err
		}

		body += string(buffer[:n])
	}

	var stream TweetStream
	sentimentCount := 0
	totalCompound := 0.0

	err = json.Unmarshal([]byte(body), &stream)

	if err != nil {
		log.Printf("Error: %v", err)
		log.Printf("Body: %s", body)

		return 0, err
	}

	for i := 0; i < len(stream.Data); i++ {
		score := a.PolarityScores(stream.Data[i].Text)
		sentimentCount++
		totalCompound += score.Compound
	}
	// TODO add next token handling
	return totalCompound / float64(sentimentCount), nil
}

// createUrl replaces the placeholder with the ticker symbol
func (t *Source) createUrl(symbol string) (string, error) {
	return strings.Replace(t.Url, t.Placeholder, symbol, 1), nil
}

// auth retreives the BEARER_TOKEN from OS env
func (t *Source) auth() (string, error) {
	return os.Getenv("BEARER_TOKEN"), nil
}
