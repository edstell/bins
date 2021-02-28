package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
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
	<div class="service-wrapper">
		{{range .Services}}
		<div class="service-name">
			<h3>{{.Name}}</h3>
			<div class="service-content">
				<div class="image-wrapper"></div>
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
								<p>{{.Status}}</p>
							</td>
						</tr>
					</tbody>
				</table>
			</div>
		</div>
		{{end}}
	</div>
</div>
`

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
