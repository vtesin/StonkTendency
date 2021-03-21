package sentiment

// Source sentiment source to parse. Will specify the URL to fecth the data, placeholder param and limit
type Source struct {
	// URL of the site to scrape information from
	Url string
	// Query string parameter to replace in Url with ticker symbol
	Placeholder string
	// Limit number of documents retreived
	Limit int16
}
