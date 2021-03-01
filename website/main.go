package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	recyclingservicesproto "github.com/edstell/bins/service.recycling-services/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"honnef.co/go/js/dom"
)

func main() {
	if err := displayProperty(); err != nil {
		println(fmt.Errorf("failed to display property: %v", err))
	}
}

const templ = `
<div class="services-wrapper">
	{{range .Services}}
	<div class="service-wrapper">
		<h3 class="service-name">{{.Name}}</h3>
		<div class="service-content">
			<div class="image-wrapper">
				<div class="image" style="background-image: url(images/{{.Name|imageURL}});"></div>
			</div>
			<table>
				<thead>
					<tr>
						<th>Schedule</th>
						<th>Last Service</th>
						<th>Next Service</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>{{.Schedule}}</td>
						<td>{{.LastService|formatDate}}</td>
						<td>{{.NextService|formatDate}}</td>
					</tr>
					<tr>
						<td colspan="3">
							<div class="task-state">
								<p>{{.Status}}</p>
							</div>
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	</div>
	</br>
	{{end}}
</div>
`

func imageURL(name string) string {
	switch name {
	case "Non-recyclable refuse":
		return "refuse.png"
	case "Paper and cardboard":
		return "paper.png"
	case "Green Garden Waste (Subscription)":
		return "garden.png"
	case "Food waste":
		return "food-waste.png"
	case "Plastic, glass and tins":
		return "recycling.png"
	case "Batteries, small electrical items and textiles":
		return "battery_bag.png"
	default:
		return ""
	}
}

func formatDate(t *timestamppb.Timestamp) string {
	if t.AsTime().IsZero() {
		return ""
	}
	return t.AsTime().Format("2 Jan 2006")
}

func displayProperty() error {
	property, err := getProperty()
	if err != nil {
		return err
	}
	t := template.New("property").Funcs(map[string]interface{}{
		"formatDate": formatDate,
		"imageURL":   imageURL,
	})
	t, err = t.Parse(templ)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, property); err != nil {
		return err
	}
	d := dom.GetWindow().Document()
	body := d.GetElementsByTagName("body")[0].(*dom.HTMLBodyElement)
	if body == nil {
		return fmt.Errorf("failed to get body")
	}
	body.SetInnerHTML(buf.String())
	return nil
}

func getProperty() (*recyclingservicesproto.Property, error) {
	rsp, err := http.Get("https://api.edstell.com/properties/100020406685")
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	property := &recyclingservicesproto.Property{}
	if err := json.Unmarshal(body, property); err != nil {
		return nil, err
	}
	return property, nil
}
