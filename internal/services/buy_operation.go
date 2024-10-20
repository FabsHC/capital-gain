package services

import (
	"capital-gain/internal/models"
)

type BuyOperation interface {
	Execute(purchase *models.StocksInfo, operation models.CapitalGainInput)
}

type buyOperation struct{}

func NewBuyOperation() BuyOperation {
	return &buyOperation{}
}

func (bo *buyOperation) Execute(stocksInfo *models.StocksInfo, operation models.CapitalGainInput) {
	stocksInfo.CalculateNewAverageSharePrice(operation)
	stocksInfo.AddShares(operation.Quantity)
}
