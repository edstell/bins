package services

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/edstell/bins/service.recycling-services/domain"
)

// Fetcher implementations should retrieve the latest services available for
// the propertyID provided.
type Fetcher interface {
	Fetch(context.Context) ([]domain.Service, error)
}

type FetcherFunc func(context.Context) ([]domain.Service, error)

func (f FetcherFunc) Fetch(ctx context.Context) ([]domain.Service, error) {
	return f(ctx)
}

// WebScraper returns a Fetcher implementation which will fetch the latest
// service information by scraping the recycling services website.
func WebScraper(client *http.Client, parser Parser, url string) Fetcher {
	return FetcherFunc(func(ctx context.Context) ([]domain.Service, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req = req.WithContext(ctx)
		rsp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer rsp.Body.Close()
		bytes, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return nil, err
		}
		return parser.Parse(bytes)
	})
}
