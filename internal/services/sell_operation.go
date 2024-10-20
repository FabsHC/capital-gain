package services

import (
	"capital-gain/internal/models"
	"capital-gain/internal/utils"
)

type SellOperation interface {
	Execute(purchase *models.StocksInfo, sale *models.Profit, operation models.CapitalGainInput)
}

type sellOperation struct{}

func NewSellOperation() SellOperation {
	return &sellOperation{}
}

func (so *sellOperation) Execute(stocksInfo *models.StocksInfo, profit *models.Profit, operation models.CapitalGainInput) {
	unitCostTotal := operation.GetTotalCost(operation.UnitCost)
	averagePriceTotal := operation.GetTotalCost(stocksInfo.AveragePrice)
	stocksInfo.SubtractShares(operation.Quantity)

	if operation.UnitCost == stocksInfo.AveragePrice {
		profit.Gains = utils.ZERO
		return
	}

	if operation.UnitCost < stocksInfo.AveragePrice {
		profit.AddLosses(averagePriceTotal - unitCostTotal)
		profit.Gains = utils.ZERO
		return
	}

	profit.Gains = unitCostTotal - averagePriceTotal
	profit.SubtractLosses()
}
