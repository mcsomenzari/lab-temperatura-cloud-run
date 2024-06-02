package config

type Config struct {
	ViaCepURL     string
	WeatherAPIURL string
	WeatherAPIKey string
}

func NewConfig() *Config {
	return &Config{
		ViaCepURL:     "https://viacep.com.br/ws",
		WeatherAPIURL: "http://api.weatherapi.com/v1/current.json",
		WeatherAPIKey: "e5bd00e528e346ff8a840254213009",
	}
}
