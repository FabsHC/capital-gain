package services

import (
	"capital-gain/internal/models"
)

type SellOperation interface {
	Execute(purchase *models.Purchase, sale *models.Sale, operation models.CapitalGainInput)
}

type sellOperation struct{}

func NewSellOperation() SellOperation {
	return &sellOperation{}
}

func (so *sellOperation) Execute(purchase *models.Purchase, sale *models.Sale, operation models.CapitalGainInput) {

	if operation.UnitCost == purchase.AveragePrice {
		sale.Gains = 0
		return
	}
	if operation.UnitCost < purchase.AveragePrice {
		sale.AddLosses((float64(operation.Quantity) * purchase.AveragePrice) - (float64(operation.Quantity) * operation.UnitCost))
		sale.Gains = 0
		return
	}
	if float64(operation.Quantity)*operation.UnitCost <= 20000 {
		sale.Gains = 0
		return
	}

	sale.Gains = (float64(operation.Quantity) * operation.UnitCost) - (float64(operation.Quantity) * purchase.AveragePrice)

	if sale.Gains > sale.Losses {
		sale.Gains = sale.Gains - sale.Losses
		sale.SubtractLosses(sale.Gains)
	} else {
		sale.SubtractLosses(sale.Gains)
		sale.Gains = 0
	}

	sale.SubtractLosses(sale.Gains)

}
