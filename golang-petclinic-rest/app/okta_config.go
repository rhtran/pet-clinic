package app

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type OktaConfig struct {
	OAuth2 OAuth2Config
}

type OAuth2Config struct {
	URL      string
	Audience string
	Issuer   string
	ClientId string
	Scopes   string
}

func (oa OAuth2Config) Validate() error {
	return validation.ValidateStruct(&oa,
		validation.Field(&oa.URL, validation.Required),
		validation.Field(&oa.Issuer, validation.Required),
	)
}
