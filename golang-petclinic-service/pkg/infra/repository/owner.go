package repository

import (
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"gorm.io/gorm"
)

/*
Owner has many Pets, OwnerID is the foreign key
Has Many configuration
*/
type Owner struct {
	gorm.Model
	model.Person
	Username  string
	Password  string
	Address   string
	City      string
	Telephone string
	Pets      []Pet
}
