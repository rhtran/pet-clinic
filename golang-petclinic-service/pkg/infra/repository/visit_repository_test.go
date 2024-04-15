package repository

import (
	"github.com/rhtran/golang-petclinic-service/pkg/infra/repository/test"
	"github.com/rhtran/golang-petclinic-service/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qiangxue/go-restful-api/pkg/log"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/suite"
)

type VisitRepoTestSuite struct {
	suite.Suite
	postgresql      *embeddedpostgres.EmbeddedPostgres
	visitRepository *VisitRepository
}

// This will run before the tests in the suite are run
func (suite *VisitRepoTestSuite) SetupSuite() {
	suite.postgresql = test.PgStart(suite.T(), "test/migrations")
	suite.visitRepository = getVisitRepository(suite.T())
}

func (suite *VisitRepoTestSuite) TearDownSuite() {
	err := suite.postgresql.Stop()
	if err != nil {
		suite.T().Fatal(err)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestVisitRepoTestSuite(t *testing.T) {
	suite.Run(t, new(VisitRepoTestSuite))
}

func getVisitRepository(t *testing.T) *VisitRepository {
	db, err := test.Connect()
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New().With(nil, "version", "visit_repository_test")
	return NewVisitRepository(logger, db)
}

func (suite *VisitRepoTestSuite) Test_FindById() {
	var testCases = []struct {
		input    int
		expected Visit
	}{
		{input: 2, expected: Visit{
			Base: model.Base{
				ID: 2,
			},
			PetID:       8,
			Description: "rabies shot",
		}},
		{input: 1, expected: Visit{
			Base: model.Base{
				ID: 1,
			},
			PetID:       7,
			Description: "rabies shot",
		}},
	}

	for _, testCase := range testCases {
		visit, _ := suite.visitRepository.FindById(testCase.input)
		assert.Equal(suite.T(), testCase.expected.ID, visit.ID)
		assert.Equal(suite.T(), testCase.expected.PetID, visit.PetID)
		assert.Equal(suite.T(), testCase.expected.Description, visit.Description)
		//assert.Equal(suite.T(), testCase.expected.VisitDate, visit.VisitDate)
	}
}
