package vet

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"gorm.io/gorm"
)

// VetRepository searches vet from the database
type VetRepository struct {
	logger log.Logger
	pg     *gorm.DB
}

func NewVetRepository(logger log.Logger, pg *gorm.DB) *VetRepository {
	return &VetRepository{
		logger: logger,
		pg:     pg,
	}
}

/*
Many-to-Many relation vets <-- vet_specialties --> specialties
Preload - eagerly loading the specialties relation and hence,
3 queries are executed.
SELECT * FROM "vet_specialties" WHERE "vet_specialties"."vet_id" = 2
SELECT * FROM "specialties" WHERE "specialties"."id" = 1 AND "specialties"."deleted_at" IS NULL
SELECT * FROM "vets" WHERE "vets"."id" = 2 AND "vets"."deleted_at" IS NULL ORDER BY "vets"."id" LIMIT 1
*/
func (repository *VetRepository) FindById(id int) (*Vet, error) {
	repository.logger.Infof("Search vet by id: %v", id)

	var vet Vet
	err := repository.pg.Preload("Specialties").First(&vet, id).Error
	if err != nil {
		repository.logger.Errorf("fails to find vet by id: %v, err: %v", id, err.Error)
		return nil, err
	}

	return &vet, err
}

/*
SELECT * FROM "vet_specialties" WHERE "vet_specialties"."vet_id" = 5
SELECT * FROM "specialties" WHERE "specialties"."id" = 1 AND "specialties"."deleted_at" IS NULL
SELECT * FROM "vets" WHERE last_name = 'Stevens' AND "vets"."deleted_at" IS NULL
*/
func (repository *VetRepository) FindByLastName(lastName string) ([]Vet, error) {
	repository.logger.Infof("Search vet by last name: %v", lastName)

	var vets []Vet
	err := repository.pg.Preload("Specialties").Where("last_name = ?", lastName).Find(&vets).Error
	if err != nil {
		repository.logger.Errorf("fails to find vet by last name: %v, err: %v", lastName, err.Error)
		return nil, err
	}

	return vets, nil
}

/*
SELECT * FROM "vets" WHERE "vets"."deleted_at" IS NULL
*/
func (repository *VetRepository) FindAll() ([]Vet, error) {
	repository.logger.Info("get all vets")

	var vets []Vet
	err := repository.pg.Find(&vets).Error
	if err != nil {
		repository.logger.Errorf("fails to find all vets, err: %v", err.Error)
		return nil, err
	}
	return vets, err
}

/*
Expect there would be n + 1 queries executed just like in Java Hibernate.  Surprisedly,
only 3 queries actually executed.

	SELECT * FROM "vet_specialties" WHERE "vet_specialties"."vet_id" IN (1,2,3,4,5,6)
	SELECT * FROM "specialties" WHERE "specialties"."id" IN (1,2,3) AND "specialties"."deleted_at" IS NULL
	SELECT * FROM "vets" WHERE "vets"."deleted_at" IS NULL
*/
func (repository *VetRepository) FindAllPreload() ([]Vet, error) {
	repository.logger.Info("get all vets with relations preloaded")

	var vets []Vet
	err := repository.pg.Preload("Specialties").Find(&vets).Error
	if err != nil {
		repository.logger.Errorf("fails to find all vets, err: %v", err.Error)
		return nil, err
	}
	return vets, err
}

func (repository *VetRepository) Insert(vet *Vet) (*Vet, error) {
	repository.logger.Infof("insert a new vet: %v", vet)

	err := repository.pg.Create(&vet).Error
	if err != nil {
		repository.logger.Errorf("fails to insert new vet, err: %v", err.Error)
		return nil, err
	}
	return vet, err
}

func (repository *VetRepository) Update(vet *Vet) (*Vet, error) {
	repository.logger.Infof("update vet id: %v", vet.ID)

	// Omit the column name from update...
	err := repository.pg.Omit("created_at").Save(&vet).Error
	if err != nil {
		repository.logger.Errorf("fails to update vet, err: %v", err.Error)
		return nil, err
	}

	return vet, err
}
