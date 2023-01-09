package domain

type (
	RatesResData struct {
		Rates `json:"rates"`
	}

	Rates struct {
		IDR float64 `json:"IDR"`
	}
)
