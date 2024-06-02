package main

import (
	"goexpert-temperature-system-by-cep/internal/application/config"
	"goexpert-temperature-system-by-cep/internal/application/service"
	"goexpert-temperature-system-by-cep/internal/infra/client"
	"goexpert-temperature-system-by-cep/internal/infra/entrypoint"
	"goexpert-temperature-system-by-cep/internal/infra/entrypoint/controller"
	"goexpert-temperature-system-by-cep/internal/infra/repository"
)

func main() {
	conf := config.NewConfig()
	httpClient := client.NewHTTPClient()

	zipCodeRepo := repository.NewZipCodeRepository(httpClient, conf.ViaCepURL)
	weatherRepo := repository.NewWeatherRepository(httpClient, conf.WeatherAPIURL, conf.WeatherAPIKey)

	zipCodeService := service.NewZipCodeService(zipCodeRepo)
	weatherService := service.NewWeatherService(weatherRepo)

	weatherController := controller.NewWeatherController(weatherService, zipCodeService)

	r := entrypoint.SetupRouter(weatherController)
	r.Run(":8080")
}
