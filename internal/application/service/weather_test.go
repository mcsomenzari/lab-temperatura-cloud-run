package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"goexpert-temperature-system-by-cep/internal/domain"
	mock_repositoy "goexpert-temperature-system-by-cep/internal/infra/repository/mock"
)

type WeatherServiceSuite struct {
	suite.Suite
	mockRepo *mock_repositoy.WeatherRepository
	service  WeatherService
}

func (suite *WeatherServiceSuite) SetupSuite() {
	suite.mockRepo = mock_repositoy.NewWeatherRepository(suite.T())
	suite.service = NewWeatherService(suite.mockRepo)
}

func (suite *WeatherServiceSuite) TestGetWeatherByLocation() {
	suite.T().Run("Success", func(t *testing.T) {
		expectedWeather := &domain.Weather{TempCelsius: 25.0}
		expectedLocation := "New York"
		suite.mockRepo.On("GetWeatherByLocation", expectedLocation).Return(expectedWeather, nil)

		result, err := suite.service.GetWeatherByLocation(expectedLocation)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedWeather, result)

		suite.mockRepo.AssertExpectations(t)
	})

	suite.T().Run("Error", func(t *testing.T) {
		expectedError := errors.New("repository error")
		expectedLocation := "London"
		suite.mockRepo.On("GetWeatherByLocation", expectedLocation).Return(nil, expectedError)

		result, err := suite.service.GetWeatherByLocation(expectedLocation)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, expectedError.Error())

		suite.mockRepo.AssertExpectations(t)
	})
}

func (suite *WeatherServiceSuite) TearDownSuite() {
	suite.mockRepo = nil
	suite.service = nil
}

func TestWeatherServiceSuite(t *testing.T) {
	suite.Run(t, new(WeatherServiceSuite))
}
