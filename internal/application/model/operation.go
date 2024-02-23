package model

import "errors"

type (
	Purchase struct {
		AveragePrice float64
		TotalShares  uint
	}

	Sale struct {
		TotalProfitLoss float64
		ProfitGains     float64
	}

	SaleError struct {
		Err string `json:"error"`
	}
)

func (p *Purchase) AddShares(shares uint) {
	p.TotalShares += shares
}

func (p *Purchase) SubtractShares(shares uint) {
	p.TotalShares -= shares
	if p.TotalShares < 0 {
		p.TotalShares = 0
	}
}

func (s *Sale) AddProfitLoss(loss float64) {
	s.TotalProfitLoss += loss
}

func (s *Sale) SubtractProfitLoss(loss float64) {
	s.TotalProfitLoss -= loss
	if s.TotalProfitLoss < 0 {
		s.TotalProfitLoss = 0
	}
}

func (s *Sale) Validate(totalShares, sharesToSell uint) error {
	if sharesToSell > totalShares {
		return errors.New("can't sell more stocks than you have")
	}

	return nil
}

func NewPurchase(averagePrice float64, totalShares uint) *Purchase {
	return &Purchase{
		AveragePrice: averagePrice,
		TotalShares:  totalShares,
	}
}

func NewSale(totalProfitLoss, profitGains float64) *Sale {
	return &Sale{
		TotalProfitLoss: totalProfitLoss,
		ProfitGains:     profitGains,
	}
}

func NewSaleError(err string) *SaleError {
	return &SaleError{
		Err: err,
	}
}
