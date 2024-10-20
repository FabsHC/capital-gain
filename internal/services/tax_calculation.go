package services

import (
	"capital-gain/internal/models"
	"capital-gain/internal/utils"
)

type TaxCalculation interface {
	Execute(operations []models.CapitalGainInput) []*models.CapitalGainOutput
}

type (
	taxCalculation struct {
		buyOperation  BuyOperation
		sellOperation SellOperation
	}
)

func NewTaxCalculation(buyOperation BuyOperation, sellOperation SellOperation) TaxCalculation {
	return &taxCalculation{
		buyOperation:  buyOperation,
		sellOperation: sellOperation,
	}
}

func (tc *taxCalculation) Execute(operations []models.CapitalGainInput) []*models.CapitalGainOutput {
	var taxes []*models.CapitalGainOutput
	stocksInfo := models.NewStocksInfo()
	profit := models.NewProfit()

	for _, operation := range operations {
		var tax *models.CapitalGainOutput

		switch operation.Operation {
		case models.BUY_OPERATION:
			tc.buyOperation.Execute(stocksInfo, operation)
			tax = models.NewCapitalGainOutput(utils.ZERO)

		case models.SELL_OPERATION:
			tc.sellOperation.Execute(stocksInfo, profit, operation)
			tax = models.NewCapitalGainOutput(getSalesGains(profit, operation))
		}

		taxes = append(taxes, tax)
	}

	return taxes
}

func getSalesGains(sale *models.Profit, operation models.CapitalGainInput) float64 {
	if operation.GetTotalCost(operation.UnitCost) <= 20000 {
		return utils.ZERO
	}

	return sale.Gains
}
