package notifier

import (
	"testing"
	"time"

	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatDate(t *testing.T) {
	t.Parallel()
	assert.Equal(t, formatDate(time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)), "Sat 9th")
}

func TestToList(t *testing.T) {
	t.Parallel()
	assert.Equal(t, binList([]recyclingservices.Service{
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
	satNinth := time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)
	result, err := ServicesTomorrow(func() time.Time { return satNinth })(recyclingservices.Property{
		Services: []recyclingservices.Service{
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
	}).Format()
	require.NoError(t, err)
	assert.Equal(t, "Hey! You've got a collection tomorrow (Sat 9th); don't forget to take your 'general waste', 'plastic and tins' and 'cardboard' bins out.", result)
}
