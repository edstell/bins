package services

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/edstell/bins/service.recycling-services/domain"
	"github.com/edstell/bins/service.recycling-services/services/assets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebScraper(t *testing.T) {
	t.Parallel()
	html, err := assets.Asset("services/assets/v1.html")
	require.NoError(t, err)
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write(html)
	}))
	defer server.Close()
	scraper := WebScraper(server.Client(), V1Parser, server.URL)
	result, err := scraper.Fetch(context.Background(), "property_id")
	require.NoError(t, err)
	fmt.Println(result)
	for i, service := range []domain.Service{
		{
			Name:        "Non-recyclable refuse",
			Status:      "Not completed.",
			Schedule:    "Thursday every other week",
			LastService: time.Date(2020, 12, 17, 0, 0, 0, 0, time.UTC),
			NextService: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:        "Paper and cardboard",
			Status:      "Not completed.",
			Schedule:    "Thursday every other week",
			LastService: time.Date(2020, 12, 17, 0, 0, 0, 0, time.UTC),
			NextService: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:        "Green Garden Waste (Subscription)",
			Status:      "Your road was completed on 09/12/2020 at 09:27.",
			Schedule:    "Wednesday every 4th week",
			LastService: time.Date(2020, 12, 9, 0, 0, 0, 0, time.UTC),
			NextService: time.Date(2021, 1, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:        "Food waste",
			Status:      "Not completed.",
			Schedule:    "Thursday every week",
			LastService: time.Date(2020, 12, 24, 0, 0, 0, 0, time.UTC),
			NextService: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:        "Plastic, glass and tins",
			Status:      "Your road was completed on 24/12/2020 at 08:49.",
			Schedule:    "Thursday every other week",
			LastService: time.Date(2020, 12, 24, 0, 0, 0, 0, time.UTC),
			NextService: time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			Name: "Batteries, small electrical items and textiles",
		},
	} {
		assert.Equal(t, service, result[i], service.Name)
	}
}
