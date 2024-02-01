package app

import validation "github.com/go-ozzo/ozzo-validation"

type AppInfoConfig struct {
	Name        string
	Description string
	Version     string
}

func (aic AppInfoConfig) Validate() error {
	return validation.ValidateStruct(&aic,
		validation.Field(&aic.Name, validation.Required),
		validation.Field(&aic.Description, validation.Required),
		validation.Field(&aic.Version, validation.Required),
	)
}
