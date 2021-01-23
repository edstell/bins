package message

import (
	"testing"
	"time"

	"github.com/edstell/lambda/service.recycling-services/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRawDate(t *testing.T) {
	t.Parallel()
	assert.Equal(t, formatDate(time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)), "Sat 9th")
}

func TestToList(t *testing.T) {
	t.Parallel()
	assert.Equal(t, binList([]domain.Service{
		{
			Name: "General waste",
		},
		{
			Name: "Plastic and tins",
		},
		{
			Name: "Cardboard",
		},
	}), "'general waste', 'plastic and tins' and 'cardboard' bins")
}

func TestServicesTomorrow(t *testing.T) {
	t.Parallel()
	friEigth := time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC)
	message, err := ServicesTomorrow(func() time.Time { return friEigth })(domain.Property{
		Services: []domain.Service{
			{
				Name: "General waste",
			},
			{
				Name: "Plastic and tins",
			},
			{
				Name: "Cardboard",
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "Hey! You've got a collection tomorrow (Sat 9th); don't forget to take your 'general waste', 'plastic and tins' and 'cardboard' bins out.", message.(*BodyOnly).Body)
}

func TestServicesNextWeek(t *testing.T) {
	t.Parallel()
	satNinth := time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)
	message, err := ServicesNextWeek(func() time.Time { return satNinth })(domain.Property{
		Services: []domain.Service{
			{
				Name:        "General waste",
				NextService: time.Date(satNinth.Year(), satNinth.Month(), satNinth.Day()+2, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:        "Plastic and tins",
				NextService: time.Date(satNinth.Year(), satNinth.Month(), satNinth.Day()+2, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:        "Cardboard",
				NextService: time.Date(satNinth.Year(), satNinth.Month(), satNinth.Day()+3, 0, 0, 0, 0, time.UTC),
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "Hey! You have 2 collection[s] next week (w/c Sun 10th): 'general waste' and 'plastic and tins' bins on Monday and 'cardboard' bin on Tuesday.", message.(*BodyOnly).Body)
}
