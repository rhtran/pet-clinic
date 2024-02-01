package vet

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wire-petclinic-service/internal/api/test"
	"wire-petclinic-service/pkg/model"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/suite"

	"github.com/qiangxue/go-restful-api/pkg/log"
)

type VetRepoTestSuite struct {
	suite.Suite
	postgresql    *embeddedpostgres.EmbeddedPostgres
	vetRepository *VetRepository
}

// This will run before the tests in the suite are run
func (suite *VetRepoTestSuite) SetupSuite() {
	suite.postgresql = test.PgStart(suite.T(), "../test/migrations")
	suite.vetRepository = getVetRepository(suite.T())
}

func (suite *VetRepoTestSuite) TearDownSuite() {
	err := suite.postgresql.Stop()
	if err != nil {
		suite.T().Fatal(err)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPetRepoTestSuite(t *testing.T) {
	suite.Run(t, new(VetRepoTestSuite))
}

func getVetRepository(t *testing.T) *VetRepository {
	db, err := test.Connect()
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New().With(nil, "version", "vet_repository_test")
	return NewVetRepository(logger, db)
}

func (suite *VetRepoTestSuite) Test_FindById() {
	var testCases = []struct {
		input    int
		expected Vet
	}{
		{1, Vet{
			Base: model.Base{
				ID: 1,
			},
			Person: model.Person{
				FirstName: "James",
				LastName:  "Carter",
			},
		}},
		{2, Vet{
			Base: model.Base{
				ID: 2,
			},
			Person: model.Person{
				FirstName: "Helen",
				LastName:  "Leary",
			},
		}},
	}

	for _, testCase := range testCases {
		vet, _ := suite.vetRepository.FindById(testCase.input)
		assert.Equal(suite.T(), testCase.expected.LastName, vet.LastName)
		assert.Equal(suite.T(), testCase.expected.FirstName, vet.FirstName)
	}
}

func (suite *VetRepoTestSuite) Test_FindByLastName() {
	var testCases = []struct {
		input    string
		expected Vet
	}{
		{"Carter", Vet{
			Person: model.Person{
				FirstName: "James",
				LastName:  "Carter",
			},
		}},
		{"Leary", Vet{
			Person: model.Person{
				FirstName: "Helen",
				LastName:  "Leary",
			},
		}},
	}

	for _, testCase := range testCases {
		vets, _ := suite.vetRepository.FindByLastName(testCase.input)
		assert.Equal(suite.T(), testCase.expected.LastName, vets[0].LastName)
		assert.Equal(suite.T(), testCase.expected.FirstName, vets[0].FirstName)
	}
}
