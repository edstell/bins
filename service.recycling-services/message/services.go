package message

import (
	"time"

	"github.com/edstell/lambda/service.recycling-services/domain"
)

type Services []domain.Service

func (ss Services) Filter(pred func(domain.Service) bool) Services {
	filtered := make([]domain.Service, 0, len(ss))
	for _, s := range ss {
		if pred(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func nextCollectionInRange(start, end time.Time) func(domain.Service) bool {
	return func(service domain.Service) bool {
		return service.NextService.After(start.Add(-1)) && service.NextService.Before(end.Add(1))
	}
}
