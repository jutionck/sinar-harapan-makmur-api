package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

var brandDummies = []model.Brand{
	{
		BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		Name:      "Honda",
	},
	{
		BaseModel: model.BaseModel{ID: "2", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		Name:      "Toyota",
	},
	{
		BaseModel: model.BaseModel{ID: "3", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		Name:      "BMW",
	},
}

type BrandRepoTestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
}

func (suite *BrandRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.mock = mock
	dialect := postgres.New(postgres.Config{
		Conn: db,
	})
	suite.DB, err = gorm.Open(dialect)
	assert.NoError(suite.T(), err)
}

func (suite *BrandRepoTestSuite) TestGetAllBrand_Success() {
	brandRowDummies := make([]model.Brand, len(brandDummies))
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	for i, brand := range brandDummies {
		brandRowDummies[i] = brand
		rows.AddRow(brand.ID, brand.Name, brand.CreatedAt, brand.UpdatedAt)
	}
	expectedQuery := `SELECT \* FROM "mst_brand"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rows)

	repo := NewBrandRepository(suite.DB)
	listBrand, err := repo.List()
	assert.Equal(suite.T(), brandRowDummies, listBrand)
	assert.NoError(suite.T(), err)
}

func (suite *BrandRepoTestSuite) TestGetAllMenu_DBErrorFail() {
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	for _, brand := range brandDummies {
		rows.AddRow(brand.ID, brand.Name, brand.CreatedAt, brand.UpdatedAt)
	}
	expectedQuery := `SELECT \* FROM "mst_brand"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("db error"))
	repo := NewBrandRepository(suite.DB)
	listMenu, err := repo.List()
	assert.Nil(suite.T(), listMenu)
	assert.Error(suite.T(), err)
}

func TestBrandRepoTestSuite(t *testing.T) {
	suite.Run(t, new(BrandRepoTestSuite))
}
