package services

import (
	"bytes"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Parser implementations should extract services from raw bytes.
type Parser interface {
	Parse([]byte) ([]*recyclingservicesproto.Service, error)
}

type ParserFunc func([]byte) ([]*recyclingservicesproto.Service, error)

func (f ParserFunc) Parse(b []byte) ([]*recyclingservicesproto.Service, error) {
	return f(b)
}

// ParseHTML will extract services from webpage HTML.
var ParseHTML = ParserFunc(func(html []byte) ([]*recyclingservicesproto.Service, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, err
	}

	sel := doc.Find(".results-table-wrapper")
	services := []*recyclingservicesproto.Service{}
	sel.Find(".service-wrapper").Each(func(_ int, sel *goquery.Selection) {
		service := &recyclingservicesproto.Service{}
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
					service.LastService = timestamppb.New(last)
				case strings.Contains(text, "Next Service"):
					next, err := time.Parse(
						"02/01/2006",
						strings.TrimSpace(strings.ReplaceAll(text, "Next Service", "")),
					)
					if err != nil {
						return
					}
					service.NextService = timestamppb.New(next)
				}
			})
			return i == 0
		})
		services = append(services, service)
	})

	return services, nil
})
