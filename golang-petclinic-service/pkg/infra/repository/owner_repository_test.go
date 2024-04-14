package repository

import (
	"github.com/rhtran/golang-petclinic-service/internal/api/test"
	model2 "github.com/rhtran/golang-petclinic-service/pkg/model"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/suite"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/assert"
)

type OwnerRepoTestSuite struct {
	suite.Suite
	postgresql      *embeddedpostgres.EmbeddedPostgres
	ownerRepository *OwnerRepository
}

// This will run before the tests in the suite are run
func (suite *OwnerRepoTestSuite) SetupSuite() {
	suite.postgresql = test.PgStart(suite.T(), "../test/migrations")
	suite.ownerRepository = getOwnerRepository(suite.T())
}

func (suite *OwnerRepoTestSuite) TearDownSuite() {
	err := suite.postgresql.Stop()
	if err != nil {
		suite.T().Fatal(err)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOwnerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(OwnerRepoTestSuite))
}

func getOwnerRepository(t *testing.T) *OwnerRepository {
	db, err := test.Connect()
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New().With(nil, "version", "owner_repository_test")
	return NewOwnerRepository(logger, db)
}

func (suite *OwnerRepoTestSuite) Test_FindById() {
	var testCases = []struct {
		input    int
		expected Owner
	}{
		{input: 2, expected: Owner{
			Base: model2.Base{
				ID: 2,
			},
			Person: model2.Person{
				FirstName: "Betty",
				LastName:  "Davis",
			},
		}},
		{input: 5, expected: Owner{
			Base: model2.Base{
				ID: 2,
			},
			Person: model2.Person{
				FirstName: "Peter",
				LastName:  "McTavish",
			},
		}},
	}

	for _, testCase := range testCases {
		owner, _ := suite.ownerRepository.FindById(testCase.input)
		assert.Equal(suite.T(), testCase.expected.FirstName, owner.FirstName)
		assert.Equal(suite.T(), testCase.expected.LastName, owner.LastName)
	}
}

func (suite *OwnerRepoTestSuite) Test_FindByLastName() {
	var testCases = []struct {
		input    string
		expected Owner
	}{
		{"Coleman", Owner{
			Person: model2.Person{
				FirstName: "Jean",
				LastName:  "Coleman",
			},
		}},
		{"Schroeder", Owner{
			Person: model2.Person{
				FirstName: "David",
				LastName:  "Schroeder",
			},
		}},
	}

	for _, testCase := range testCases {
		owners, _ := suite.ownerRepository.FindByLastName(testCase.input)
		assert.Equal(suite.T(), testCase.expected.LastName, owners[0].LastName)
		assert.Equal(suite.T(), testCase.expected.FirstName, owners[0].FirstName)
	}
}
