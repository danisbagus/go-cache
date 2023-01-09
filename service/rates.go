package service

import (
	"github.com/danisbagus/go-cache/domain"
	"github.com/danisbagus/go-cache/repository"
)

type (
	IRatesService interface {
		GetUSDToIDRRates() (domain.Rates, error)
	}

	RatesService struct {
		ratesRepo repository.IRatesRepository
	}
)

func NewRatesService(ratesRepo repository.IRatesRepository) IRatesService {
	return &RatesService{
		ratesRepo: ratesRepo,
	}
}

func (s *RatesService) GetUSDToIDRRates() (domain.Rates, error) {
	var rates domain.Rates

	IDRRates, err := s.ratesRepo.GetUSDToIDRRates()
	if err != nil {
		return rates, err
	}

	rates.IDR = IDRRates
	return rates, nil
}
