package usecase

import (
	"ganho-capital/internal/application/model"
)

type SellOperation struct{}

func NewSellOperation() *SellOperation {
	return &SellOperation{}
}

func (so *SellOperation) Execute(purchase *model.Purchase, sale *model.Sale, operation model.CapitalGainInput) {

	if operation.UnitCost == purchase.AveragePrice {
		sale.ProfitGains = 0
		return
	}
	if operation.UnitCost < purchase.AveragePrice {
		sale.AddProfitLoss((float64(operation.Quantity) * purchase.AveragePrice) - (float64(operation.Quantity) * operation.UnitCost))
		sale.ProfitGains = 0
		return
	}
	if float64(operation.Quantity)*operation.UnitCost <= 20000 {
		sale.ProfitGains = 0
		return
	}

	sale.ProfitGains = (float64(operation.Quantity) * operation.UnitCost) - (float64(operation.Quantity) * purchase.AveragePrice)

	if sale.ProfitGains > sale.TotalProfitLoss {
		sale.ProfitGains = sale.ProfitGains - sale.TotalProfitLoss
		sale.SubtractProfitLoss(sale.ProfitGains)
	} else {
		sale.SubtractProfitLoss(sale.ProfitGains)
		sale.ProfitGains = 0
	}

	sale.SubtractProfitLoss(sale.ProfitGains)

}
