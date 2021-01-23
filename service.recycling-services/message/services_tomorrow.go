package message

import (
	"bytes"
	"text/template"
	"time"

	"github.com/edstell/lambda/service.recycling-services/domain"
)

func ServicesTomorrow(timeNow func() time.Time) func(domain.Property) (Message, error) {
	t, err := template.New("ServicesTomorrow").Funcs(map[string]interface{}{
		"tomorrow": func() string {
			return formatDate(timeNow().Add(time.Hour * 24))
		},
		"binList": binList,
	}).Parse(`Hey! You've got a collection tomorrow ({{tomorrow}}); don't forget to take your {{.Services|binList}} out.`)
	if err != nil {
		panic(err)
	}
	return func(property domain.Property) (Message, error) {
		var out bytes.Buffer
		if err := t.Execute(&out, property); err != nil {
			return nil, err
		}
		return &BodyOnly{out.String()}, nil
	}
}
