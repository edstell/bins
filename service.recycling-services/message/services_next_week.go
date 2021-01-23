package message

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/edstell/lambda/service.recycling-services/domain"
)

type collection struct {
	Services []domain.Service
	Weekday  time.Weekday
}

// ServicesNextWeek returns a function which when called will construct a
// message with details of bin collections in the coming week.
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
			"plural":          plural,
			"word":            word,
		},
	).Parse(`Hey! You have {{len .Collections|word}} collection{{len .Collections|plural}} next week (w/c {{weekstart}}){{if .Collections}}: {{.Collections|listCollections}}{{end}}`)
	if err != nil {
		panic(err)
	}
	return func(property domain.Property) (Message, error) {
		c := map[time.Weekday][]domain.Service{}
		for _, service := range Services(property.Services).Filter(nextCollectionInRange(start, end)) {
			services, ok := c[service.NextService.Weekday()]
			if !ok {
				services = []domain.Service{}
			}
			c[service.NextService.Weekday()] = append(services, service)
		}
		cs := make([]collection, 0, len(c))
		for weekday, services := range c {
			cs = append(cs, collection{
				Weekday:  weekday,
				Services: services,
			})
		}
		sort.Slice(cs, func(i, j int) bool {
			return cs[i].Weekday < cs[j].Weekday
		})
		var out bytes.Buffer
		err := t.Execute(&out, struct {
			Collections []collection
		}{
			Collections: cs,
		})
		if err != nil {
			return nil, err
		}
		return &BodyOnly{out.String()}, nil
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

func listServices(services []domain.Service) string {
	if len(services) == 1 {
		return fmt.Sprintf("'%s'", strings.ToLower(services[0].Name))
	}
	names := make([]string, 0, len(services))
	for _, service := range services {
		names = append(names, fmt.Sprintf("'%s'", strings.ToLower(service.Name)))
	}
	list := strings.Join(names[:len(names)-1], ", ")
	return list + " and " + names[len(names)-1]
}

func listCollections(collections []collection) string {
	l := ""
	i := 0
	for _, collection := range collections {
		sep := ", "
		if i == 0 {
			sep = ""
		} else if i == len(collections)-1 {
			sep = " and "
		}
		l = l + sep + listServices(collection.Services) + " on " + fmt.Sprint(collection.Weekday)
		i++
	}
	return l + "."
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

func word(n int) string {
	if n > 10 {
		return fmt.Sprint(n)
	}
	return []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}[n]
}
