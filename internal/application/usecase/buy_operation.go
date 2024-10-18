package usecase

import (
	"capital-gain/internal/application/model"
	"capital-gain/internal/utils"
)

type BuyOperation interface {
	Execute(purchase *model.Purchase, operation model.CapitalGainInput)
}

type buyOperation struct{}

func NewBuyOperation() BuyOperation {
	return &buyOperation{}
}

func (bo *buyOperation) Execute(purchase *model.Purchase, operation model.CapitalGainInput) {
	purchase.AveragePrice = utils.CalculateAveragePrice(
		purchase.TotalShares,
		operation.Quantity,
		purchase.AveragePrice,
		operation.UnitCost)
	purchase.AddShares(operation.Quantity)
}
