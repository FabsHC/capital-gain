package models

import "capital-gain/internal/utils"

type (
	Purchase struct {
		AveragePrice float64
		Stocks       uint
	}

	Sale struct {
		Losses float64
		Gains  float64
	}
)

func (p *Purchase) AddShares(shares uint) {
	p.Stocks += shares
}

func (p *Purchase) SubtractShares(shares uint) {
	p.Stocks -= shares
}

func (s *Sale) AddLosses(loss float64) {
	s.Losses += loss
}

func (s *Sale) SubtractLosses() {
	if s.Gains > s.Losses {
		s.Gains -= s.Losses
		s.Losses -= s.Gains
	} else {
		s.Losses -= s.Gains
		s.Gains = utils.ZERO
	}
	if s.Losses < utils.ZERO {
		s.Losses = utils.ZERO
	}
}

func NewPurchase() *Purchase {
	return &Purchase{
		AveragePrice: utils.ZERO,
		Stocks:       utils.ZERO,
	}
}

func NewSale() *Sale {
	return &Sale{
		Losses: utils.ZERO,
		Gains:  utils.ZERO,
	}
}
