package message

import (
	"testing"
	"time"

	"github.com/edstell/lambda/service.recycling-services/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServicesTomorrow(t *testing.T) {
	t.Parallel()
	friEigth := time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC)
	message, err := ServicesTomorrow(func() time.Time { return friEigth })(domain.Property{
		Services: []domain.Service{
			{
				Name:        "General waste",
				NextService: friEigth.Add(time.Hour * 24),
			},
			{
				Name:        "Plastic and tins",
				NextService: friEigth,
			},
			{
				Name:        "Cardboard",
				NextService: friEigth.Add(time.Hour * 24),
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "Hey! You've got a collection tomorrow (Sat 9th); don't forget to take your 'general waste' and 'cardboard' bins out.", message.(*BodyOnly).Body)
}

func TestServicesTomorrowNone(t *testing.T) {
	t.Parallel()
	message, err := ServicesTomorrow(func() time.Time { return time.Now() })(domain.Property{})
	require.NoError(t, err)
	assert.Implements(t, (*NotSendable)(nil), message)
}
