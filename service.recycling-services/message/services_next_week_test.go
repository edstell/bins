package message

import (
	"testing"
	"time"

	"github.com/edstell/lambda/service.recycling-services/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
	assert.Equal(t, "Hey! You have two collections next week (w/c Sun 10th): 'general waste' and 'plastic and tins' on Monday and 'cardboard' on Tuesday.", message.(*BodyOnly).Body)
}
