package models

import "capital-gain/internal/utils"

type (
	StocksInfo struct {
		AveragePrice float64
		Shares       uint
	}

	Profit struct {
		Losses float64
		Gains  float64
	}
)

func (s *StocksInfo) AddShares(shares uint) {
	s.Shares += shares
}

func (s *StocksInfo) SubtractShares(shares uint) {
	s.Shares -= shares
}

func (s *StocksInfo) CalculateNewAverageSharePrice(operation CapitalGainInput) {
	s.AveragePrice = utils.CalculateAverageSharePrice(
		s.Shares,
		operation.Quantity,
		s.AveragePrice,
		operation.UnitCost)
}

func (p *Profit) AddLosses(loss float64) {
	p.Losses += loss
}

func (p *Profit) SubtractLosses() {
	if p.Gains > p.Losses {
		p.Gains -= p.Losses
		p.Losses -= p.Gains
	} else {
		p.Losses -= p.Gains
		p.Gains = utils.ZERO
	}
	if p.Losses < utils.ZERO {
		p.Losses = utils.ZERO
	}
}

func NewStocksInfo() *StocksInfo {
	return &StocksInfo{
		AveragePrice: utils.ZERO,
		Shares:       utils.ZERO,
	}
}

func NewProfit() *Profit {
	return &Profit{
		Losses: utils.ZERO,
		Gains:  utils.ZERO,
	}
}
