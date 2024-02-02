package visit

import (
	"github.com/qiangxue/go-restful-api/pkg/log"

	"gorm.io/gorm"
)

// VetRepository searches vet from the database
type VisitRepository struct {
	logger log.Logger
	pg     *gorm.DB
}

func NewVisitRepository(logger log.Logger, pg *gorm.DB) *VisitRepository {
	return &VisitRepository{
		logger: logger,
		pg:     pg,
	}
}

/*
SELECT * FROM "pets" WHERE "pets"."id" = 8 AND "pets"."deleted_at" IS NULL
SELECT * FROM "types" WHERE "types"."id" = 1 AND "types"."deleted_at" IS NULL
SELECT * FROM "visits" WHERE "visits"."id" = 2 AND "visits"."deleted_at" IS NULL ORDER BY "visits"."id" LIMIT 1
*/
func (repository *VisitRepository) FindById(id int) (*Visit, error) {
	repository.logger.Infof("Search visit by id: %v", id)

	var visit Visit
	err := repository.pg.Preload("Pet.Type").First(&visit, id).Error
	if err != nil {
		repository.logger.Errorf("fails to find visit by id: %v, err: %v", id, err.Error)
		return nil, err
	}

	return &visit, err
}

/*
Find all visits
SELECT * FROM "pets" WHERE "pets"."id" IN (7,8) AND "pets"."deleted_at" IS NULL
SELECT * FROM "types" WHERE "types"."id" = 1 AND "types"."deleted_at" IS NULL
SELECT * FROM "visits" WHERE "visits"."deleted_at" IS NULL
*/
func (repository *VisitRepository) FindAll() ([]Visit, error) {
	repository.logger.Info("get list of visits")

	var visits []Visit
	// not needed to preload for all visits - will be a performance if get large result.
	err := repository.pg.Preload("Pet.Type").Find(&visits).Error
	if err != nil {
		repository.logger.Errorf("fails to find all visits, err: %v", err.Error)
		return nil, err
	}

	return visits, err
}
