package notifier

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
)

type Services []*recyclingservicesproto.Service

func (ss Services) Filter(pred func(*recyclingservicesproto.Service) bool) Services {
	filtered := make([]*recyclingservicesproto.Service, 0, len(ss))
	for _, s := range ss {
		if pred(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

type Message interface {
	Format() (string, error)
}

type MessageFunc func() (string, error)

func (f MessageFunc) Format() (string, error) {
	return f()
}

func ServicesTomorrow(timeNow func() time.Time) func(*recyclingservicesproto.Property) Message {
	t, err := template.New("ServicesTomorrow").Funcs(map[string]interface{}{
		"tomorrow": func() string {
			return formatDate(timeNow().Add(time.Hour * 24))
		},
		"binList": binList,
	}).Parse(`Hey! You've got a collection tomorrow ({{tomorrow}}); don't forget to take your {{.Services|binList}} out.`)
	if err != nil {
		panic(err)
	}
	return func(property *recyclingservicesproto.Property) Message {
		return MessageFunc(func() (string, error) {
			var out bytes.Buffer
			if err := t.Execute(&out, property); err != nil {
				return "", err
			}
			return out.String(), nil
		})
	}
}

func ServicesNextWeek(timeNow func() time.Time) func(*recyclingservicesproto.Property) Message {
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
	inRange := func(service *recyclingservicesproto.Service) bool {
		return service.NextService.AsTime().After(start.Add(-1)) && service.NextService.AsTime().Before(end.Add(1))
	}
	type input struct {
		Collections map[time.Weekday][]*recyclingservicesproto.Service
	}
	return func(property *recyclingservicesproto.Property) Message {
		collections := map[time.Weekday][]*recyclingservicesproto.Service{}
		for _, service := range Services(property.Services).Filter(inRange) {
			services, ok := collections[service.NextService.AsTime().Weekday()]
			if !ok {
				services = []*recyclingservicesproto.Service{}
			}
			collections[service.NextService.AsTime().Weekday()] = append(services, service)
		}
		var out bytes.Buffer
		err := t.Execute(&out, input{
			Collections: collections,
		})
		return MessageFunc(func() (string, error) {
			if err != nil {
				return "", err
			}
			return out.String(), nil
		})
	}
}

func DescribeProperty() func(*recyclingservicesproto.Property) Message {
	return func(property *recyclingservicesproto.Property) Message {
		return MessageFunc(func() (string, error) {
			return "", nil
		})
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

func binList(services []*recyclingservicesproto.Service) string {
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

func listCollections(collections map[time.Weekday][]*recyclingservicesproto.Service) string {
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
