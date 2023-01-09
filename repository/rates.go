package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/danisbagus/go-cache/domain"
	"github.com/danisbagus/go-cache/utils/net"
)

type (
	IRatesRepository interface {
		GetUSDToIDRRates() (float64, error)
	}

	RatesRepository struct {
	}
)

func NewRatesRepository() *RatesRepository {
	return &RatesRepository{}
}

func (r *RatesRepository) GetUSDToIDRRates() (float64, error) {
	url := os.Getenv("RATES_URL")
	apiKey := os.Getenv("API_KEY")
	method := "get"

	url = fmt.Sprintf("%s/latest?symbol=IDR&base=USD", url)

	header := make(map[string]string)
	header["apikey"] = apiKey

	response, err := net.HttpRequest(url, method, "", header)
	if err != nil {
		return 0, err
	}

	resData := new(domain.RatesResData)
	if response != nil {
		json.Unmarshal(response, resData)
	}

	return resData.Rates.IDR, nil
}
