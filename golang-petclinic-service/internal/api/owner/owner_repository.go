package owner

import (
	"github.com/qiangxue/go-restful-api/pkg/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repositorier interface {
	FindById(id int) (*Owner, error)
	FindByLastName(lastName string) ([]Owner, error)
	FindAll() ([]Owner, error)
	FindAllWithPets() ([]Owner, error)
	Insert(owner *Owner) (*Owner, error)
	Update(owner *Owner) (*Owner, error)
}

// Repository searches owner from the database
type Repository struct {
	logger log.Logger
	pg     *gorm.DB
}

func NewOwnerRepository(logger log.Logger, pg *gorm.DB) *Repository {
	return &Repository{
		logger: logger,
		pg:     pg,
	}
}

// FindById
// find the owner by id
//
//	SELECT * FROM "pets" WHERE "pets"."owner_id" = 1 AND "pets"."deleted_at" IS NULL
//	SELECT * FROM "types" WHERE "types"."id" = 1 AND "types"."deleted_at" IS NULL
//	SELECT * FROM "owners" WHERE "owners"."id" = 1 AND "owners"."deleted_at" IS NULL ORDER BY "owners"."id" LIMIT 1
func (repository *Repository) FindById(id int) (*Owner, error) {
	repository.logger.Infof("search owner by id: %v", id)

	var owner Owner
	// Pets.Type - Nested Preloading (Eager Loading)
	err := repository.pg.Preload("Pets.Type").First(&owner, id).Error
	if err != nil {
		repository.logger.Errorf("fails to find owner by id: %v, err: %v", id, err.Error)
		return nil, err
	}
	return &owner, err
}

// FindByLastName
// find owner by last name
//
// SELECT * FROM "pets" WHERE "pets"."owner_id" = 3 AND "pets"."deleted_at" IS NULL
// SELECT * FROM "types" WHERE "types"."id" = 2 AND "types"."deleted_at" IS NULL
// SELECT * FROM "owners" WHERE last_name = 'Rodriquez' AND "owners"."deleted_at" IS NULL
func (repository *Repository) FindByLastName(lastName string) ([]Owner, error) {
	repository.logger.Infof("Search owner by last name: %v", lastName)

	// Pets.Type - Nested Preloading (Eager Loading)
	var owners []Owner
	err := repository.pg.Preload("Pets.Type").Where("last_name = ?", lastName).Find(&owners).Error
	if err != nil {
		repository.logger.Errorf("fails to find owners by last name: %v, err: %v", lastName, err.Error)
		return nil, err
	}
	return owners, err
}

/*
Find all owners
SELECT * FROM "owners" WHERE "owners"."deleted_at" IS NULL
*/
func (repository *Repository) FindAll() ([]Owner, error) {
	repository.logger.Info("get list of owners")

	var owners []Owner
	// not needed to preload for all owners - will be a performance if get large result.
	err := repository.pg.Find(&owners).Error
	if err != nil {
		repository.logger.Errorf("fails to find all owners, err: %v", err.Error)
		return nil, err
	}

	return owners, err
}

/*
Find all owners with pets and its nested associations.

SELECT * FROM "pets" WHERE "pets"."owner_id" IN (1,2,3,4,5,6,7,8,9,10) AND "pets"."deleted_at" IS NULL
SELECT * FROM "types" WHERE "types"."id" IN (1,6,2,3,4,5) AND "types"."deleted_at" IS NULL
SELECT * FROM "owners" WHERE "owners"."deleted_at" IS NULL
*/
func (repository *Repository) FindAllWithPets() ([]Owner, error) {
	repository.logger.Info("get list of owners with pets")

	var owners []Owner
	/*
	 clause.Associations won’t preload nested associations, but can use it with Nested Preloading together
	*/
	err := repository.pg.Preload("Pets.Type").Preload(clause.Associations).Find(&owners).Error
	if err != nil {
		repository.logger.Errorf("fails to find all owners with pets, err: %v", err.Error)
		return nil, err
	}

	return owners, err
}

func (repository *Repository) Insert(owner *Owner) (*Owner, error) {
	repository.logger.Infof("insert a new owner: %v", owner)

	err := repository.pg.Create(&owner).Error
	if err != nil {
		repository.logger.Errorf("fails to insert new owner, err: %v", err.Error)
		return nil, err
	}
	return owner, err
}

func (repository *Repository) Update(owner *Owner) (*Owner, error) {
	repository.logger.Infof("update owner id: %v", owner.ID)

	// Omit the column name from update...
	err := repository.pg.Omit("created_at").Save(&owner).Error
	if err != nil {
		repository.logger.Errorf("fails to update owner, err: %v", err.Error)
		return nil, err
	}

	return owner, err
}
