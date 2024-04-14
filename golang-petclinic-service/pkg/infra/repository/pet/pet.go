package pet

import (
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"time"
)

/*
Belongs To
A belongs to association sets up a one-to-one connection with another model,
Pet

Notes:
	can not add this: Owner     owner.Owner `gorm:"foreignKey:OwnerID"` because
	Owner struct has Pets attribute relationship.  Therefore, it creates
	circular relationship and cycle relationship not allowed
*/
// Pet type belongs to Type, TypeID is the foreign key
type Pet struct {
	model.Base
	Name      string     `json:"name"`
	BirthDate *time.Time `json:"birthDate"`
	TypeID    int        `json:"typeId"`
	OwnerID   int        `json:"ownerId"`
	Type      Type       `gorm:"foreignKey:TypeID"`
}

type Type struct {
	model.Base
	Name string `json:"name"`
}
