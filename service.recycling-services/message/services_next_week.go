package message

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/edstell/lambda/service.recycling-services/domain"
)

func ServicesNextWeek(timeNow func() time.Time) func(domain.Property) (Message, error) {
	now := timeNow()
	start := time.Date(now.Year(), now.Month(), now.Day()+int(7-now.Weekday()), 0, 0, 0, 0, time.UTC)
	end := start.Add(7 * 24 * time.Hour)
	t, err := template.New("ServicesNextWeek").Funcs(
		map[string]interface{}{
			"weekstart": func() string {
				return formatDate(start)
			},
			"listCollections": listCollections,
		},
	).Parse(`Hey! You have {{len .Collections}} collection[s] next week (w/c {{weekstart}}){{if .Collections}}: {{.Collections|listCollections}}{{end}}`)
	if err != nil {
		panic(err)
	}
	return func(property domain.Property) (Message, error) {
		collections := map[time.Weekday][]domain.Service{}
		for _, service := range Services(property.Services).Filter(nextCollectionInRange(start, end)) {
			services, ok := collections[service.NextService.Weekday()]
			if !ok {
				services = []domain.Service{}
			}
			collections[service.NextService.Weekday()] = append(services, service)
		}
		var out bytes.Buffer
		err := t.Execute(&out, struct {
			Collections map[time.Weekday][]domain.Service
		}{
			Collections: collections,
		})
		if err != nil {
			return nil, err
		}
		return &BodyOnly{out.String()}, nil
	}
}

func DescribeProperty() func(domain.Property) (Message, error) {
	return func(property domain.Property) (Message, error) {
		return nil, nil
	}
}

func formatDate(t time.Time) string {
	suffix := "th"
	switch t.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}
	return t.Format("Mon 2") + suffix
}

func binList(services []domain.Service) string {
	if len(services) == 1 {
		return fmt.Sprintf("'%s'", strings.ToLower(services[0].Name)) + " bin"
	}
	names := make([]string, 0, len(services))
	for _, service := range services {
		names = append(names, fmt.Sprintf("'%s'", strings.ToLower(service.Name)))
	}
	list := strings.Join(names[:len(names)-1], ", ")
	return list + " and " + names[len(names)-1] + " bins"
}

func listCollections(collections map[time.Weekday][]domain.Service) string {
	list := ""
	i := 0
	for weekday, services := range collections {
		sep := ", "
		if i == 0 {
			sep = ""
		} else if i == len(collections)-1 {
			sep = " and "
		}
		list = list + sep + binList(services) + " on " + fmt.Sprint(weekday)
		i++
	}
	return list + "."
}
