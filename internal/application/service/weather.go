package service

import (
	weather "goexpert-temperature-system-by-cep/internal/domain"
	"goexpert-temperature-system-by-cep/internal/infra/repository"
)

//go:generate mockery --name WeatherService --outpkg mock --output mock --filename weather.go --with-expecter=true

type WeatherService interface {
	GetWeatherByLocation(location string) (*weather.Weather, error)
}

type weatherService struct {
	repository repository.WeatherRepository
}

func NewWeatherService(repo repository.WeatherRepository) WeatherService {
	return &weatherService{
		repository: repo,
	}
}

func (s *weatherService) GetWeatherByLocation(location string) (*weather.Weather, error) {
	return s.repository.GetWeatherByLocation(location)
}
