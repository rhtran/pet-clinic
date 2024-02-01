package app

import validation "github.com/go-ozzo/ozzo-validation"

type ServerConfig struct {
	HttpPort string
}

func (sc ServerConfig) Validate() error {
	return validation.ValidateStruct(&sc,
		validation.Field(&sc.HttpPort, validation.Required),
	)
}
