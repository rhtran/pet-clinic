package health

import validation "github.com/go-ozzo/ozzo-validation/v3"

type HealthCheck struct {
	Status string `json:"status"`
}

func (hc HealthCheck) Validate() error {
	return validation.ValidateStruct(&hc,
		validation.Field(&hc.Status, validation.Required),
	)
}
