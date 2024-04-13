package app

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"regexp"
)

type ServerConfig struct {
	HttpPort string `yaml:"httpPort"`
}

func (sc ServerConfig) Validate() error {
	return validation.ValidateStruct(&sc,
		validation.Field(&sc.HttpPort, validation.Required),
		validation.Field(&sc.HttpPort, validation.Match(regexp.MustCompile("^:\\d+$")).Error("must be in the format :dddd")),
	)
}
