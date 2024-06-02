package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"goexpert-temperature-system-by-cep/internal/domain"
	mock_repository "goexpert-temperature-system-by-cep/internal/infra/repository/mock"
)

type ZipCodeServiceSuite struct {
	suite.Suite
	mockRepo *mock_repository.ZipCodeRepository
	service  ZipCodeService
}

func (suite *ZipCodeServiceSuite) SetupSuite() {
	suite.mockRepo = mock_repository.NewZipCodeRepository(suite.T())
	suite.service = NewZipCodeService(suite.mockRepo)
}

func (suite *ZipCodeServiceSuite) TestGetLocationByZipCode() {
	suite.T().Run("Success", func(t *testing.T) {
		expectedLocation := &domain.Location{
			Cep:        "12345678",
			Logradouro: "Rua Exemplo",
			Bairro:     "Bairro Exemplo",
			Localidade: "Cidade Exemplo",
			Uf:         "UF",
		}

		expectedZipCode := "10001"
		suite.mockRepo.On("GetLocationByZipCode", expectedZipCode).Return(expectedLocation, nil)

		result, err := suite.service.GetLocationByZipCode(expectedZipCode)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedLocation, result)

		suite.mockRepo.AssertExpectations(t)
	})

	suite.T().Run("Error", func(t *testing.T) {
		expectedError := errors.New("repository error")
		expectedZipCode := "90210"
		suite.mockRepo.On("GetLocationByZipCode", expectedZipCode).Return(nil, expectedError)

		result, err := suite.service.GetLocationByZipCode(expectedZipCode)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, expectedError.Error())

		suite.mockRepo.AssertExpectations(t)
	})
}

func (suite *ZipCodeServiceSuite) TearDownSuite() {
	suite.mockRepo = nil
	suite.service = nil
}

func TestZipCodeServiceSuite(t *testing.T) {
	suite.Run(t, new(ZipCodeServiceSuite))
}
