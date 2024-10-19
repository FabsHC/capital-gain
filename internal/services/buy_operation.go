package services

import (
	"capital-gain/internal/models"
	"capital-gain/internal/utils"
)

type BuyOperation interface {
	Execute(purchase *models.Purchase, operation models.CapitalGainInput)
}

type buyOperation struct{}

func NewBuyOperation() BuyOperation {
	return &buyOperation{}
}

func (bo *buyOperation) Execute(purchase *models.Purchase, operation models.CapitalGainInput) {
	purchase.AveragePrice = utils.CalculateAveragePrice(
		purchase.Stock,
		operation.Quantity,
		purchase.AveragePrice,
		operation.UnitCost)
	purchase.AddShares(operation.Quantity)
}
