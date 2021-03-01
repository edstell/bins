package message

import (
	"bytes"
	"text/template"
	"time"

	"github.com/edstell/bins/service.recycling-services/domain"
)

func ServicesTomorrow(timeNow func() time.Time) func(domain.Property) (Message, error) {
	t, err := template.New("ServicesTomorrow").Funcs(map[string]interface{}{
		"tomorrow": func() string {
			return formatDate(timeNow().Add(time.Hour * 24))
		},
		"plural":       plural,
		"listServices": listServices,
	}).Parse(`Hey! You've got a collection tomorrow ({{tomorrow}}); don't forget to take your {{.Services|listServices}} bin{{len .Services|plural}} out.`)
	if err != nil {
		panic(err)
	}
	return func(property domain.Property) (Message, error) {
		services := Services(property.Services).Filter(nextCollectionOnDay(timeNow().Add(24 * time.Hour)))
		if len(services) == 0 {
			return DontSend(), nil
		}
		var out bytes.Buffer
		if err := t.Execute(&out, struct {
			Services Services
		}{
			Services: services,
		}); err != nil {
			return nil, err
		}
		return &BodyOnly{out.String()}, nil
	}
}

func nextCollectionOnDay(date time.Time) func(domain.Service) bool {
	return func(service domain.Service) bool {
		return date.Year() == service.NextService.Year() &&
			date.Day() == service.NextService.Day()
	}
}
