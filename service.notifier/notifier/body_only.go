package notifier

import "encoding/json"

type BodyOnly struct {
	body string
}

func (b *BodyOnly) Raw() []byte {
	bytes, err := json.Marshal(b)
	if err != nil {
		// Marshaling BodyOnly will only ever error if 'b' is nil, in which case
		// there's been a mistake in the logic - hence a panic rather than
		// proporgating the error.
		panic(err)
	}
	return bytes
}

func (b *BodyOnly) Body() string {
	return b.body
}
