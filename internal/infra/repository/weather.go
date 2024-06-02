package repository

import (
	"encoding/json"
	"fmt"
	weather "goexpert-temperature-system-by-cep/internal/domain"
	"net/http"
	"net/url"
)

//go:generate mockery --name WeatherRepository --outpkg mock --output mock --filename weather.go --with-expecter=true

type WeatherRepository interface {
	GetWeatherByLocation(location string) (*weather.Weather, error)
}

type weatherRepository struct {
	client *http.Client
	url    string
	apiKey string
}

func NewWeatherRepository(client *http.Client, url, apiKey string) WeatherRepository {
	return &weatherRepository{
		client: client,
		url:    url,
		apiKey: apiKey,
	}
}

func (r *weatherRepository) GetWeatherByLocation(location string) (*weather.Weather, error) {
	escapedLocation := url.QueryEscape(location)
	url := fmt.Sprintf("%s?key=%s&q=%s", r.url, r.apiKey, escapedLocation)
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var weatherResponse weather.WeatherResponse
		if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
			return nil, err
		}
		return &weatherResponse.Current, nil
	}

	return nil, fmt.Errorf("could not fetch weather data")
}
