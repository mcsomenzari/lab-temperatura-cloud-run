package service

import (
	zipcode "goexpert-temperature-system-by-cep/internal/domain"
	"goexpert-temperature-system-by-cep/internal/infra/repository"
)

//go:generate mockery --name ZipCodeService --outpkg mock --output mock --filename zipcode.go --with-expecter=true

type ZipCodeService interface {
	GetLocationByZipCode(zipCode string) (*zipcode.Location, error)
}

type zipCodeService struct {
	repository repository.ZipCodeRepository
}

func NewZipCodeService(repo repository.ZipCodeRepository) ZipCodeService {
	return &zipCodeService{
		repository: repo,
	}
}

func (s *zipCodeService) GetLocationByZipCode(zipCode string) (*zipcode.Location, error) {
	return s.repository.GetLocationByZipCode(zipCode)
}
