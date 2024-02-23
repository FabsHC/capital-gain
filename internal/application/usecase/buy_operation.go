package usecase

import (
	"ganho-capital/internal/application/model"
	"ganho-capital/internal/application/util"
)

type BuyOperation struct{}

func NewBuyOperation() *BuyOperation {
	return &BuyOperation{}
}

func (bo *BuyOperation) Execute(purchase *model.Purchase, operation model.CapitalGainInput) {
	purchase.AveragePrice = util.CalculateAveragePrice(
		purchase.TotalShares,
		operation.Quantity,
		purchase.AveragePrice,
		operation.UnitCost)
	purchase.AddShares(operation.Quantity)
}
