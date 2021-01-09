package recyclingservices

import "github.com/edstell/lambda/libraries/errors"

func (r ReadPropertyRequest) Validate() error {
	if r.PropertyID == "" {
		return errors.MissingParam("property_id")
	}
	return nil
}

func (r WritePropertyRequest) Validate() error {
	if r.PropertyID == "" {
		return errors.MissingParam("property_id")
	}
	if len(r.Services) == 0 {
		return errors.MissingParam("services")
	}
	for _, service := range r.Services {
		if err := service.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (r SyncPropertyRequest) Validate() error {
	if r.PropertyID == "" {
		return errors.MissingParam("property_id")
	}
	return nil
}

func (s Service) Validate() error {
	if s.Name == "" {
		return errors.MissingParam("name")
	}
	if s.Status == "" {
		return errors.MissingParam("status")
	}
	if s.Schedule == "" {
		return errors.MissingParam("schedule")
	}
	if s.LastService.IsZero() {
		return errors.MissingParam("last_service")
	}
	if s.NextService.IsZero() {
		return errors.MissingParam("next_service")
	}
	return nil
}
