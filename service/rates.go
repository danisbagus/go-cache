package service

import (
	"time"

	"github.com/danisbagus/go-cache/domain"
	"github.com/danisbagus/go-cache/repository"

	"github.com/danisbagus/go-cache/utils/cache"
)

type (
	IRatesService interface {
		GetUSDToIDRRates() (domain.Rates, error)
	}

	RatesService struct {
		cache     *cache.Cache
		ratesRepo repository.IRatesRepository
	}
)

func NewRatesService(ratesRepo repository.IRatesRepository, cache *cache.Cache) IRatesService {
	return &RatesService{
		ratesRepo: ratesRepo,
		cache:     cache,
	}
}

func (s *RatesService) GetUSDToIDRRates() (domain.Rates, error) {
	var rates domain.Rates

	value, found := s.cache.Get("IDR_RATES")
	if found {
		rates.IDR = value.(float64)
		return rates, nil
	}

	IDRRates, err := s.ratesRepo.GetUSDToIDRRates()
	if err != nil {
		return rates, err
	}

	s.cache.Set("IDR_RATES", IDRRates, 1*time.Minute)
	rates.IDR = IDRRates
	return rates, nil
}
