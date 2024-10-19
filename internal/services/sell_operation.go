package services

import (
	"capital-gain/internal/models"
	"capital-gain/internal/utils"
)

type SellOperation interface {
	Execute(purchase *models.Purchase, sale *models.Sale, operation models.CapitalGainInput)
}

type sellOperation struct{}

func NewSellOperation() SellOperation {
	return &sellOperation{}
}

func (so *sellOperation) Execute(purchase *models.Purchase, sale *models.Sale, operation models.CapitalGainInput) {
	unitCostTotal := operation.GetTotalCost(operation.UnitCost)
	averagePriceTotal := operation.GetTotalCost(purchase.AveragePrice)

	if operation.UnitCost == purchase.AveragePrice {
		sale.Gains = utils.ZERO
		return
	}

	if operation.UnitCost < purchase.AveragePrice {
		sale.AddLosses(averagePriceTotal - unitCostTotal)
		sale.Gains = utils.ZERO
		return
	}

	sale.Gains = unitCostTotal - averagePriceTotal
	sale.SubtractLosses()
}
