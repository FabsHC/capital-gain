package models

type (
	Purchase struct {
		AveragePrice float64
		Stock        uint
	}

	Sale struct {
		TotalProfitLoss float64
		ProfitGains     float64
	}
)

func (p *Purchase) AddShares(shares uint) {
	p.Stock += shares
}

func (p *Purchase) SubtractShares(shares uint) {
	p.Stock -= shares
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

func NewPurchase(averagePrice float64, stocks uint) *Purchase {
	return &Purchase{
		AveragePrice: averagePrice,
		Stock:        stocks,
	}
}

func NewSale(totalProfitLoss, profitGains float64) *Sale {
	return &Sale{
		TotalProfitLoss: totalProfitLoss,
		ProfitGains:     profitGains,
	}
}
