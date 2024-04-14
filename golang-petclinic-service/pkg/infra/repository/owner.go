package repository

import (
	"github.com/rhtran/golang-petclinic-service/pkg/model"
)

/*
Owner has many Pets, OwnerID is the foreign key
Has Many configuration
*/
type Owner struct {
	model.Base
	model.Person
	Username  string `json:"username"`
	Password  string `json:"-"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
	Pets      []Pet  `json:"pets"`
}
