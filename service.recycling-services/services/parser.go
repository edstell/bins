package services

import (
	"bytes"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
)

// Parser implementations should extract services from raw bytes.
type Parser interface {
	Parse([]byte) ([]recyclingservices.Service, error)
}

type ParserFunc func([]byte) ([]recyclingservices.Service, error)

func (f ParserFunc) Parse(b []byte) ([]recyclingservices.Service, error) {
	return f(b)
}

// ParseHTML will extract services from webpage HTML.
var ParseHTML = ParserFunc(func(html []byte) ([]recyclingservices.Service, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, err
	}

	sel := doc.Find(".results-table-wrapper")
	services := []recyclingservices.Service{}
	sel.Find(".service-wrapper").Each(func(_ int, sel *goquery.Selection) {
		service := recyclingservices.Service{}
		service.Name = strings.TrimSpace(sel.Find(".service-name").Text())
		service.Status = strings.TrimPrefix(strings.TrimSpace(sel.Find(".task-state").Text()), "Last collection: ")
		sel.Find("tr").EachWithBreak(func(i int, sel *goquery.Selection) bool {
			sel.Find("td").Each(func(_ int, sel *goquery.Selection) {
				text := sel.Text()
				switch {
				case strings.Contains(text, "Schedule"):
					service.Schedule = strings.TrimSpace(strings.ReplaceAll(text, "Schedule", ""))
				case strings.Contains(text, "Last Service"):
					last, err := time.Parse(
						"02/01/2006",
						strings.TrimSpace(strings.ReplaceAll(text, "Last Service", "")),
					)
					if err != nil {
						return
					}
					service.LastService = last
				case strings.Contains(text, "Next Service"):
					next, err := time.Parse(
						"02/01/2006",
						strings.TrimSpace(strings.ReplaceAll(text, "Next Service", "")),
					)
					if err != nil {
						return
					}
					service.NextService = next
				}
			})
			return i == 0
		})
		services = append(services, service)
	})

	return services, nil
})
