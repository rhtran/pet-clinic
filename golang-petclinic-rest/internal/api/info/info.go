package info

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type Info struct {
	AppName     string `json:"appName"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Ip          string `json:"ip"`
	GitCommit   string `json:"gitCommit"`
}

func (i Info) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.AppName, validation.Required),
		validation.Field(&i.Description, validation.Required),
		validation.Field(&i.Version, validation.Required),
	)
}
