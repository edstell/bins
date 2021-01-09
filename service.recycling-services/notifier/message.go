package notifier

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
)

type Message interface {
	Format() (string, error)
}

type MessageFunc func() (string, error)

func (f MessageFunc) Format() (string, error) {
	return f()
}

func ServicesTomorrow(timeNow func() time.Time) func(recyclingservices.Property) Message {
	t, err := template.New("ServicesTomorrow").Funcs(map[string]interface{}{
		"timeNow":    timeNow,
		"formatDate": formatDate,
		"binList":    binList,
	}).Parse(`Hey! You've got a collection tomorrow ({{timeNow|formatDate}}); don't forget to take your {{.Services|binList}} out.`)
	if err != nil {
		panic(err)
	}
	return func(property recyclingservices.Property) Message {
		return MessageFunc(func() (string, error) {
			var out bytes.Buffer
			if err := t.Execute(&out, property); err != nil {
				return "", err
			}
			return out.String(), nil
		})
	}
}

func ServicesThisWeek(timeNow func() time.Time) func(recyclingservices.Property) Message {
	return func(property recyclingservices.Property) Message {
		return MessageFunc(func() (string, error) {
			return "", nil
		})
	}
}

func DescribeProperty() func(recyclingservices.Property) Message {
	return func(property recyclingservices.Property) Message {
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

func binList(services []recyclingservices.Service) string {
	if len(services) == 1 {
		return services[0].Name + " bin"
	}
	names := make([]string, 0, len(services))
	for _, service := range services {
		names = append(names, fmt.Sprintf("'%s'", strings.ToLower(service.Name)))
	}
	list := strings.Join(names[:len(names)-1], ", ")
	return list + " and " + names[len(names)-1] + " bins"
}
