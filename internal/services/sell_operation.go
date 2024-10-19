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
	unitCostTotal := float64(operation.Quantity) * operation.UnitCost
	averagePriceTotal := float64(operation.Quantity) * purchase.AveragePrice

	if operation.UnitCost == purchase.AveragePrice {
		sale.Gains = 0
		return
	}

	if operation.UnitCost < purchase.AveragePrice {
		sale.AddLosses(averagePriceTotal - unitCostTotal)
		sale.Gains = 0
		return
	}

	if unitCostTotal <= 20000 {
		sale.Gains = 0
		return
	}

	sale.Gains = unitCostTotal - averagePriceTotal
	if sale.Gains > sale.Losses {
		sale.Gains -= sale.Losses
		sale.SubtractLosses(sale.Losses)
	} else {
		sale.SubtractLosses(sale.Gains)
		sale.Gains = 0
	}
}
