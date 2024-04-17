package repository

import (
	"gorm.io/gorm"
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
	gorm.Model
	Name      string     `gorm:"name"`
	BirthDate *time.Time `gorm:"birth_date"`
	TypeID    int
	OwnerID   int
	Type      Type `gorm:"foreignKey:TypeID"`
}

type Type struct {
	gorm.Model
	Name string
}
