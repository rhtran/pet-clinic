package repository

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"gorm.io/gorm"
)

type PetRepositorier interface {
	FindById(id int) (*Pet, error)
	FindByName(name string) ([]Pet, error)
	Insert(pet *Pet) (*Pet, error)
	Update(pet *Pet) (*Pet, error)
}

// PetRepository searches pet from the database
type PetRepository struct {
	logger log.Logger
	pg     *gorm.DB
}

func NewPetRepository(logger log.Logger, pg *gorm.DB) *PetRepository {
	return &PetRepository{
		logger: logger,
		pg:     pg,
	}
}

/*
Join query reduces into 1 query by joining 2 tables together through LEFT JOIN.

SELECT "pets"."id","pets"."created_at","pets"."updated_at","pets"."deleted_at","pets"."name","pets"."birth_date",

	"pets"."type_id","pets"."owner_id","Type"."id" AS "Type__id","Type"."created_at" AS "Type__created_at",
	"Type"."updated_at" AS "Type__updated_at","Type"."deleted_at" AS "Type__deleted_at","Type"."name" AS "Type__name"
	FROM "pets" LEFT JOIN "types" "Type" ON "pets"."type_id" = "Type"."id"
	WHERE pets.id = 2 AND "pets"."deleted_at" IS NULL ORDER BY "pets"."id" LIMIT 1
*/
func (repository *PetRepository) FindById(id int) (*Pet, error) {
	repository.logger.Infof("Search pet by id: %v", id)

	var pet Pet
	result := repository.pg.Joins("Type").Where("pets.id = ?", id).First(&pet)
	return &pet, result.Error
}

/*
SELECT * FROM "types" WHERE "types"."id" = 1 AND "types"."deleted_at" IS NULL
SELECT * FROM "pets" WHERE name = 'Leo' AND "pets"."deleted_at" IS NULL
*/
func (repository *PetRepository) FindByName(name string) ([]Pet, error) {
	repository.logger.Infof("Search pet by name: %v", name)

	var pets []Pet
	result := repository.pg.Preload("Type").Where("name = ?", name).Find(&pets)
	return pets, result.Error
}

func (repository *PetRepository) Insert(pet *Pet) (*Pet, error) {
	repository.logger.Infof("insert a new pet: %v", pet.Name)

	err := repository.pg.Create(&pet).Error
	if err != nil {
		repository.logger.Error("fails to insert new pet, err: %v", err.Error)
		return nil, err
	}
	return pet, err
}

func (repository *PetRepository) Update(pet *Pet) (*Pet, error) {
	repository.logger.Infof("update vet id: %v", pet.ID)

	// Omit the column name from update...
	err := repository.pg.Omit("created_at").Save(&pet).Error
	if err != nil {
		repository.logger.Error("fails to update pet, err: %v", err.Error)
		return nil, err
	}

	return pet, err
}
