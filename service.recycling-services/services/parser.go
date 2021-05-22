package services

import (
	"bytes"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/edstell/bins/service.recycling-services/domain"
)

// Parser implementations should extract services from raw bytes.
type Parser interface {
	Parse([]byte) ([]domain.Service, error)
}

type ParserFunc func([]byte) ([]domain.Service, error)

func (f ParserFunc) Parse(b []byte) ([]domain.Service, error) {
	return f(b)
}

// V2Parser
var V2Parser = ParserFunc(func(html []byte) ([]domain.Service, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, err
	}
	tz, _ := time.LoadLocation("Europe/London")
	now := time.Now().In(tz)
	sel := doc.Find(".waste__collections")
	sel = sel.Find(".govuk-grid-column-two-thirds")
	services := []domain.Service{}
	sel.Find(".waste-service-name").Each(func(_ int, sel *goquery.Selection) {
		services = append(services, domain.Service{
			Name: sel.Text(),
		})
	})
	sel.Find(".govuk-grid-row").Each(func(i int, sel *goquery.Selection) {
		sel.Find(".govuk-summary-list__row").Each(func(_ int, sel *goquery.Selection) {
			value := strings.Split(strings.TrimSpace(sel.Find(".govuk-summary-list__value").Text()), "\n")[0]
			switch sel.Find(".govuk-summary-list__key").Text() {
			case "Frequency":
				services[i].Schedule = value
			case "Next collection":
				value = dropOrdinals(value)
				t, _ := time.ParseInLocation("Monday, 2 January", value, tz)
				t = t.AddDate(now.Year(), 0, 0)
				if t.YearDay() < now.YearDay() {
					t = t.AddDate(now.Year()+1, 0, 0)
				}
				services[i].NextService = t
			case "Last collection":
				services[i].Status = value
				value = dropOrdinals(value)
				t, _ := time.ParseInLocation("Monday, 2 January, at 3:04pm", value, tz)
				t = t.AddDate(now.Year(), 0, 0)
				if t.YearDay() > now.YearDay() {
					t = t.AddDate(now.Year()-1, 0, 0)
				}
				services[i].LastService = t
			}
		})
	})
	return services, nil
})

func dropOrdinals(s string) string {
	for _, ordinal := range []string{"st", "nd", "rd", "th"} {
		s = strings.ReplaceAll(s, ordinal, "")
	}
	return s
}

// V1Parser
var V1Parser = ParserFunc(func(html []byte) ([]domain.Service, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, err
	}

	sel := doc.Find(".results-table-wrapper")
	services := []domain.Service{}
	sel.Find(".service-wrapper").Each(func(_ int, sel *goquery.Selection) {
		service := domain.Service{}
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
