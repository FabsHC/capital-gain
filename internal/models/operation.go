package models

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

func (s *Sale) SubtractLosses(loss float64) {
	s.Losses -= loss
	if s.Losses < 0 {
		s.Losses = 0
	}
}

func NewPurchase(averagePrice float64, stocks uint) *Purchase {
	return &Purchase{
		AveragePrice: averagePrice,
		Stocks:       stocks,
	}
}

func NewSale(losses, gains float64) *Sale {
	return &Sale{
		Losses: losses,
		Gains:  gains,
	}
}
