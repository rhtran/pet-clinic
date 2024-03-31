package pet

import (
	"github.com/rhtran/golang-petclinic-service/internal/api/test"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"

	"github.com/stretchr/testify/suite"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"gopkg.in/go-playground/assert.v1"
)

type PetRepoTestSuite struct {
	suite.Suite
	postgresql    *embeddedpostgres.EmbeddedPostgres
	petRepository *Repository
}

// This will run before the tests in the suite are run
func (suite *PetRepoTestSuite) SetupSuite() {
	suite.postgresql = test.PgStart(suite.T(), "../test/migrations")
	suite.petRepository = getPetRepository(suite.T())
}

func (suite *PetRepoTestSuite) TearDownSuite() {
	err := suite.postgresql.Stop()
	if err != nil {
		suite.T().Fatal(err)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPetRepoTestSuite(t *testing.T) {
	suite.Run(t, new(PetRepoTestSuite))
}

func getPetRepository(t *testing.T) *Repository {
	db, err := test.Connect()
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New().With(nil, "version", "pet_repository_test")
	return NewPetRepository(logger, db)
}

func (suite *PetRepoTestSuite) Test_FindById() {

	var testCases = []struct {
		input    int
		expected string
	}{
		{1, "Leo"},
		{2, "Basil"},
	}

	for _, testCase := range testCases {
		pet, _ := suite.petRepository.FindById(testCase.input)
		assert.Equal(suite.T(), testCase.expected, pet.Name)
	}
}

func (suite *PetRepoTestSuite) Test_FindByName() {
	var testCases = []struct {
		input    string
		expected string
	}{
		{"Leo", "Leo"},
		{"Basil", "Basil"},
	}

	for _, testCase := range testCases {
		pets, _ := suite.petRepository.FindByName(testCase.input)
		assert.Equal(suite.T(), testCase.expected, &pets[0].Name)
	}
}
