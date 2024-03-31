package health

import validation "github.com/go-ozzo/ozzo-validation/v3"

type Check struct {
	Status string `json:"status"`
}

func (hc Check) Validate() error {
	return validation.ValidateStruct(&hc,
		validation.Field(&hc.Status, validation.Required),
	)
}
